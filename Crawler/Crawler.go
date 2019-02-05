package crawler

import (
	"io/ioutil"
	"net/http"
)

func Scrape(URLs []string) []byte {
	// c := colly.NewCollector(
	// 	// colly.AllowedDomains(URLs...),
	// 	colly.MaxDepth(2),
	// 	colly.Async(true),
	// )

	// c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})

	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	link := e.Attr("href")
	// 	fmt.Println(link)
	// 	e.Request.Visit(link)
	// })

	// c.Visit(URLs[0])
	// c.Wait()
	resp, err := http.Get(URLs[0])
	if err != nil {
		// err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}
