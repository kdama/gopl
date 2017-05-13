// Package github は、GitHub に対する Go の API を提供します。
package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetIssues は、指定された GitHub リポジトリから一定数の Issue を取得します。
func GetIssues(owner, repo string) (*[]Issue, error) {
	req, err := http.NewRequest("GET", getIssuesURL(owner, repo), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// この行よりも下の全てのパスで、resp.Body をクローズする必要があります。
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed: %s", resp.Status)
	}

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return &issues, nil
}
