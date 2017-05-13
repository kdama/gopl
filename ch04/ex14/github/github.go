// Package github は、GitHub に対する Go の API を提供します。
package github

import (
	"fmt"
	"time"
)

func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=all", owner, repo)
}

// Issue は、GitHub Issue を表します。
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Assignees []*User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	UpdatedAt time.Time `json:"updated_at"`
	Milestone *Milestone
}

// User は、GitHub 上のユーザーを表します。
type User struct {
	AvatarURL string `json:"avatar_url"`
	HTMLURL   string `json:"html_url"`
	ID        int
	Login     string
}

// Milestone は、GitHub 上のマイルストーンを表します。
type Milestone struct {
	Description string
	HTMLURL     string `json:"html_url"`
	ID          int
	State       string
	Title       string
}

// Equals は、与えられた User と等しいとみなしてよいかどうかを返します。
func (u *User) Equals(x *User) bool {
	return u.ID == x.ID
}

// Equals は、与えられた Milestone と等しいとみなしてよいかどうかを返します。
func (m *Milestone) Equals(x *Milestone) bool {
	return m.ID == x.ID
}
