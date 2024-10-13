package main

import (
	"os"
	"os/signal"

	"github.com/kurnhyalcantara/teman-petani/config"
	"github.com/kurnhyalcantara/teman-petani/libs/database"
	"github.com/kurnhyalcantara/teman-petani/libs/log"
	"github.com/rs/zerolog"
)

var (
	appConfig *config.Config
	dbSql     *database.DB
	logger    zerolog.Logger
)

func init() {
	appConfig = config.InitConfig()
	logger = log.SetupZerolog(appConfig)
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	startDBConnections()

	<-ch

	closeDBConnections()

}
