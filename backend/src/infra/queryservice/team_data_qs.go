package queryservice

import (
	appteam "github.com/baron7151/praha-ddd-go/src/app/team"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"gorm.io/gorm"
)

type TeamDataQS struct {
	db *gorm.DB
	appteam.ITeamDataQS
}

func NewTeamDataQS(db *gorm.DB) *TeamDataQS {
	return &TeamDataQS{
		db:          db,
		ITeamDataQS: &TeamDataQS{}, // 自身の型で埋め込み
	}
}

func (t *TeamDataQS) FindAllTeams() ([]appteam.TeamDataDTO, error) {
	var teams []infra.Team
	result := t.db.Preload("Users").Preload("Pairs").Find(&teams)
	if result != nil {
		return []appteam.TeamDataDTO{}, result.Error
	}

	var teamDataDTOs []appteam.TeamDataDTO
	for _, team := range teams {
		var userIds []string
		for _, user := range team.Users {
			userIds = append(userIds, user.UserId)
		}
		var pairIds []string
		for _, pair := range team.Pairs {
			pairIds = append(pairIds, pair.PairId)
		}
		teamDataDTO := appteam.NewTeamDataDTO(team.TeamId, team.TeamName, &userIds, &pairIds)
		teamDataDTOs = append(teamDataDTOs, teamDataDTO)
	}
	return teamDataDTOs, nil
}

func (t *TeamDataQS) FindByTeamName(teamname string) (appteam.TeamDataDTO, error) {
	var team infra.Team
	result := t.db.Where("team_name = ?", teamname).Preload("Users").Preload("Pairs").First(&team)
	if result.Error != nil {
		return appteam.TeamDataDTO{}, result.Error
	}
	var userIds []string
	for _, user := range team.Users {
		userIds = append(userIds, user.UserId)
	}
	var pairIds []string
	for _, pair := range team.Pairs {
		pairIds = append(pairIds, pair.PairId)
	}
	teamDataDTO := appteam.NewTeamDataDTO(team.TeamId, team.TeamName, &userIds, &pairIds)
	return teamDataDTO, nil
}
