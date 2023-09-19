package infra

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	dsn := `host=` + host + ` user=` + user + ` password=` + password + ` dbname=` + dbname + ` port=` + port + ` sslmode=disable TimeZone=Asia/Tokyo`
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() {
	rootdir := os.Getenv("PROJECT_ROOT_DIR")
	db, err := ConnectDB()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(Task{}, TaskProgress{})
	db.AutoMigrate(Team{}, Pair{}, User{})

	//Init Test Data.
	CleaningAllTable(db)
	err = CreateTestData(db, rootdir)
	if err != nil {
		fmt.Println(err)
		panic("failed to create test data")
	}
}

func CleaningAllTable(db *gorm.DB) {
	db.Exec("DELETE FROM task_progresses")
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM users")
	db.Exec("DELETE FROM pairs")
	db.Exec("DELETE FROM teams")
}

func CreateTestData(db *gorm.DB, rootdir string) error {

	err := InsertCSVDataToDB(db, rootdir+"src/testutil/testdata/Team.csv", &Team{})
	if err != nil {
		return err
	}
	err = InsertCSVDataToDB(db, rootdir+"src/testutil/testdata/Pair.csv", &Pair{})
	if err != nil {
		return err
	}
	err = InsertCSVDataToDB(db, rootdir+"src/testutil/testdata/User.csv", &User{})
	if err != nil {
		return err
	}
	err = InsertCSVDataToDB(db, rootdir+"src/testutil/testdata/Task.csv", &Task{})
	if err != nil {
		return err
	}
	err = InsertCSVDataToDB(db, rootdir+"src/testutil/testdata/TaskProgress.csv", &TaskProgress{})
	if err != nil {
		return err
	}
	return nil
}

func setFieldValue(fieldValue reflect.Value, value string) error {
	fieldType := fieldValue.Type()

	// ポインタ型の場合に値を設定
	if fieldType.Kind() == reflect.Ptr {
		if fieldValue.IsNil() {
			fieldValue.Set(reflect.New(fieldType.Elem()))
		}
		fieldValue = fieldValue.Elem()
	}

	switch fieldValue.Kind() {
	case reflect.String:
		fieldValue.SetString(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		fieldValue.SetInt(int64(intValue))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, fieldType.Bits())
		if err != nil {
			return err
		}
		fieldValue.SetUint(uintValue)
	// 他のデータ型に対する変換ルールを追加できます

	default:
		return fmt.Errorf("サポートされていないデータ型: %v", fieldType)
	}
	return nil
}

func InsertCSVDataToDB(db *gorm.DB, filePath string, dbModel interface{}) error {
	// CSVファイルをオープン
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// CSVリーダーを作成
	csvReader := csv.NewReader(file)

	// ヘッダー行を読み込み
	header, err := csvReader.Read()
	if err != nil {
		return err
	}
	// フィールド名から構造体のフィールドをマッピング
	fieldMap := make(map[string]int)
	modelType := reflect.TypeOf(dbModel).Elem()
	for i, name := range header {
		_, exists := modelType.FieldByName(name)
		if exists {
			fieldMap[name] = i
		}
	}
	for {
		record, err := csvReader.Read()
		if err != nil {
			break // ファイルの終わりに達した場合、ループを終了
		}

		modelValue := reflect.New(modelType).Elem()
		for fieldName, fieldIndex := range fieldMap {
			fieldValue := modelValue.FieldByName(fieldName)
			if fieldValue.IsValid() {
				fieldStr := record[fieldIndex]
				if err := setFieldValue(fieldValue, fieldStr); err != nil {
					return err
				}
			}
		}

		// データベースに挿入
		if err := db.Create(modelValue.Addr().Interface()).Error; err != nil {
			return err
		}
	}
	return nil
}
