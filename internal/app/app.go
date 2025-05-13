package app

import (
	"github.com/dan-nicholls/project-sourdough/internal/db"
)

type AppService struct {
	Db db.Database
}

func New(db db.Database) *AppService {
	return &AppService{
		Db: db,
	}
}
