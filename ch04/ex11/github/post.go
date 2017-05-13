// Package github は、GitHub に対する Go の API を提供します。
package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateIssue は、指定された GitHub 上のリポジトリに Issue を作成します。
func CreateIssue(owner, repo string, fields map[string]string) error {
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(fields)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", getIssuesURL(owner, repo), buf)
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("post issue failed: %s", resp.Status)
	}
	return nil
}
