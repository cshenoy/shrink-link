package main

import (
	"encoding/json"
	// "fmt"
	shrinklink "github.com/cshenoy/shrinklink/shrinklink"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Error struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

func createHandler(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var link shrinklink.UrlPayload
	log.Println(request.Body == nil)
	// THIS IS WRONG. DO NOT USE IN PROD
	if request.ContentLength == 0 {
		// http.Error(responseWriter, "Please send a request body", 400)
		thisErr := Error{400, "Body missing or malformed"}
		// http.Error(w, err, http.StatusBadRequest)
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(thisErr.Code)
		json.NewEncoder(responseWriter).Encode(thisErr)
		// http.Error(responseWriter, thisErr.Message, http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&link)
	log.Println(err, err == io.EOF)
	if err != nil {
		http.Error(responseWriter, err.Error(), 400)
		return
	}
	log.Println(link)
}

func findHandler(responseWriter http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	// log.Println("sup dawg", ps.ByName("hash"))
	// fmt.Fprintf(responseWriter, "%s", ps.ByName("hash"))
	// CheckHash
	hash := ps.ByName("hash")
	// If good -> redirect to long_url
	if strings.EqualFold(hash, "blanc") {
		http.Redirect(responseWriter, request, "http://corp.urbanstems.com/blanc", 301)
	} else {
		http.Redirect(responseWriter, request, "https://urbanstems.com", 301)
	}
	// If bad -> redirect to urbanstems.com
}

func indexHandler(responseWriter http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	http.Redirect(responseWriter, request, "https://urbanstems.com", 301)
}

func validToken(token string) bool {
	// return token == slashCommandToken()
	return true
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := httprouter.New()
	router.GET("/", indexHandler) // redirect to urbanstems.com
	router.GET("/:hash", findHandler)
	router.POST("/api/url", createHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
