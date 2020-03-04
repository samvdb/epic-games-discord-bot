package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Request struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables"`
}

type Variables struct {
	Namespace string `json:"namespace"`
	Country   string `json:"country"`
	Locale    string `json:"locale"`
}

const URL = "https://graphql.epicgames.com/graphql"

type Response struct {
	Data struct {
		Catalog struct {
			CatalogOffers struct {
				Elements []Game `json:"elements"`
			} `json:"catalogOffers"`
		} `json:"Catalog"`
	} `json:"data"`
}

type Game struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      []struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"keyImages"`
}

type Api struct {
}

func (a *Api) get() (*Response, error) {

	query := a.createQuery()
	request := Request{
		Query: query,
		Variables: Variables{
			Country:   "BE",
			Locale:    "en-US",
			Namespace: "epic",
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New("HTTP client error " + string(res.StatusCode) + ": " + res.Status)
	}

	response := &Response{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (a *Api) createQuery() string {
	category := "freegames"
	sort := "effectiveDate"
	sortDir := "asc"

	q := "query promotionsQuery($namespace: String!, $country: String!, $locale: String!) { Catalog { "
	q += fmt.Sprintf(`catalogOffers(namespace: $namespace, locale: $locale, params: {category: "%s", country: $country, sortBy: "%s", sortDir: "%s"})`, category, sort, sortDir)
	// response element
	q += "{ elements { title description id keyImages { type url } productSlug promotions { promotionalOffers { promotionalOffers { startDate endDate } } } } }"
	q += "} }"

	return q
}
