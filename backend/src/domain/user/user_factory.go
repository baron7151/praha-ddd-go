package domainuser

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
)

func Create(userIdStr string, userNameStr string, emailStr string, userStatusStr string, pairIdStr *string, teamIdStr *string) (UserEntity, error) {
	userId, err := domaincommon.NewBaseUUID(userIdStr)
	if err != nil {
		return UserEntity{}, err
	}
	userName, err := NewUserName(userNameStr)
	if err != nil {
		return UserEntity{}, err
	}
	email, err := domaincommon.NewEmail(emailStr)
	if err != nil {
		return UserEntity{}, err
	}
	userStatus, err := StringToUserStatus(userStatusStr)
	if err != nil {
		return UserEntity{}, err
	}
	var pairId *domaincommon.BaseUUID
	if pairIdStr != nil {
		tempPairId, err := domaincommon.NewBaseUUID(*pairIdStr)
		if err != nil {
			return UserEntity{}, err
		}
		pairId = &tempPairId
	}

	var teamId *domaincommon.BaseUUID
	if teamIdStr != nil {
		tempTeamId, err := domaincommon.NewBaseUUID(*teamIdStr)
		if err != nil {
			return UserEntity{}, err
		}
		teamId = &tempTeamId
	}
	return NewUserEntity(userId, userName, email, userStatus, WithPairId(pairId), WithTeamId(teamId)), nil
}

func Reconstruct(userEntity UserEntity, newUserNameStr string, newUserEmailStr, newUserStatus string, newPairIdStr *string, newTeamIdStr *string) (UserEntity, error) {
	newUserName, err := NewUserName(newUserNameStr)
	if err != nil {
		return UserEntity{}, err
	}
	newUserEmail, err := domaincommon.NewEmail(newUserEmailStr)
	if err != nil {
		return UserEntity{}, err
	}
	newStatus, err := StringToUserStatus(newUserStatus)
	if err != nil {
		return UserEntity{}, err
	}

	var newPairId *domaincommon.BaseUUID
	if newPairIdStr != nil {
		tempPairId, err := domaincommon.NewBaseUUID(*newPairIdStr)
		if err != nil {
			return UserEntity{}, err
		}
		newPairId = &tempPairId
	}

	var newTeamId *domaincommon.BaseUUID
	if newTeamIdStr != nil {
		tempTeamId, err := domaincommon.NewBaseUUID(*newTeamIdStr)
		if err != nil {
			return UserEntity{}, err
		}
		newTeamId = &tempTeamId
	}
	return NewUserEntity(userEntity.GetUserId(), newUserName, newUserEmail, newStatus, WithPairId(newPairId), WithTeamId(newTeamId)), nil
}
