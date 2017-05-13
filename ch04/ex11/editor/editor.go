// Package editor は、環境変数 EDITOR で指定されたエディタによってデータを操作します。
package editor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

// Edit は、環境変数 EDITOR で指定されたエディタを起動し、ユーザーに value を編集させます。
func Edit(value map[string]string) error {
	editor := os.Getenv("EDITOR")

	// vi は、ほとんど全ての Linux ディストリビューションに含まれます。
	if editor == "" {
		editor = "vi"
	}

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}

	tempFileName := tempFile.Name()
	defer os.Remove(tempFileName)

	encoder := json.NewEncoder(tempFile)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(value)
	if err != nil {
		return err
	}

	err = tempFile.Close()
	if err != nil {
		return err
	}

	cmd := exec.Command(editor, tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return err
	}

	edited, err := ioutil.ReadFile(tempFileName)
	if err != nil {
		return err
	}

	// Windows Notepad などいくつかのエディタは、UTF-8 の保存時に常に BOM を追加します。
	// json.Unmarshal は、UTF-8 BOM が付いたデータに対応していないので、予め削除します。
	err = json.Unmarshal(removeUTF8BOM(edited), &value)
	if err != nil {
		return err
	}
	return nil
}

// removeUTF8BOM は、バイト列の先頭に UTF-8 BOM があった場合、それを削除します。
func removeUTF8BOM(s []byte) []byte {
	utf8Bom := []byte{239, 187, 191}
	return bytes.TrimPrefix(s, utf8Bom)
}
