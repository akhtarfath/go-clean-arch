package main

import (
	"github.com/akhtarfath/gol-clean-arch-api/app/databases"
	_articleHttpDelivery "github.com/akhtarfath/gol-clean-arch-api/features/article/delivery/http"
	_articleRepo "github.com/akhtarfath/gol-clean-arch-api/features/article/repository/mysql"
	_articleUcase "github.com/akhtarfath/gol-clean-arch-api/features/article/usecase"
	_authorRepo "github.com/akhtarfath/gol-clean-arch-api/features/author/repository/mysql"
	"github.com/labstack/echo"
	"time"
)

func Routes(e *echo.Echo, timeoutContext time.Duration) {
	// Database injection from go wire
	db := databases.InitializedDatabaseRepository()

	// Repository injection to use a case
	articleRepository := _articleRepo.NewMysqlArticleRepository(db.DatabaseMySQL.GetDatabaseConnection())
	authorRepo := _authorRepo.NewMysqlAuthorRepository(db.DatabaseMySQL.GetDatabaseConnection())

	// Use case injection
	articleUseCase := _articleUcase.NewArticleUseCase(articleRepository, authorRepo, timeoutContext)

	// Delivery injection (HTTP)
	g := e.Group("/api/v1")
	_articleHttpDelivery.NewArticleHandler(g, articleUseCase)
}
