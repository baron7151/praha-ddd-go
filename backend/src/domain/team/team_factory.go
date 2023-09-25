package domainteam

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
)

func Create(teamIdStr string, teamNameStr string, pairIdsStr *[]string, userIdsStr *[]string) (TeamEntity, error) {
	teamId, err := domaincommon.NewBaseUUID(teamIdStr)
	if err != nil {
		return TeamEntity{}, err
	}
	teamName, err := NewTeamName(teamNameStr)
	if err != nil {
		return TeamEntity{}, err
	}

	var pairIds *[]domaincommon.BaseUUID
	var tempPairIds []domaincommon.BaseUUID
	if pairIdsStr != nil && len(*pairIdsStr) > 0 {
		for _, pairIdStr := range *pairIdsStr {
			tempPairId, err := domaincommon.NewBaseUUID(pairIdStr)
			if err != nil {
				return TeamEntity{}, err
			}
			tempPairIds = append(tempPairIds, tempPairId)
		}
		pairIds = &tempPairIds
	}

	var userIds *[]domaincommon.BaseUUID
	var tempUserIds []domaincommon.BaseUUID
	if userIdsStr != nil && len(*userIdsStr) > 0 {
		for _, userIdStr := range *userIdsStr {
			tempUserId, err := domaincommon.NewBaseUUID(userIdStr)
			if err != nil {
				return TeamEntity{}, err
			}
			tempUserIds = append(tempUserIds, tempUserId)
		}
		userIds = &tempUserIds
	}
	return NewTeamEntity(teamId, teamName, WithPairIds(pairIds), WithUserIds(userIds)), nil
}

func Reconstruct(teamEntity TeamEntity, newTeamNameStr string, teamRepository ITeamRepository) (TeamEntity, error) {
	newTeamName, err := NewTeamName(newTeamNameStr)
	if err != nil {
		return TeamEntity{}, err
	}
	duplicatedTeamNameCheck, err := teamRepository.Exists(newTeamName)
	if duplicatedTeamNameCheck {
		return TeamEntity{}, domaincommon.NewDomainError("teamName is duplicated")
	}
	return teamEntity.ChangeTeamName(newTeamName), nil

}
