package services

import (
	"database/sql"
	"github.com/eif-courses/golab/utils"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

type Models struct {
	User         User
	JsonResponse utils.JsonResponse
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
