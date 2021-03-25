# GinGolang-FeedReaderApi
this is a simple back-end with auth system and feed reader \
1. clone the proeject
2. config database connection info in 'app/Database.go'
3. run main.go (run project) 
4. after you run the project, 3 tables will be created in your database, open your database go to table feeds and add new feed source 
* name -> name of the feed source for example isna
* url -> feed url for example 'https://www.isna.ir/rss'
* content_selector -> css selector of article body for example '#item > div.full-news > div.full-news-text'

*done*


####
the crawler engine go and get contents like news or etc and update 'contents' table
