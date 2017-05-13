// ch04/ex14 は、GitHub への一度の問い合わせで、バグレポート、マイルストーン、ユーザーの一覧を閲覧可能にするウェブサーバです。
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/kdama/gopl/ch04/ex14/github"
)

var navigation = `
<p>
<a href='/'>Issues</a> /
<a href='/milestones'>Milestones</a> /
<a href='/users'>Users</a>
</p>
`
var issuesTemplate = template.Must(template.New("issues").Parse(navigation + `
<h1>{{len .}} issue{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var milestonesTemplate = template.Must(template.New("milestones").Parse(navigation + `
<h1>{{len .}} milestone{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th>Title</th>
<th>State</th>
</tr>
{{range .}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td>{{.State}}</td>
</tr>
{{end}}
</table>
`))

var usersTemplate = template.Must(template.New("users").Parse(navigation + `
<h1>{{len .}} user{{if ne (len .) 1}}s{{end}}</h1>
<table>
<tr style='text-align: left'>
<th>Avatar</th>
<th>Username</th>
</tr>
{{range .}}
<tr>
	<td><a href='{{.HTMLURL}}'><img src='{{.AvatarURL}}' width='32' height='32'></td>
	<td><a href='{{.HTMLURL}}'>{{.Login}}</td>
</tr>
{{end}}
</table>
`))

var issues []github.Issue
var milestones []github.Milestone
var users []github.User

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: ex14 OWNER REPO")
		os.Exit(1)
	}

	owner, repo := os.Args[1], os.Args[2]
	err := generateCache(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handleIssues)
	http.HandleFunc("/milestones", handleMilestones)
	http.HandleFunc("/users", handleUsers)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func generateCache(owner, repo string) error {
	got, err := github.GetIssues(owner, repo)
	if err != nil {
		return err
	}

	issues = *got

	for _, issue := range issues {
		if issue.Milestone != nil {
			milestones = appendMilestoneAsSet(milestones, issue.Milestone)
		}
		for _, assignee := range issue.Assignees {
			users = appendUserAsSet(users, assignee)
		}
		users = appendUserAsSet(users, issue.User)
	}
	return nil
}

func handleIssues(w http.ResponseWriter, r *http.Request) {
	issuesTemplate.Execute(w, issues)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	usersTemplate.Execute(w, users)
}

func handleMilestones(w http.ResponseWriter, r *http.Request) {
	milestonesTemplate.Execute(w, milestones)
}

// appendMilestoneAsSet は、与えられたマイルストーンの配列 set に、与えられたマイルストーン milestone を追加します。
// ただし、set が milestone を既に含んでいる場合は、milestone を追加しません。
func appendMilestoneAsSet(set []github.Milestone, milestone *github.Milestone) []github.Milestone {
	if !includesMilestone(set, milestone) {
		return append(set, *milestone)
	}
	return set
}

// includesMilestone は、与えられたマイルストーンの配列 array が、与えられたマイルストーン milestone を含んでいるかどうかを返します。
func includesMilestone(array []github.Milestone, milestone *github.Milestone) bool {
	for _, value := range array {
		if value.Equals(milestone) {
			return true
		}
	}
	return false
}

// appendUserAsSet は、与えられたユーザーの配列 set に、与えられたユーザー user を追加します。
// ただし、set が user を既に含んでいる場合は、user を追加しません。
func appendUserAsSet(set []github.User, user *github.User) []github.User {
	if !includesUser(set, user) {
		return append(set, *user)
	}
	return set
}

// includesUser は、与えられたユーザーの配列 array が、与えられたユーザー user を含んでいるかどうかを返します。
func includesUser(array []github.User, user *github.User) bool {
	for _, value := range array {
		if value.Equals(user) {
			return true
		}
	}
	return false
}
