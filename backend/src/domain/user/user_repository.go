package domainuser

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainpair "github.com/baron7151/praha-ddd-go/src/domain/pair"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
)

type IUserRepository interface {
	FindByUserId(userId UserId) (UserEntity, error)
	FindByEmail(email domaincommon.Email) (UserEntity, error)
	FindByTeamId(teamId domainteam.TeamId) ([]UserEntity, error)
	FindByPairId(pairId domainpair.PairId) ([]UserEntity, error)
	FindByManyUserIds(userIds []UserId) ([]UserEntity, error)
	Save(saveUesr UserEntity) error
	Exists(email domaincommon.Email) (bool, error)
}
