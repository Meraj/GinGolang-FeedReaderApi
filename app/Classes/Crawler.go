package Classes

import (
	".."
	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"log"
	"net/http"
)
type Feeds struct {
	id int
	Name            string
	Url             string
	ContentSelector string
}
type FeedClass struct {

}
func (FeedClass) All() []Feeds {
	db := app.Database()
	rows, err := db.Query("SELECT id,name,url,content_selector FROM feeds")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	var allFeeds []Feeds
	for rows.Next(){
		var feed Feeds
		rows.Scan(&feed.id,&feed.Name,&feed.Url,&feed.ContentSelector)
		allFeeds = append(allFeeds,feed)
	}
	return allFeeds
}

func GetSources(){
	var feedClass FeedClass

	feeds := feedClass.All()
	for i := range feeds {
		feed := feeds[i]
		fp := gofeed.NewParser()
		fp.UserAgent = "Latest News 1.0"
		rssFeed, _ := fp.ParseURL(feed.Url)
		db := app.Database()
		for i2 := range rssFeed.Items {
			title := rssFeed.Items[i2].Title

			url := rssFeed.Items[i2].Link
			var total int
				err := db.QueryRow("SELECT COUNT(*) AS total FROM contents WHERE url =?", url).Scan(&total)
				if err != nil {
					panic(err.Error()) // proper error handling instead of panic in your app
				}
				if total == 0 {
					Crawl(title,url,feed)
				}
		}
	}
}
func Crawl(title string,url string,feed Feeds){
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	db := app.Database()
	bod, _ := doc.Find(feed.ContentSelector).Html()
	db.Exec("INSERT INTO contents (url,title,body,source_id) VALUES (?,?,?,?)",url,title,bod,feed.id)
}
