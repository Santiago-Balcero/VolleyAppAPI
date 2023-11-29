package ports

import (
	"volleyapp/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type SetRepository interface {
	SaveNewSet(domain.SetMainInfo) (int, error)
	FinishSet(int, domain.SetMainInfo) (int, error)
}

type SetService interface {
	CreateSet(domain.SetMainInfo) (int, error)
	FinishSet(int) (int, error)
}

type SetController interface {
	CreateSet(*gin.Context)
	FinishSet(*gin.Context)
	InitSetRoutes()
}
