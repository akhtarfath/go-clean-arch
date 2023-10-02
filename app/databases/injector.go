//go:build wireinject
// +build wireinject

package databases

import "github.com/google/wire"

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMySQL,
		NewDatabasePgSQL,
		NewDatabaseMSSQL,
		NewDatabaseRepository,
	)
	return nil
}
