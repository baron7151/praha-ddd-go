package domainteam

import (
	"testing"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
)

func TestNewTeamName_InValidName(t *testing.T) {
	_, err := NewTeamName("test")
	if err == nil {
		t.Error("エラーが発生しませんでした")
	}
}

func TestNewTeamName_ValidName(t *testing.T) {
	_, err := NewTeamName("1")
	if err != nil {
		t.Errorf("エラーが発生しました: %v", err)
	}
}

func TestNewTeamEntity_ValidTeamEntity(t *testing.T) {
	teamId, _ := domaincommon.NewBaseUUID("")
	teamName, _ := NewTeamName("1")
	pairId1, _ := domaincommon.NewBaseUUID("")
	pairId2, _ := domaincommon.NewBaseUUID("")
	pairIds := []domaincommon.BaseUUID{pairId1, pairId2}
	userId1, _ := domaincommon.NewBaseUUID("")
	userId2, _ := domaincommon.NewBaseUUID("")
	userIds := []domaincommon.BaseUUID{userId1, userId2}
	team := NewTeamEntity(teamId, teamName, WithPairIds(&pairIds), WithUserIds(&userIds))
	if team.teamId != teamId {
		t.Error("teamIdが設定されていません")
	}
	if team.teamName != teamName {
		t.Error("teamNameが設定されていません")
	}
	if team.pairIds == nil {
		t.Error("pairIdsが設定されていません")
	}
	if team.userIds == nil {
		t.Error("userIdsが設定されていません")
	}
}

func TestValidateTeamUserCount_ValidTeamUserCount(t *testing.T) {
	userId1, _ := domaincommon.NewBaseUUID("")
	userId2, _ := domaincommon.NewBaseUUID("")
	userId3, _ := domaincommon.NewBaseUUID("")
	userIds := []domaincommon.BaseUUID{userId1, userId2, userId3}
	result := ValidateTeamUserCount(&userIds)
	if result != true {
		t.Error("チームのユーザー数が不正です")
	}
}

func TestCountTeamUser(t *testing.T) {
	teamId, _ := domaincommon.NewBaseUUID("")
	teamName, _ := NewTeamName("1")
	pairId1, _ := domaincommon.NewBaseUUID("")
	pairId2, _ := domaincommon.NewBaseUUID("")
	pairIds := []domaincommon.BaseUUID{pairId1, pairId2}
	userId1, _ := domaincommon.NewBaseUUID("")
	userId2, _ := domaincommon.NewBaseUUID("")
	userIds := []domaincommon.BaseUUID{userId1, userId2}
	team := NewTeamEntity(teamId, teamName, WithPairIds(&pairIds), WithUserIds(&userIds))
	result := team.CountTeamUser()
	if result != 2 {
		t.Error("チームのユーザー数が不正です")
	}
}

func TestIsMinimumTeamUser(t *testing.T) {
	teamId, _ := domaincommon.NewBaseUUID("")
	teamName, _ := NewTeamName("1")
	pairId1, _ := domaincommon.NewBaseUUID("")
	pairId2, _ := domaincommon.NewBaseUUID("")
	pairIds := []domaincommon.BaseUUID{pairId1, pairId2}
	userId1, _ := domaincommon.NewBaseUUID("")
	userId2, _ := domaincommon.NewBaseUUID("")
	userIds := []domaincommon.BaseUUID{userId1, userId2}
	team := NewTeamEntity(teamId, teamName, WithPairIds(&pairIds), WithUserIds(&userIds))
	result := team.IsMinimumTeamUser()
	if result != true {
		t.Error("チームのユーザー数が不正です")
	}
}

func TestChangeTeamName(t *testing.T) {
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
	result := team.ChangeTeamName(newTeamName)
	if result.teamName != newTeamName {
		t.Error("チーム名が変更されていません")
	}
}
