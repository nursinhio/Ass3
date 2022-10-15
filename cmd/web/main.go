package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)
import _ "github.com/go-sql-driver/mysql"

type config struct {
	addr      string
	staticDir string
	dsn       string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	file, err := os.OpenFile("Logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//logs
	infoLog := log.New(file, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(file, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{errorLog: errorLog, infoLog: infoLog}

	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	srv := &http.Server{
		Addr:     cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	srcError := srv.ListenAndServe()
	infoLog.Println("staring server in port %v", cfg.addr)
	if srcError != nil {
		srv.ErrorLog.Fatal(err)
	}

}

