package data

import (
	"net/http"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet()

type repo struct {
	client *http.Client
	db     *gorm.DB
}

func NewRepo(client *http.Client, db *gorm.DB) *repo {
	return &repo{}
}
