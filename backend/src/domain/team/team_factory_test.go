package domainteam

import (
	"testing"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreate_ValidInput(t *testing.T) {
	teamId, _ := domaincommon.NewBaseUUID("")
	teamName := "1"
	pairId1, _ := domaincommon.NewBaseUUID("")
	pairId2, _ := domaincommon.NewBaseUUID("")
	pairIds := []string{pairId1.GetValue(), pairId2.GetValue()}
	userId1, _ := domaincommon.NewBaseUUID("")
	userId2, _ := domaincommon.NewBaseUUID("")
	userIds := []string{userId1.GetValue(), userId2.GetValue()}
	teamEntity, err := Create(teamId.GetValue(), teamName, &pairIds, &userIds)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	result := teamEntity.GetAllProperties()
	assert.Equal(t, teamId, result.TeamId)
}

func TestReconstruct(t *testing.T) {
	teamId, _ := domaincommon.NewBaseUUID("")
	teamName, _ := NewTeamName("1")
	pairId1, _ := domaincommon.NewBaseUUID("")
	pairId2, _ := domaincommon.NewBaseUUID("")
	pairIds := []domaincommon.BaseUUID{pairId1, pairId2}
	userId1, _ := domaincommon.NewBaseUUID("")
	userId2, _ := domaincommon.NewBaseUUID("")
	userIds := []domaincommon.BaseUUID{userId1, userId2}
	team := NewTeamEntity(teamId, teamName, WithPairIds(&pairIds), WithUserIds(&userIds))
	newTeamName, _ := NewTeamName("2")
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	pr := NewMockITeamRepository(ctrl)
	pr.EXPECT().Exists(newTeamName).Return(false, nil)

	teamEntity, err := Reconstruct(team, newTeamName.value, pr)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, newTeamName, teamEntity.teamName)
}
