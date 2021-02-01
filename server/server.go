package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	port       string = ":5812"
	name       string = "BREND-API"
	logPath    string = "logs/output.log"
	configPath string = "config.json"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	movegenPath string
	options     int
)

// TODO: make a method to print this
// func (rd *reqDetails) rqPrint()
type reqDetails struct {
	method   string
	urlPath  string
	time     time.Time
	unixTime int64
	addr     string
}

func initLogger() {
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	infoLogger = log.New(file, "[INFO]  ", log.LstdFlags)
	errorLogger = log.New(file, "[ERROR] ", log.LstdFlags)
	log.SetOutput(file)
}

func initHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		infoLogger.Printf("received request:")
		infoLogger.Print(reqDetails{
			r.Method,
			r.URL.Path,
			time.Now(),
			time.Now().Unix(),
			r.RemoteAddr,
		})
		fmt.Fprintf(w, "Thanks for using BrendAPI!\n")
		fmt.Fprintf(w, "Use /fen for a random move.\n")
	})
	http.HandleFunc("/fen", func(w http.ResponseWriter, r *http.Request) {
	})
}

func main() {
	initLogger()
	infoLogger.Println(" #### Starting " + name)
	initHandlers()
	infoLogger.Println("handlers initialized")

	errorLogger.Fatal(http.ListenAndServe(port, nil))
}
