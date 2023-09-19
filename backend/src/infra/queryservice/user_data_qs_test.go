package queryservice

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/baron7151/praha-ddd-go/src/infra"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../.test.env")
	if err != nil {
		panic(err)
	}
	status := m.Run()

	os.Exit(status)
}

func TestFindAllUsers(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userDataQS := NewUserDataQS(db)
	result, err := userDataQS.FindAllUsers()
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, 9, len(result))
}

func TestFindUserName(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userDataQS := NewUserDataQS(db)
	var user infra.User
	result1 := db.Where("id=1").Find(&user)
	if result1.Error != nil {
		t.Fatal(result1.Error)
	}
	result2, err := userDataQS.FindByUserName(user.UserName)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, user.UserName, result2[0].UserName)
}
