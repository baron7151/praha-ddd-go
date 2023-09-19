package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var UsersColumn []string = []string{"id", "user_id", "user_name", "email", "user_status", "pair_id", "team_id", "created_at", "updated_at"}
var PairsColumn []string = []string{"id", "pair_id", "pair_name", "team_id", "created_at", "updated_at"}
var TeamsColumn []string = []string{"id", "team_id", "team_name", "created_at", "updated_at"}
var TasksColumn []string = []string{"id", "task_id", "task_name", "task_content", "task_category", "created_at", "updated_at"}
var TaskProgressesColumn []string = []string{"id", "task_progress_id", "task_status", "user_id", "task_id", "created_at", "updated_at"}

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, _ := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}
