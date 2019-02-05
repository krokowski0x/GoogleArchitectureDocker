package storeserver

import (
	"net/http"
)

func SaveToDB(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//urls := urlserver.GenerateURLs()
	//fmt.Println(urls)
	// 	document := crawler.Scrape(urls)
	// 	storeserver.SaveToDB(document)

	// type Document struct {
	// 	gorm.Model
	// 	docID      int
	// 	ecode      int
	// 	urlLength  int
	// 	pageLength int
	// 	url        string
	// 	page       string
	// }

	// 	var b bytes.Buffer
	// w := zlib.NewWriter(&b)
	// w.Write([]byte("hello, world\n"))
	// w.Close()
}
