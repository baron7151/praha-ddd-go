package teamrepository

import (
	"os"
	"testing"

	domaincommon "github.com/baron7151/praha-ddd-go/src/domain/common"
	domainteam "github.com/baron7151/praha-ddd-go/src/domain/team"
	"github.com/baron7151/praha-ddd-go/src/infra"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../../../.test.env")
	if err != nil {
		panic(err)
	}
	status := m.Run()

	os.Exit(status)
}

func TestFindByTeamId(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	teamRepository := NewTeamRepository(db)
	var team infra.Team
	result := db.Where("id=1").Find(&team)
	if result.Error != nil {
		t.Fatal(result.Error)
	}

	teamId, _ := domaincommon.NewBaseUUID(team.TeamId)
	// 実行
	res, err := teamRepository.FindByTeamId(teamId)
	if err != nil {
		t.Fatal(err)
	}
	// 検証

	assert.Equal(t, team.TeamId, res.GetTeamId().GetValue())
}

func TestExists(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	teamRepository := NewTeamRepository(db)
	var team infra.Team
	result := db.Where("id=1").Find(&team)
	if result.Error != nil {
		t.Fatal(result.Error)
	}

	teamName, _ := domainteam.NewTeamName(team.TeamName)
	// 実行
	res, err := teamRepository.Exists(teamName)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, true, res)
}

func TestFindAllTeams(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	teamRepository := NewTeamRepository(db)
	// 実行
	res, err := teamRepository.FindAllTeams()
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	assert.Equal(t, 2, len(res))
}

func TestSave_Create(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	teamRepository := NewTeamRepository(db)

	teamId, _ := domaincommon.NewBaseUUID("")
	teamName, _ := domainteam.NewTeamName("99")
	teamEntity := domainteam.NewTeamEntity(teamId, teamName)
	// 実行
	err := teamRepository.Save(teamEntity)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	res, _ := teamRepository.FindByTeamId(teamId)

	assert.Equal(t, teamEntity, res)

}

func TestSave_Update(t *testing.T) {
	infra.InitDB()
	db, _ := infra.ConnectDB()
	teamRepository := NewTeamRepository(db)
	var team infra.Team
	result := db.Where("id=1").Find(&team)
	if result.Error != nil {
		t.Fatal(result.Error)
	}
	teamId, _ := domaincommon.NewBaseUUID(team.TeamId)
	teamName, _ := domainteam.NewTeamName("99")
	teamEntity := domainteam.NewTeamEntity(teamId, teamName)
	// 実行
	err := teamRepository.Save(teamEntity)
	if err != nil {
		t.Fatal(err)
	}
	// 検証
	res, _ := teamRepository.FindByTeamId(teamId)

	assert.Equal(t, teamEntity.GetTeamName(), res.GetTeamName())

}
