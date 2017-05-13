// Package github は、GitHub に対する Go の API を提供します。
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// patchIssue は、GitHub Issue に関する書き込み操作を行います。
func patchIssue(owner, repo, number string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("PATCH", getIssueURL(owner, repo, number), buf)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github.v3.text-match+json")
	req.Header.Set("Content-Type", "application/json")
	err = setAuthorization(req)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// この行よりも下の全てのパスで、resp.Body をクローズする必要があります。
	resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("patch issue failed: %s", resp.Status)
	}

	return nil
}

// UpdateIssue は、GitHub Issue を更新します。
func UpdateIssue(owner, repo, number string, fields map[string]string) error {
	return patchIssue(owner, repo, number, fields)
}

// ReopenIssue は、Close された GitHub Issue を再び Open します。
func ReopenIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "open",
	}
	return patchIssue(owner, repo, number, fields)
}

// CloseIssue は、GitHub Issue を編集します。
func CloseIssue(owner, repo, number string) error {
	fields := map[string]string{
		"state": "closed",
	}
	return patchIssue(owner, repo, number, fields)
}
