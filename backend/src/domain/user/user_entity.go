package domainuser

import (
	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainpair "github.com/baron7151/praha-ddd-go/src/domain/pair"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
)

type UserStatus string

const (
	ACTIVE   UserStatus = "ACTIVE"
	INACTIVE UserStatus = "INACTIVE"
	DELETE   UserStatus = "DELETE"
)

func StringToUserStatus(statusStr string) (UserStatus, error) {
	switch statusStr {
	case string(ACTIVE):
		return ACTIVE, nil
	case string(INACTIVE):
		return INACTIVE, nil
	case string(DELETE):
		return DELETE, nil
	default:
		return "", domaincommon.NewDomainError("UserStatus is invalid")
	}
}

type UserId struct {
	domaincommon.UUIDProvider
}

func NewUserId(value string) (UserId, error) {
	baseUUID, err := domaincommon.NewBaseUUID(value)
	if err != nil {
		return UserId{}, err
	}
	return UserId{UUIDProvider: baseUUID}, nil
}

type UserName struct {
	value string
}

func NewUserName(value string) (UserName, error) {
	if len(value) == 0 {
		return UserName{}, domaincommon.NewDomainError("User name is invalid")
	}
	return UserName{value: value}, nil
}
func (u UserName) GetValue() string {
	return u.value
}
func (u *UserName) Equals(other UserName) bool {
	return u.value == other.value
}

type UserEntity struct {
	userId   UserId
	userName UserName
	email    domaincommon.Email
	status   UserStatus
	pairId   *domainpair.PairId
	teamId   *domainteam.TeamId
}

type UserProperties struct {
	UserId     UserId
	UserName   UserName
	Email      domaincommon.Email
	UserStatus UserStatus
	PairId     *domainpair.PairId
	TeamId     *domainteam.TeamId
}

type UserOption func(*UserEntity)

func WithPairId(pairId *domainpair.PairId) UserOption {
	return func(u *UserEntity) {
		u.pairId = pairId
	}
}

func WithTeamId(teamId *domainteam.TeamId) UserOption {
	return func(u *UserEntity) {
		u.teamId = teamId
	}
}

func NewUserEntity(
	userId UserId,
	userName UserName,
	email domaincommon.Email,
	status UserStatus,
	options ...UserOption,
) UserEntity {
	userEntity := UserEntity{
		userId:   userId,
		userName: userName,
		email:    email,
		status:   status,
	}

	for _, option := range options {
		option(&userEntity)
	}

	return userEntity
}

func (u UserEntity) GetUserId() UserId {
	return u.userId
}

func (u UserEntity) GetUserName() UserName {
	return u.userName
}

func (u UserEntity) GetEmail() domaincommon.Email {
	return u.email
}

func (u UserEntity) GetUserStatus() UserStatus {
	return u.status
}

func (u UserEntity) GetPairId() *domainpair.PairId {
	return u.pairId
}

func (u UserEntity) GetTeamId() *domainteam.TeamId {
	return u.teamId
}

func (u UserEntity) GetAllProperties() UserProperties {
	return UserProperties{
		UserId:     u.userId,
		UserName:   u.userName,
		Email:      u.email,
		UserStatus: u.status,
		PairId:     u.pairId,
		TeamId:     u.teamId,
	}
}

func (u UserEntity) Equals(other UserEntity) bool {
	return u.userId.GetValue() == other.userId.GetValue()
}

func (u UserEntity) ChangeStatus(status UserStatus) (UserEntity, error) {
	switch status {
	case ACTIVE:
		return NewUserEntity(u.userId, u.userName, u.email, ACTIVE, WithPairId(u.pairId), WithTeamId(u.teamId)), nil
	case INACTIVE:
		return NewUserEntity(u.userId, u.userName, u.email, INACTIVE), nil
	case DELETE:
		return NewUserEntity(u.userId, u.userName, u.email, DELETE), nil
	}
	return UserEntity{}, domaincommon.NewDomainError("Status is invalid")
}

func (u UserEntity) ChangeUserName(userName UserName) (UserEntity, error) {
	return NewUserEntity(u.userId, userName, u.email, u.status, WithPairId(u.pairId), WithTeamId(u.teamId)), nil
}

func (u UserEntity) ChangeEmail(email domaincommon.Email) (UserEntity, error) {
	return NewUserEntity(u.userId, u.userName, email, u.status, WithPairId(u.pairId), WithTeamId(u.teamId)), nil
}

func (u UserEntity) ChangePairId(pairId domainpair.PairId) (UserEntity, error) {
	return NewUserEntity(u.userId, u.userName, u.email, u.status, WithPairId(&pairId), WithTeamId(u.teamId)), nil
}

func (u UserEntity) ChangeTeamId(teamId domainteam.TeamId) (UserEntity, error) {
	return NewUserEntity(u.userId, u.userName, u.email, u.status, WithPairId(u.pairId), WithTeamId(&teamId)), nil
}
