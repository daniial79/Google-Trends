package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	inputErrMsg = "error: please provide region name."

	getGoogleTrendsErrStat = iota + 1
	readGoogleTrendsErrStat
	unmarshalErrStat
	inputErrStat
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(inputErrMsg)
		os.Exit(inputErrStat)
	}

	region := os.Args[1]

	var (
		r    RSS
		data = readGoogleTrends(region)
	)

	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println(err)
		os.Exit(unmarshalErrStat)
	}

	items := r.Channel.Items

	fmt.Printf("%s Google Trends For Today!\n", region)
	fmt.Printf("####################################\n")

	for i, item := range items {
		fmt.Printf("Title %d: %s\n", i+1, item.Title)
		fmt.Println("Link: ", item.Link)
		fmt.Println("Traffic approx: ", item.Traffic)
		fmt.Printf("\nitem's news: \n\n")

		for _, news := range item.NewsItem {
			fmt.Println(news.HeadLine)
			fmt.Println(news.HeadLineLink)
			fmt.Println("------------------------------")
		}
		fmt.Printf("\n\n")
	}
}

func getFromGooleTrends(region string) *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=" + region)

	if err != nil {
		fmt.Println(err)
		os.Exit(getGoogleTrendsErrStat)
	}

	return resp
}

func readGoogleTrends(region string) []byte {
	resp := getFromGooleTrends(region)

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(readGoogleTrendsErrStat)
	}

	return data
}
