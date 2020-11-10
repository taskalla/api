package unsplash

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/graphql-go/graphql"
	"github.com/taskalla/api/pkg/env"
)

func RandomImage(p graphql.ResolveParams) (interface{}, error) {
	key := env.String("UNSPLASH_ACCESS_KEY", "")
	if key == "" {
		return nil, errors.New("No Unsplash access key provided")
	}

	random_url := &url.URL{
		Scheme: "https",
		Host:   "api.unsplash.com",
		Path:   "/photos/random",
		RawQuery: url.Values{
			//"content_filter": {"high"},
			"collections": {"1026767"},
			"orientation": {"landscape"},
		}.Encode(),
	}

	req := &http.Request{
		Method: "GET",
		URL:    random_url,
		Header: http.Header{
			"Authorization": {fmt.Sprintf("Client-ID %s", key)},
		},
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if !(res.StatusCode >= 200 && res.StatusCode <= 299) {
		return nil, errors.New("Error getting image")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var parsedBody struct {
		URLs struct {
			Full    string `json:"full"`
			Regular string `json:"regular"`
		} `json:"urls"`
		BlurHash string `json:"blur_hash"`
		User     struct {
			Name  string `json:"name"`
			Links struct {
				HTML string `json:"html"`
			} `json:"links"`
		} `json:"user"`
		Links struct {
			HTML string `json:"html"`
		} `json:"links"`
	}

	err = json.Unmarshal(body, &parsedBody)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"url":         parsedBody.URLs.Full,
		"url_full":    parsedBody.URLs.Full,
		"url_regular": parsedBody.URLs.Regular,
		"meta_url":    parsedBody.Links.HTML,
		"blurhash":    parsedBody.BlurHash,
		"user": map[string]interface{}{
			"name": parsedBody.User.Name,
			"url":  parsedBody.User.Links.HTML,
		},
	}, nil
}

var UnsplashImageObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "UnsplashImage",
	Fields: graphql.Fields{
		"blurhash": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type:              graphql.String,
			DeprecationReason: "Please use url_full or url_regular instead",
		},
		"url_full": &graphql.Field{
			Type: graphql.String,
		},
		"url_regular": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UnsplashUserObj,
		},
		"meta_url": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var UnsplashUserObj = graphql.NewObject(graphql.ObjectConfig{
	Name: "UnsplashUser",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
