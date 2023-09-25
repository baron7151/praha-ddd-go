package domainuser

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
)

type IUserRepository interface {
	FindByUserId(userId domaincommon.BaseUUID) (UserEntity, error)
	FindByEmail(email domaincommon.Email) (UserEntity, error)
	FindByTeamId(teamId domaincommon.BaseUUID) ([]UserEntity, error)
	FindByPairId(pairId domaincommon.BaseUUID) ([]UserEntity, error)
	FindByManyUserIds(userIds []domaincommon.BaseUUID) ([]UserEntity, error)
	Save(saveUesr UserEntity) error
	Exists(email domaincommon.Email) (bool, error)
}
