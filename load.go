package iso3166ja

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const src = "https://ja.wikipedia.org/w/api.php?format=xml&action=query&titles=ISO_3166-1&prop=revisions&rvprop=content"

// LoadWikipediaJA は Wikipedia 日本語版の ISO_3166-1 のページをダウンロードし、 wiki text を取り出します。
func LoadWikipediaJA() (string, error) {
	return Load(src)
}

// Load は Wikipedia の ISO_3166-1 のページの XML データを渡すと、ダウンロードして wiki text を取り出します。
func Load(sourceURL string) (string, error) {
	resp, err := http.Get(sourceURL)
	if err != nil {
		return "", fmt.Errorf("failed to get source: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			println("failed to close response body:", err.Error())
		}
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	tree := struct {
		Query struct {
			Pages []struct {
				Page struct {
					Revisions []struct {
						Revision string `xml:"rev"`
					} `xml:"revisions"`
				} `xml:"page"`
			} `xml:"pages"`
		} `xml:"query"`
	}{}
	if err := xml.Unmarshal(body, &tree); err != nil {
		return "", fmt.Errorf("failed to unmarshal content: %w", err)
	}
	if len(tree.Query.Pages) < 1 || len(tree.Query.Pages[0].Page.Revisions) < 1 {
		return "", errors.New("no such content")
	}
	return tree.Query.Pages[0].Page.Revisions[0].Revision, nil
}
