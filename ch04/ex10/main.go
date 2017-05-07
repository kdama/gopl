// ch04/ex10 は、検索語に一致した GitHub Issue の表を、Issue が作成された期間で分類して表示します。
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kdama/gopl/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()

	// 1 か月 (30 日) 未満に作成された Issue を報告します。
	fmt.Println("\n-- created at less than a month --")
	for _, item := range result.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days < 30 {
			printIssue(item)
		}
	}

	// 1 か月 (30 日) 以上 1 年 (365 日) 未満に作成された Issue を報告します。
	fmt.Println("\n-- created at less than a year --")
	for _, item := range result.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days >= 30 && days < 365 {
			printIssue(item)
		}
	}

	// 1 年 (365 日) 以上に作成された Issue を報告します。
	fmt.Println("\n-- created at more than a year --")
	for _, item := range result.Items {
		days := now.Sub(item.CreatedAt).Hours() / 24
		if days >= 365 {
			printIssue(item)
		}
	}
}

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
}
