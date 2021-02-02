package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"time"
)

const (
	port       string = ":5812"
	cert       string = "certs/fullchain.pem"
	key        string = "certs/privkey.pem"
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
	time     string
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
		logRequest(r)
		fmt.Fprintf(w, "Thanks for using BrendAPI!\n")
		fmt.Fprintf(w, "Use /fen for a random move.\n")
	})
	http.HandleFunc("/fen", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)

		m, _ := url.ParseQuery(r.URL.RawQuery)
		infoLogger.Printf("inputFEN: %s", m["fen"][0])

		cmd := exec.Command("/home/server/restapi/chess/move_generator/out", "-f", m["fen"][0])
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
			errorLogger.Fatal(err)
		}

		infoLogger.Printf("outputFEN: %s", out.String())
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, out.String())
	})
}

func logRequest(r *http.Request) {
	infoLogger.Printf("received request:")
	infoLogger.Print(reqDetails{
		r.Method,
		r.URL.Path,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Unix(),
		r.RemoteAddr,
	})
}

func main() {
	initLogger()
	infoLogger.Println(" #### Starting " + name)
	initHandlers()
	infoLogger.Println("handlers initialized")

	errorLogger.Fatal(http.ListenAndServeTLS(port, cert, key, nil))
}
