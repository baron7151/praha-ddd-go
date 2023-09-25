package domainteam

import (
	"regexp"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
)

const (
	MIN_TEAM_USER = 3
)

type TeamName struct {
	value string
}

func NewTeamName(value string) (TeamName, error) {
	var TeamNamePattern = regexp.MustCompile(`^[0-9]+$`)
	if !TeamNamePattern.MatchString(value) {
		return TeamName{}, domaincommon.NewDomainError("team name is invalid " + value)
	} else if len(value) > 3 || len(value) < 1 {
		return TeamName{}, domaincommon.NewDomainError("team name is invalid " + value)
	} else {
		return TeamName{value: value}, nil
	}
}

func (t TeamName) GetValue() string {
	return t.value
}

type TeamEntity struct {
	teamId   domaincommon.BaseUUID
	teamName TeamName
	pairIds  *[]domaincommon.BaseUUID
	userIds  *[]domaincommon.BaseUUID
}
type TeamOption func(*TeamEntity)

func WithPairIds(pairIds *[]domaincommon.BaseUUID) TeamOption {
	return func(u *TeamEntity) {
		u.pairIds = pairIds
	}
}

func WithUserIds(userIds *[]domaincommon.BaseUUID) TeamOption {
	return func(u *TeamEntity) {
		u.userIds = userIds
	}
}

func NewTeamEntity(
	teamId domaincommon.BaseUUID,
	teamName TeamName,
	options ...TeamOption,
) TeamEntity {
	teamEntity := TeamEntity{
		teamId:   teamId,
		teamName: teamName,
	}
	for _, option := range options {
		option(&teamEntity)
	}
	return teamEntity
}

func (t TeamEntity) GetTeamId() domaincommon.BaseUUID {
	return t.teamId
}

func (t TeamEntity) GetTeamName() TeamName {
	return t.teamName
}

func (t TeamEntity) GetPairIds() *[]domaincommon.BaseUUID {
	return t.pairIds
}

func (t TeamEntity) GetUserIds() *[]domaincommon.BaseUUID {
	return t.userIds
}

func (t TeamEntity) Equals(other TeamEntity) bool {
	return t.teamId == other.teamId
}

type TeamProperties struct {
	TeamId   domaincommon.BaseUUID
	TeamName TeamName
	PairIds  *[]domaincommon.BaseUUID
	UserIds  *[]domaincommon.BaseUUID
}

func (t TeamEntity) GetAllProperties() TeamProperties {
	return TeamProperties{
		TeamId:   t.teamId,
		TeamName: t.teamName,
		PairIds:  t.pairIds,
		UserIds:  t.userIds,
	}
}

func ValidateTeamUserCount(userIds *[]domaincommon.BaseUUID) bool {
	if userIds == nil {
		return false
	} else if len(*userIds) >= MIN_TEAM_USER {
		return true
	} else {
		return false
	}
}

func (t TeamEntity) CountTeamUser() int {
	if t.userIds == nil {
		return 0
	} else {
		return len(*t.userIds)
	}
}

func (t TeamEntity) IsMinimumTeamUser() bool {
	if t.CountTeamUser() > MIN_TEAM_USER {
		return false
	} else {
		return true
	}
}

func (t TeamEntity) ChangeTeamName(teamName TeamName) TeamEntity {
	return NewTeamEntity(t.teamId, teamName, WithPairIds(t.pairIds), WithUserIds(t.userIds))
}
