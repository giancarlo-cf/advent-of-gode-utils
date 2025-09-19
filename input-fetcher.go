package advent_of_gode_utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type HttpFetcher interface {
	FetchInput(year int, day int, sessionCookie string) (string, error)
}

type InputFetcher struct{}

const url string = "https://adventofcode.com/%d/day/%d/input"

var client = http.Client{}

func (f *InputFetcher) FetchInput(year int, day int, sessionCookie string) (string, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf(url, year, day), nil)
	if err != nil {
		return "", err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error fetching the Input: %v\n", err))
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error reading response body: %v\n", err))
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error parsing the fetch response: %v\n", err))
	}
	title := doc.Find("title").First().Text()
	if title == "500 Internal Server Error" {
		return "", errors.New(fmt.Sprintf("Could not find the Input data. Check your session cookie.\nResponse body:\n%s", doc.Text()))
	}

	return string(bodyBytes), nil
}
