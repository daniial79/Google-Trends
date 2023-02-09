package main

import "encoding/xml"

type RSS struct {
	XMLname xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title    string `xml:"title"`
	Link     string `xml:"link"`
	Traffic  string `xml:"approx_traffic"`
	NewsItem []News `xml:"news_item"`
}

type News struct {
	HeadLine     string `xml:"news_item_title"`
	HeadLineLink string `xml:"news_item_url"`
}
