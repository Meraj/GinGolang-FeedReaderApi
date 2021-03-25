package Models

import (
	".."
)
type ContentClass struct {

}
type Content struct {
	Id int
	Title string
	Url string
	Body string
	SourceId int
}
func (ContentClass) All() []Content {
	db := app.Database()
	var contents []Content
	rows, err := db.Query("SELECT id,title,url,body,source_id FROM contents")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for rows.Next(){
		var content Content
		rows.Scan(&content.Id,&content.Title,&content.Url,&content.Body,&content.SourceId)
		contents = append(contents,content)
	}
	return contents
}
