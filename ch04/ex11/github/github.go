// Package github は、GitHub に対する Go の API を提供します。
package github

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getIssueURL(owner string, repo string, number string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues/%s", owner, repo, number)
}

func getIssuesURL(owner string, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
}

func setAuthorization(req *http.Request) error {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return fmt.Errorf("GITHUB_TOKEN is not set")
	}
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	return nil
}

// Issue は、GitHub Issue を表します。
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Comments  int
	UpdatedAt time.Time `json:"updated_at"`
}

// User は、GitHub 上のユーザーを表します。
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
