package appuser

import (
	"testing"

	testutilcommon "github.com/baron7151/praha-ddd-go/src/testutil/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserData_FindAllUsers(t *testing.T) {
	// (1) モックを呼び出すための Controller を生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// (2) モックの生成
	pr := NewMockIUserDataQS(ctrl)

	// (3) テストに呼ばれるべきメソッドと引数・戻り値を指定
	userId1 := testutilcommon.GenerateUUID()
	userId2 := testutilcommon.GenerateUUID()
	pairId := testutilcommon.GenerateUUID()
	teamId := testutilcommon.GenerateUUID()
	expect := []UserDataDTO{
		{
			UserId:     userId1,
			UserName:   "test1",
			Email:      "test1@example.com",
			UserStatus: "ACTIVE",
			PairId:     &pairId,
			TeamId:     &teamId,
		},
		{
			UserId:     userId2,
			UserName:   "test2",
			Email:      "test2@example.com",
			UserStatus: "ACTIVE",
			PairId:     &pairId,
			TeamId:     &teamId,
		},
	}
	pr.EXPECT().FindAllUsers().Return(expect, nil)

	// (4) テストの本体
	result, err := NewGetUserDataUsecase(pr).GetUserData("")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, expect, result)
}

func TestGetUserData_FindByUserName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	pr := NewMockIUserDataQS(ctrl)

	// (3) テストに呼ばれるべきメソッドと引数・戻り値を指定
	userId1 := testutilcommon.GenerateUUID()
	pairId := testutilcommon.GenerateUUID()
	teamId := testutilcommon.GenerateUUID()
	expect := []UserDataDTO{
		{
			UserId:     userId1,
			UserName:   "test1",
			Email:      "test1@example.com",
			UserStatus: "ACTIVE",
			PairId:     &pairId,
			TeamId:     &teamId,
		},
	}

	pr.EXPECT().FindByUserName("test1").Return(expect, nil)

	// (4) テストの本体
	result, err := NewGetUserDataUsecase(pr).GetUserData("test1")
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, expect, result)
}
