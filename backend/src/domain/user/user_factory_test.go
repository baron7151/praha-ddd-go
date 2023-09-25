package domainuser

import (
	"testing"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	userId, _ := domaincommon.NewBaseUUID("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	pairId, _ := domaincommon.NewBaseUUID("")
	teamId, _ := domaincommon.NewBaseUUID("")
	pairIdStr := pairId.GetValue()
	teamIdStr := teamId.GetValue()
	userEntity := NewUserEntity(userId, userName, email, status, WithPairId(&pairId), WithTeamId(&teamId))
	user, _ := Create(userId.GetValue(), userName.GetValue(), email.GetValue(), "ACTIVE", &pairIdStr, &teamIdStr)
	assert.Equal(t, userEntity, user)
}

func TestReconstruct(t *testing.T) {
	userId, _ := domaincommon.NewBaseUUID("")
	userName, _ := NewUserName("test1")
	email, _ := domaincommon.NewEmail("test1@example.com")
	status := ACTIVE
	pairId, _ := domaincommon.NewBaseUUID("")
	teamId, _ := domaincommon.NewBaseUUID("")
	userEntity := NewUserEntity(userId, userName, email, status, WithPairId(&pairId), WithTeamId(&teamId))
	newUserName := "test2"
	newUserEmail := "test2@example.com"
	newUserStatus := "INACTIVE"
	newPairId, _ := domaincommon.NewBaseUUID("")
	newPairIdStr := newPairId.GetValue()
	newTeamId, _ := domaincommon.NewBaseUUID("")
	newTeamIdStr := newTeamId.GetValue()
	newUser, _ := Reconstruct(userEntity, newUserName, newUserEmail, newUserStatus, &newPairIdStr, &newTeamIdStr)
	assert.Equal(t, newUserName, newUser.GetUserName().GetValue())
	assert.Equal(t, newUserEmail, newUser.GetEmail().GetValue())
	assert.Equal(t, INACTIVE, newUser.GetUserStatus())
	assert.Equal(t, &newPairId, newUser.GetPairId())
	assert.Equal(t, &newTeamId, newUser.GetTeamId())
}
