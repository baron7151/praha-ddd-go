package domainteam

import domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"

type TeamId struct {
	domaincommon.UUIDProvider
}

func NewTeamId(value string) (TeamId, error) {
	baseUUID, err := domaincommon.NewBaseUUID(value)
	if err != nil {
		return TeamId{}, err
	}
	return TeamId{UUIDProvider: baseUUID}, nil
}
