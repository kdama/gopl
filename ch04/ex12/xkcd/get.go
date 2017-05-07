// Package xkcd は、xkcd に対する Go の API を提供します。
package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetComic は xkcd からコミックを取得します。
func GetComic(comicID int) (*Comic, error) {
	resp, err := http.Get(getComicURL(comicID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get comic failed: %s", resp.Status)
	}

	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
