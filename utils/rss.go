package utils

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	models "github.com/Triyaambak/RSS-Aggregator/models"
)

func urlToFeed(url string) (models.RSSFEED, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return models.RSSFEED{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.RSSFEED{}, err
	}
	rssFeed := models.RSSFEED{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return models.RSSFEED{}, err
	}

	return rssFeed, nil
}
