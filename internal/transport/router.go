package transport

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/itsmandrew/Central-Bank/internal/config"
	"github.com/itsmandrew/Central-Bank/internal/service"
)

func NewRouter(cfg *config.Config, db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/hello", service.HelloHandler)

	return r
}
