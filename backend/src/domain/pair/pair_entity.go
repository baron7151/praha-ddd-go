package domainpair

import domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"

type PairId struct {
	domaincommon.UUIDProvider
}

func NewPairId(value string) (PairId, error) {
	baseUUID, err := domaincommon.NewBaseUUID(value)
	if err != nil {
		return PairId{}, err
	}
	return PairId{UUIDProvider: baseUUID}, nil
}
