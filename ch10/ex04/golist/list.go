package golist

import (
	"bytes"
	"encoding/json"
	"io"
	"os/exec"
)

// List は、go list ツールを利用して、Go パッケージのメタデータを返します。
func List(template ...string) ([]Package, error) {
	cmd := exec.Command("go", append([]string{"list", "-json"}, template...)...)
	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var packages []Package
	dec := json.NewDecoder(bytes.NewReader(b))
	for {
		var val Package
		err := dec.Decode(&val)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		packages = append(packages, val)
	}
	return packages, nil
}
