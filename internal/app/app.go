package app

import (
	"github.com/dan-nicholls/project-sourdough/internal/db"
)

type AppService struct {
	Db db.Database
	FormOptions []Option
	AccessCode string 
	TokenStore TokenStore
}

func New(db db.Database, code string) *AppService {
	store := NewDemoTokenStore()
	return &AppService{
		Db: db,
		FormOptions: FormOptions,
		AccessCode: code,
		TokenStore: &store,
	}
}
