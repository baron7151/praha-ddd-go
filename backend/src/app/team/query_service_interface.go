package appteam

type TeamDataDTO struct {
	TeamId   string
	TeamName string
	PairIds  *[]string
	UserIds  *[]string
}

func NewTeamDataDTO(teamId string, teamName string, pairIds *[]string, userIds *[]string) TeamDataDTO {
	return TeamDataDTO{
		TeamId:   teamId,
		TeamName: teamName,
		PairIds:  pairIds,
		UserIds:  userIds,
	}
}

type ITeamDataQS interface {
	FindAllTeams() ([]TeamDataDTO, error)
	FindByTeamName(teamname string) (TeamDataDTO, error)
}
