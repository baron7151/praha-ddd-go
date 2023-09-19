package userrepository

import (
	"errors"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainpair "github.com/baron7151/praha-ddd-go/src/domain/pair"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
	domainuser "github.com/baron7151/praha-ddd-go/src/domain/user"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
	domainuser.IUserRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:              db,
		IUserRepository: &UserRepository{}, // 自身の型で埋め込み
	}
}

func (r *UserRepository) FindByUserId(userId domainuser.UserId) (domainuser.UserEntity, error) {

	var user infra.User
	result := r.db.Where("user_id = ?", userId.GetValue()).First(&user)

	if result.Error != nil {
		return domainuser.UserEntity{}, result.Error
	}

	//ToDo resultをUserEntityに変換する
	userEntity, err := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)

	if err != nil {
		return domainuser.UserEntity{}, err
	}

	return userEntity, nil
}

func (r *UserRepository) FindByEmail(email domaincommon.Email) (domainuser.UserEntity, error) {
	var user infra.User
	result := r.db.Where("email = ?", email.GetValue()).First(&user)
	if result.Error != nil {
		return domainuser.UserEntity{}, result.Error
	}
	userEntity, err := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
	if err != nil {
		return domainuser.UserEntity{}, err
	}
	return userEntity, nil
}

func (r *UserRepository) FindByTeamId(teamId domainteam.TeamId) ([]domainuser.UserEntity, error) {
	var users []infra.User
	result := r.db.Where("team_id = ?", teamId.GetValue()).Find(&users)
	if result.Error != nil {
		return []domainuser.UserEntity{}, result.Error
	}
	var userEntities []domainuser.UserEntity
	for _, user := range users {
		userEntity, err := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
		if err == nil {
			userEntities = append(userEntities, userEntity)
		}
	}
	return userEntities, nil
}

func (r *UserRepository) FindByPairId(pairId domainpair.PairId) ([]domainuser.UserEntity, error) {
	var users []infra.User
	result := r.db.Where("pair_id = ?", pairId.GetValue()).Find(&users)
	if result.Error != nil {
		return []domainuser.UserEntity{}, result.Error
	}
	var userEntities []domainuser.UserEntity
	for _, user := range users {
		userEntity, err := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
		if err == nil {
			userEntities = append(userEntities, userEntity)
		}
	}
	return userEntities, nil
}

func (r *UserRepository) FindByManyUserIds(userIds []domainuser.UserId) ([]domainuser.UserEntity, error) {
	var users []infra.User
	var userIdsStr []string
	for _, userId := range userIds {
		userIdsStr = append(userIdsStr, userId.GetValue())
	}
	result := r.db.Where("user_id IN ?", userIdsStr).Find(&users)
	if result.Error != nil {
		return []domainuser.UserEntity{}, result.Error
	}
	var userEntities []domainuser.UserEntity
	for _, user := range users {
		userEntity, err := domainuser.Create(user.UserId, user.UserName, user.Email, user.UserStatus, user.PairId, user.TeamId)
		if err == nil {
			userEntities = append(userEntities, userEntity)
		}
	}
	return userEntities, nil
}

func (r *UserRepository) Save(saveUser domainuser.UserEntity) error {
	var user infra.User
	userProserties := saveUser.GetAllProperties()
	newUser := infra.User{
		UserId:     userProserties.UserId.GetValue(),
		UserName:   userProserties.UserName.GetValue(),
		Email:      userProserties.Email.GetValue(),
		UserStatus: string(userProserties.UserStatus),
	}
	if userProserties.PairId != nil {
		PairID := userProserties.PairId.GetValue()
		newUser.PairId = &PairID
	}
	if userProserties.TeamId != nil {
		TeamID := userProserties.TeamId.GetValue()
		newUser.TeamId = &TeamID
	}
	tx := r.db.Begin()
	result := tx.Where("user_id = ?", saveUser.GetUserId().GetValue()).First(&user)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	if user.Id == 0 {
		result := tx.Create(&newUser)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	} else {
		tx.Model(&infra.User{}).Where("user_id = ?", userProserties.UserId.GetValue()).Omit("user_id").Updates(newUser)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	tx.Commit()
	return nil
}

func (r *UserRepository) Exists(email domaincommon.Email) (bool, error) {
	// メールアドレスが存在するかを確認するコードをここに記述
	var user infra.User
	result := r.db.Where("email = ?", email.GetValue()).First(&user)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	} else if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	} else {
		return true, nil
	}
}
