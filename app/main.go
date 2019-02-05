package main

import (
	"database"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func createDocumentHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newDocument database.Document
	if err := decoder.Decode(&newDocument); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	database.DB.Create(&newDocument)
	res, err := json.Marshal(newDocument)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}

func showDocumentHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var document database.Document
	database.DB.Where("ID = ?", ps.ByName("DocID")).First(&document)
	res, err := json.Marshal(document)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write(res)
}

func main() {
	//  urls := urlserver.GenerateURLs()
	// 	documents := crawler.Scrape(urls)
	// 	storeserver.save(documents)
	//  pages := indexer.get(documents)
	//  hits, anchors := indexer.parse(pages)

	router := httprouter.New()
	router.POST("/posts", createDocumentHandler)
	router.GET("/posts/:postId", showDocumentHandler)

	// add database
	_, err := database.Init()
	if err != nil {
		log.Println("connection to DB failed, aborting...")
		log.Fatal(err)
	}
	defer database.DB.Close()
	log.Println("connected to DB")

	// print env
	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.Println("Running api server in production mode")
	} else {
		log.Println("Running api server in dev mode")
	}

	http.ListenAndServe(":8080", router)
}
