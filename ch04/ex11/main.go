package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kdama/gopl/ch04/ex11/editor"
	"github.com/kdama/gopl/ch04/ex11/github"
)

const (
	usage = `usage:
    create OWNER REPO
    get    OWNER REPO ISSUE_NUMBER
    edit   OWNER REPO ISSUE_NUMBER
    close  OWNER REPO ISSUE_NUMBER
    reopen OWNER REPO ISSUE_NUMBER`
)

func create(owner, repo string) {
	fields := map[string]string{
		"title": "",
		"body":  "",
	}
	err := editor.Edit(fields)
	if err != nil {
		log.Fatal(err)
	}

	err = github.CreateIssue(owner, repo, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func get(owner, repo, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issue:    %s/%s (#%d)\n", owner, repo, issue.Number)
	fmt.Printf("url:      %v\n", issue.HTMLURL)
	fmt.Printf("title:    %v\n", issue.Title)
	fmt.Printf("user:     %s\n", issue.User.Login)
	fmt.Printf("state:    %v\n", issue.State)
	fmt.Printf("comments: %v\n", issue.Comments)
	fmt.Printf("created:  %v\n", issue.CreatedAt)
	fmt.Printf("updated:  %v\n", issue.UpdatedAt)
	fmt.Printf("\n%s\n", issue.Body)
}

func edit(owner string, repo string, number string) {
	issue, err := github.GetIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]string{
		"title": issue.Title,
		"body":  issue.Body,
	}
	err = editor.Edit(fields)
	if err != nil {
		log.Fatal(err)
	}

	err = github.UpdateIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func close(owner string, repo string, number string) {
	err := github.CloseIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
}

func reopen(owner string, repo string, number string) {
	err := github.ReopenIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 4 {
		command, owner, repo := os.Args[1], os.Args[2], os.Args[3]
		switch command {
		case "create":
			create(owner, repo)
		default:
			log.Fatal(usage)
		}
	} else if len(os.Args) == 5 {
		command, owner, repo, number := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
		switch command {
		case "get":
			get(owner, repo, number)
		case "edit":
			edit(owner, repo, number)
		case "close":
			close(owner, repo, number)
		case "reopen":
			reopen(owner, repo, number)
		default:
			log.Fatal(usage)
		}
	} else {
		log.Fatal(usage)
	}
}
