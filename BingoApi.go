package BingoApi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://api.bing.microsoft.com/v7.0/"

// This struct formats the answer provided by the Bing News Search API.
type NewsAnswer struct {
	ReadLink     string `json:"readLink"`
	QueryContext struct {
		OriginalQuery string `json:"originalQuery"`
		AdultIntent   bool   `json:"adultIntent"`
	} `json:"queryContext"`
	TotalEstimatedMatches int `json:"totalEstimatedMatches"`
	Sort                  []struct {
		Name       string `json:"name"`
		ID         string `json:"id"`
		IsSelected bool   `json:"isSelected"`
		URL        string `json:"url"`
	} `json:"sort"`
	Value []struct {
		Name  string `json:"name"`
		URL   string `json:"url"`
		Image struct {
			Thumbnail struct {
				ContentUrl string `json:"thumbnail"`
				Width      int    `json:"width"`
				Height     int    `json:"height"`
			} `json:"thumbnail"`
		} `json:"image"`
		Description string `json:"description"`
		Provider    []struct {
			Type string `json:"_type"`
			Name string `json:"name"`
		} `json:"provider"`
		DatePublished string `json:"datePublished"`
	} `json:"value"`
}

type reqClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type BingoApi struct {
	ClientKey string
	Client    reqClient
}

func New(ClientKey string) *BingoApi {
	return &BingoApi{ClientKey, http.DefaultClient}
}

type Args struct {
	Key   string
	Value string
}

func (b *BingoApi) NewsSearch(q string) (*NewsAnswer, error) {
	url := baseUrl + "news/search"

	req, err := http.NewRequest("GET", url, nil)
	var ans *NewsAnswer
	if req != nil && err == nil {

		param := req.URL.Query()
		req.Header.Add("Ocp-Apim-Subscription-Key", b.ClientKey)
		param.Add("q", q)
		param.Add("freshness", "Day")
		req.URL.RawQuery = param.Encode()
		res, err := b.Client.Do(req)
		if err != nil {
			return nil, err
		}
		if res.StatusCode >= 400 {
			return nil, errors.New(fmt.Sprintf("server error Status:%d", res.StatusCode))
		}
		// Close the connection.
		defer func() {
			_ = res.Body.Close()
		}()

		// Read the results
		body, err := ioutil.ReadAll(res.Body)
		if body != nil && err == nil {
			ans = new(NewsAnswer)
			err = json.Unmarshal(body, ans)
			if err != nil {
				return nil, err
			}
		}

	}
	return ans, err
}
