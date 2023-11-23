package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	server := &http.Server{Addr: serverAddr, Handler: ctrl}

	go handleSignals(server)

	log.Print(server.ListenAndServe())
}

func handleSignals(server *http.Server) {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs

	log.Print("handling signal: ", sig)

	log.Print(server.Shutdown(context.Background()))
}

func mustNotError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
