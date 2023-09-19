package domainuser

import (
	"reflect"
	"testing"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainpair "github.com/baron7151/praha-ddd-go/src/domain/pair"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
)

func TestNewUserId_ValidId(t *testing.T) {
	userId, err := NewUserId("")
	if err != nil {
		t.Errorf("エラーが発生しました: %v", err)
	}
	if reflect.TypeOf(userId).String() != "domainuser.UserId" {
		t.Errorf("期待値: domainuser.UserId, 実際の値: %s", reflect.TypeOf(userId).String())
	}
}

func TestNewUserName_ValidName(t *testing.T) {
	userName, err := NewUserName("test")
	if err != nil {
		t.Errorf("エラーが発生しました: %v", err)
	}
	if reflect.TypeOf(userName).String() != "domainuser.UserName" {
		t.Errorf("期待値: domainuser.UserName, 実際の値: %s", reflect.TypeOf(userName).String())
	}
}

func TestEquals_Equal(t *testing.T) {
	userName1, _ := NewUserName("test")
	userName2, _ := NewUserName("test")
	result := userName1.Equals(userName2)
	if !result {
		t.Error("ユーザ名は同じであるはずですが、異なります")
	}
}

func TestWithPairId_ValidPairId(t *testing.T) {
	pairId, _ := domainpair.NewPairId("")
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@exmaple.com")
	status := ACTIVE
	user := NewUserEntity(userId, userName, email, status, WithPairId(&pairId))
	if user.pairId == nil {
		t.Error("pairIdが設定されていません")
	}
}

func TestWithTeamId_ValidTeamId(t *testing.T) {
	teamId, _ := domainteam.NewTeamId("")
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	user := NewUserEntity(userId, userName, email, status, WithTeamId(&teamId))
	if user.teamId == nil {
		t.Error("teamIdが設定されていません")
	}
}

func TestNewUserEntity_ValidUserEntity(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	pairId, _ := domainpair.NewPairId("")
	teamId, _ := domainteam.NewTeamId("")
	user := NewUserEntity(userId, userName, email, status, WithPairId(&pairId), WithTeamId(&teamId))
	if reflect.TypeOf(user).String() != "domainuser.UserEntity" {
		t.Errorf("期待値: domainuser.UserEntity, 実際の値: %s", reflect.TypeOf(user).String())
	}
}

func TestGetAllProperties_ValidProperties(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	user := NewUserEntity(userId, userName, email, status)
	result := user.GetAllProperties()
	if result.UserId.GetValue() != userId.GetValue() {
		t.Errorf("期待値: %s, 実際の値: %s", userId.GetValue(), result.UserId.GetValue())
	}
	if result.UserName.GetValue() != userName.GetValue() {
		t.Errorf("期待値: %s, 実際の値: %s", userName.GetValue(), result.UserName.GetValue())
	}
	if result.Email.GetValue() != email.GetValue() {
		t.Errorf("期待値: %s, 実際の値: %s", email.GetValue(), result.Email.GetValue())
	}
	if result.UserStatus != status {
		t.Errorf("期待値: %s, 実際の値: %s", status, result.UserStatus)
	}

	if result.TeamId != nil {
		t.Error("teamIdが設定されています。")
	}
	if result.PairId != nil {
		t.Error("pairIdが設定されています。")
	}
}

func TestEquals_ValidUserEntity(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	user1 := NewUserEntity(userId, userName, email, status)
	user2 := NewUserEntity(userId, userName, email, status)
	result := user1.Equals(user2)
	if !result {
		t.Error("ユーザは同じであるはずですが、異なります")
	}
}

func TestChangeStatus_ValidStatus(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	user1 := NewUserEntity(userId, userName, email, status)
	user2, _ := user1.ChangeStatus(INACTIVE)
	if user2.status != INACTIVE {
		t.Errorf("期待値: %s, 実際の値: %s", INACTIVE, user2.status)
	}
	if user2.pairId != nil {
		t.Error("pairIdが設定されています。")
	}
	if user2.teamId != nil {
		t.Error("teamIdが設定されています。")
	}
}

func TestChangeUserName_ValidUserName(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	user1 := NewUserEntity(userId, userName, email, status)
	changeUserName, _ := NewUserName("test2")
	user2, _ := user1.ChangeUserName(changeUserName)
	if user2.GetUserName().value != "test2" {
		t.Errorf("期待値: %s, 実際の値: %s", "test2", user2.userName.GetValue())
	}
}

func TestChangeEmail_ValidEmail(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	user1 := NewUserEntity(userId, userName, email, status)
	changeEmail, _ := domaincommon.NewEmail("test2@example.com")
	user2, _ := user1.ChangeEmail(changeEmail)
	if user2.GetEmail().GetValue() != "test2@example.com" {
		t.Errorf("期待値: %s, 実際の値: %s", "test2@example.com", user2.email.GetValue())
	}
}

func TestChangePairId_ValidPairId(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	pairId, _ := domainpair.NewPairId("")
	teamId, _ := domainteam.NewTeamId("")
	user1 := NewUserEntity(userId, userName, email, status, WithPairId(&pairId), WithTeamId(&teamId))
	changePairId, _ := domainpair.NewPairId("")
	user2, _ := user1.ChangePairId(changePairId)
	if user2.GetPairId().GetValue() != changePairId.GetValue() {
		t.Errorf("期待値: %s, 実際の値: %s", changePairId.GetValue(), user2.pairId.GetValue())
	}
}

func TestChangeTeamId_ValidTeamId(t *testing.T) {
	userId, _ := NewUserId("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	pairId, _ := domainpair.NewPairId("")
	teamId, _ := domainteam.NewTeamId("")
	user1 := NewUserEntity(userId, userName, email, status, WithPairId(&pairId), WithTeamId(&teamId))
	changeTeamId, _ := domainteam.NewTeamId("")
	user2, _ := user1.ChangeTeamId(changeTeamId)
	if user2.GetTeamId().GetValue() != changeTeamId.GetValue() {
		t.Errorf("期待値: %s, 実際の値: %s", changeTeamId.GetValue(), user2.teamId.GetValue())
	}
}
