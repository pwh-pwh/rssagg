package scraper

import (
	"context"
	"encoding/xml"
	"github.com/pwh-pwh/rssagg/config"
	"github.com/pwh-pwh/rssagg/internal/database"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func StartScraping(feedNum int, duration time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", duration, feedNum)
	ticker := time.NewTicker(duration)
	for ; ; <-ticker.C {
		feeds, err := config.Config.DB.GetNextFeedsToFetch(context.Background(), int32(feedNum))
		if err != nil {
			log.Printf("Couldn't get next feeds to fetch%s\n", err)
			continue
		}
		log.Printf("get %v feed to fetch", len(feeds))
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(wg, feed)
		}
		wg.Wait()
		log.Println("end wait")
	}
}

func scrapeFeed(wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := config.Config.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("can not mark feed name:%v id:%v\n", feed.Name, feed.ID)
		return
	}
	feedData, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		log.Println("Found post", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}

func fetchFeed(feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := httpClient.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}
