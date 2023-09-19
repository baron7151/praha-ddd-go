package queryservice

import (
	appuser "github.com/baron7151/praha-ddd-go/src/app/user"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"gorm.io/gorm"
)

type UserDataQS struct {
	db *gorm.DB
	appuser.IUserDataQS
}

func NewUserDataQS(db *gorm.DB) *UserDataQS {
	return &UserDataQS{
		db:          db,
		IUserDataQS: &UserDataQS{}, // 自身の型で埋め込み
	}
}

func (u *UserDataQS) FindAllUsers() ([]appuser.UserDataDTO, error) {
	var users []infra.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return []appuser.UserDataDTO{}, result.Error
	}
	var userDataDTOs []appuser.UserDataDTO
	for _, user := range users {
		userDataDTO := appuser.NewUserDataDTO(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
		userDataDTOs = append(userDataDTOs, userDataDTO)
	}
	return userDataDTOs, nil
}

func (u *UserDataQS) FindByUserName(name string) ([]appuser.UserDataDTO, error) {
	var users []infra.User
	result := u.db.Where("user_name = ?", name).Find(&users)
	if result.Error != nil {
		return []appuser.UserDataDTO{}, result.Error
	}
	var userDataDTOs []appuser.UserDataDTO
	for _, user := range users {
		userDataDTO := appuser.NewUserDataDTO(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
		userDataDTOs = append(userDataDTOs, userDataDTO)
	}
	return userDataDTOs, nil
}
