package main

import (
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func init() {
	// Set dbEnv file
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// Set debug mode
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// Middleware injection
	echo := middleware()
	// Set timeout for context, this timeout will be applied to every use case request to a database
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// Routes injection
	Routes(echo, timeoutContext)
	// Start server
	server(echo)
}
