package teamrepository

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
	domainteam.ITeamRepository
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		db:              db,
		ITeamRepository: &TeamRepository{},
	}
}
func (r *TeamRepository) FindByTeamId(teamId domaincommon.BaseUUID) (domainteam.TeamEntity, error) {
	var team infra.Team
	result := r.db.Where("team_id = ?", teamId.GetValue()).Preload("Users").Preload("Pairs").First(&team)
	if result.Error != nil {
		return domainteam.TeamEntity{}, result.Error
	}
	var userIds []string
	for _, user := range team.Users {
		userIds = append(userIds, user.UserId)
	}
	var pairIds []string
	for _, pair := range team.Pairs {
		pairIds = append(pairIds, pair.PairId)
	}
	teamEntity, err := domainteam.Create(team.TeamId, team.TeamName, &userIds, &pairIds)
	if err != nil {
		return domainteam.TeamEntity{}, err
	}
	return teamEntity, nil
}

func (r *TeamRepository) Save(saveTeam domainteam.TeamEntity) error {
	var team infra.Team
	newTeam := infra.Team{
		TeamId:   saveTeam.GetTeamId().GetValue(),
		TeamName: saveTeam.GetTeamName().GetValue(),
	}
	tx := r.db.Begin()
	result := tx.Where("team_id = ?", newTeam.TeamId).Find(&team)
	if result.Error != nil {
		return result.Error
	}
	if team.Id == 0 {
		result := tx.Create(&newTeam)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	} else {
		tx.Model(&team).Where("team_id = ?", newTeam.TeamId).Omit("team_id").Updates(newTeam)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	// result := r.db.Save(&team)
	// if result.Error != nil {
	// 	return result.Error
	// }
	return nil
}

func (r *TeamRepository) Exists(teamName domainteam.TeamName) (bool, error) {
	var teams []infra.Team
	result := r.db.Where("team_name = ?", teamName.GetValue()).Find(&teams)
	if result.Error != nil {
		return false, result.Error
	} else if result.RowsAffected == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (r *TeamRepository) FindAllTeams() ([]domainteam.TeamEntity, error) {
	var teams []infra.Team
	result := r.db.Preload("Users").Preload("Pairs").Find(&teams)
	if result.Error != nil {
		return []domainteam.TeamEntity{}, result.Error
	}
	var teamEntities []domainteam.TeamEntity
	for _, team := range teams {
		var userIds []string
		for _, user := range team.Users {
			userIds = append(userIds, user.UserId)
		}
		var pairIds []string
		for _, pair := range team.Pairs {
			pairIds = append(pairIds, pair.PairId)

		}
		teamEntity, err := domainteam.Create(team.TeamId, team.TeamName, &userIds, &pairIds)
		if err == nil {
			teamEntities = append(teamEntities, teamEntity)
		}
	}
	return teamEntities, nil
}
