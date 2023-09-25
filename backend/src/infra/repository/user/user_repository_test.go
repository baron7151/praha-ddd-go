package userrepository

import (
	"os"
	"testing"

	"github.com/joho/godotenv"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainuser "github.com/baron7151/praha-ddd-go/src/domain/user"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../../.test.env")
	if err != nil {
		panic(err)
	}
	status := m.Run()

	os.Exit(status)
}

func TestFindByUserId(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	var user infra.User
	result := db.Where("id=1").Find(&user)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	userEntity, _ := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
	userId, _ := domaincommon.NewBaseUUID(user.UserId)
	// 実行
	res, err := userRepsitory.FindByUserId(userId)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, userEntity, res)
}

func TestFindByEmail(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	var user infra.User
	result := db.Where("id=1").Find(&user)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	userEntity, _ := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
	email, _ := domaincommon.NewEmail(user.Email)
	// 実行
	res, err := userRepsitory.FindByEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, userEntity, res)
}

func TestFindByTeamId(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	var user infra.User
	result := db.Where("id=1").Find(&user)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	//userEntity, _ := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
	teamId, _ := domaincommon.NewBaseUUID(*user.TeamId)
	// 実行
	res, err := userRepsitory.FindByTeamId(teamId)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	for _, user := range res {
		assert.Equal(t, teamId, *user.GetTeamId())
	}
}

func TestFindByPairId(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	var user infra.User
	result := db.Where("id=1").Find(&user)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	//userEntity, _ := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
	pairId, _ := domaincommon.NewBaseUUID(*user.PairId)
	// 実行
	res, err := userRepsitory.FindByPairId(pairId)
	if err != nil {
		t.Fatal(err)
	}

	for _, user := range res {
		assert.Equal(t, pairId, *user.GetPairId())
	}
}

func TestFindByManyUserIds(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	var user1 infra.User
	result1 := db.Where("id=1").Find(&user1)
	if result1.Error != nil {
		t.Fatal(result1.Error)
	}
	var user2 infra.User
	result2 := db.Where("id=2").Find(&user2)
	if result2.Error != nil {
		t.Fatal(result1.Error)
	}
	userId1, _ := domaincommon.NewBaseUUID(user1.UserId)
	userId2, _ := domaincommon.NewBaseUUID(user2.UserId)
	userIds := []domaincommon.BaseUUID{userId1, userId2}

	// 実行
	res, err := userRepsitory.FindByManyUserIds(userIds)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, 2, len(res))
}

func TestSave_Update(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	var user infra.User
	result := db.Where("id=1").Find(&user)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	updateUserName := "test100"
	userEntity, _ := domainuser.Create(user.UserId, updateUserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
	// 実行
	err := userRepsitory.Save(userEntity)
	if err != nil {
		t.Fatal(err)
	}

	updateUser, _ := userRepsitory.FindByUserId(userEntity.GetUserId())
	// 検証

	assert.Equal(t, updateUserName, updateUser.GetUserName().GetValue())
	assert.Equal(t, user.UserId, updateUser.GetUserId().GetValue())
}

func TestSave_Create(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	userRepsitory := NewUserRepository(db)
	userId, _ := domaincommon.NewBaseUUID("")
	userName, _ := domainuser.NewUserName("test100")
	email, _ := domaincommon.NewEmail("test100@example.com")
	status := domainuser.ACTIVE
	//pairId, _ := domainpair.NewPairId("")
	//teamId, _ := domainteam.NewTeamId("")
	userEntity := domainuser.NewUserEntity(userId, userName, email, status)

	// 実行
	err := userRepsitory.Save(userEntity)
	if err != nil {
		t.Fatal(err)
	}

	saveUser, _ := userRepsitory.FindByUserId(userEntity.GetUserId())
	// 検証
	assert.Equal(t, saveUser, userEntity)
}
