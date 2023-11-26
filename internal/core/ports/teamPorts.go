package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type TeamRepository interface {
	SaveNewTeam(domain.TeamMainInfo) (int, error)
	GetUserTeams(int) ([]domain.TeamSummary, error)
	GetTeam(string) (domain.Team, error)
	CheckTeamExistence(string) (bool, error)
}

type TeamService interface {
	CreateTeam(domain.TeamMainInfo) (int, error)
	GetUserTeams(int) ([]domain.TeamSummary, error)
	GetTeam(string) (domain.Team, error)
	UpdateTeamInfo(domain.TeamMainInfo) (bool, error)
}

type TeamController interface {
	CreateTeam(c *gin.Context)
	GetUserTeams(c *gin.Context)
	GetTeam(c *gin.Context)
	UpdateTeamInfo(c *gin.Context)
	InitTeamRoutes()
}
