package appuser

type UserDataDTO struct {
	UserId     string
	UserName   string
	Email      string
	UserStatus string
	PairId     *string
	TeamId     *string
}

func NewUserDataDTO(userId string, userName string, email string, userStatus string, pairId *string, teamId *string) UserDataDTO {
	return UserDataDTO{
		UserId:     userId,
		UserName:   userName,
		Email:      email,
		UserStatus: userStatus,
		PairId:     pairId,
		TeamId:     teamId,
	}
}

type IUserDataQS interface {
	FindAllUsers() ([]UserDataDTO, error)
	FindByUserName(name string) ([]UserDataDTO, error)
}
