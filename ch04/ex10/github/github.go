// Package github は、GitHub に対する Go の API を提供します。
package github

import "time"

// IssuesURL は、GitHub Issue を検索するための URL です。
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult は、GitHub Issue の検索結果を表します。
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue は、GitHub Issue を表します。
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User は、GitHub User を表します。
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
