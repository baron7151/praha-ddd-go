package domainuser

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainpair "github.com/baron7151/praha-ddd-go/src/domain/pair"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
)

func Create(userIdStr string, userNameStr string, emailStr string, userStatusStr string, pairIdStr *string, teamIdStr *string) (UserEntity, error) {
	userId, err := NewUserId(userIdStr)
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
	var pairId *domainpair.PairId
	if pairIdStr != nil {
		tempPairId, err := domainpair.NewPairId(*pairIdStr)
		if err != nil {
			return UserEntity{}, err
		}
		pairId = &tempPairId
	}

	var teamId *domainteam.TeamId
	if teamIdStr != nil {
		tempTeamId, err := domainteam.NewTeamId(*teamIdStr)
		if err != nil {
			return UserEntity{}, err
		}
		teamId = &tempTeamId
	}
	return NewUserEntity(userId, userName, email, userStatus, WithPairId(pairId), WithTeamId(teamId)), nil
	/*
		if pairId == nil && teamId == nil {
			return NewUserEntity(userId, userName, email, userStatus), nil
		} else if pairId == nil {
			return NewUserEntity(userId, userName, email, userStatus, WithTeamId(teamId)), nil
		} else if teamId == nil {
			return NewUserEntity(userId, userName, email, userStatus, WithPairId(pairId)), nil
		} else {
			return NewUserEntity(userId, userName, email, userStatus, WithPairId(pairId), WithTeamId(teamId)), nil
		}
	*/
}
