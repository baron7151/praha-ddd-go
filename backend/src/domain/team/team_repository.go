package domainteam

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
)

type ITeamRepository interface {
	FindByTeamId(teamId domaincommon.BaseUUID) (TeamEntity, error)
	Save(team TeamEntity) error
	Exists(teamName TeamName) (bool, error)
	FindAllTeams() ([]TeamEntity, error)
}
