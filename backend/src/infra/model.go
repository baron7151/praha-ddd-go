package infra

import "time"

type User struct {
	Id             uint   `gorm:"primaryKey;autoIncrement"`
	UserId         string `gorm:"unique;not null"`
	UserName       string
	Email          string `gorm:"unique;not null"`
	UserStatus     string
	PairId         *string
	TeamId         *string
	TaskProgresses []TaskProgress `gorm:"foreignKey:UserId;references:UserId;OnDelete:CASCADE;OnUpdate:CASCADE;"`
	Created_at     time.Time      `gorm:"autoCreateTime"`
	Updated_at     time.Time      `gorm:"autoUpdateTime"`
}

type Pair struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	PairId     string `gorm:"unique; not null"`
	PairName   string
	TeamId     *string
	Users      []User    `gorm:"foreignKey:PairId;references:PairId;OnDelete:SET NULL;OnUpdate:SET NULL;"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}

type Team struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	TeamId     string `gorm:"unique;not null"`
	TeamName   string
	Users      []User    `gorm:"foreignKey:TeamId;references:TeamId;OnDelete:SET NULL;OnUpdate:SET NULL;"`
	Pairs      []Pair    `gorm:"foreignKey:TeamId;references:TeamId;OnDelete:SET NULL;OnUpdate:SET NULL;"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}

type TaskProgress struct {
	Id             uint   `gorm:"primaryKey;autoIncrement"`
	TaskProgressId string `gorm:"unique;not null"`
	TaskStatus     string
	TaskId         string
	UserId         string
	Created_at     time.Time `gorm:"autoCreateTime"`
	Updated_at     time.Time `gorm:"autoUpdateTime"`
}

type Task struct {
	Id             uint   `gorm:"primaryKey;autoIncrement"`
	TaskId         string `gorm:"unique;not null"`
	TaskName       string
	TaskContent    string
	TaskCategory   string
	TaskProgresses []TaskProgress `gorm:"foreignKey:TaskId;references:TaskId;OnDelete:CASCADE;OnUpdate:CASCADE;"`
	Created_at     time.Time      `gorm:"autoCreateTime"`
	Updated_at     time.Time      `gorm:"autoUpdateTime"`
}
