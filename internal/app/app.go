package app

import (
	"github.com/dan-nicholls/project-sourdough/internal/db"
)

type AppService struct {
	Db db.Database
	Config BreadConfigurations
}

func New(db db.Database) *AppService {
	return &AppService{
		Db: db,
		Config: availableConfigs,
	}
}
