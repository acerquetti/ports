package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/acerquetti/ports/app"
	"github.com/acerquetti/ports/infra/controller"
	"github.com/acerquetti/ports/infra/database"
)

var (
	portsPath  string
	serverAddr string
)

func init() {
	flag.StringVar(&portsPath, "portsPath", "./ports.json", "Path pointing ports.json file")
	flag.StringVar(&serverAddr, "serverAddr", ":8080", `Server address (e.g. ":8080")`)
	flag.Parse()
}

func main() {
	var err error

	portsFile, err := os.Open(portsPath)
	mustNotError(err)

	db, err := database.NewMemoryDB(portsFile)
	mustNotError(err)

	err = portsFile.Close()
	mustNotError(err)

	svc := app.NewService(db)
	ctrl := controller.NewREST(svc)

	log.Fatal(http.ListenAndServe(serverAddr, ctrl))
}

func mustNotError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
