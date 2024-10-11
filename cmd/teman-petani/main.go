package main

import (
	"os"
	"os/signal"

	"github.com/kurnhyalcantara/teman-petani/config"
	"github.com/kurnhyalcantara/teman-petani/libs/database"
)

var (
	appConfig *config.Config
	dbSql     *database.DB
)

func init() {
	appConfig = config.InitConfig()
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	startDBConnections()

	<-ch

	closeDBConnections()

}
