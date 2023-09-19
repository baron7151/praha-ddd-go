package appuser

type GetUserDataUseCase struct {
	userDataQS IUserDataQS
}

func NewGetUserDataUsecase(userDataQS IUserDataQS) *GetUserDataUseCase {
	return &GetUserDataUseCase{
		userDataQS: userDataQS,
	}
}

func (u *GetUserDataUseCase) GetUserData(username string) ([]UserDataDTO, error) {
	if username == "" {
		result, err := u.userDataQS.FindAllUsers()
		if err != nil {
			return []UserDataDTO{}, err
		}
		return result, err
	} else {
		result, err := u.userDataQS.FindByUserName(username)
		if err != nil {
			return []UserDataDTO{}, err
		}
		return result, err
	}
}
