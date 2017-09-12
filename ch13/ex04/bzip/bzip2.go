// Package bzip は、bzip2 圧縮を使うライターを提供します。
package bzip

import (
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	mu  sync.Mutex
	wc  io.WriteCloser
	cmd *exec.Cmd
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("/bin/bzip2")
	cmd.Stdout = out
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	w := &writer{wc: stdin, cmd: cmd}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (w *writer) Write(data []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.wc.Write(data)
}

func (w *writer) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	err := w.wc.Close()
	if waitErr := w.cmd.Wait(); err == nil {
		err = waitErr
	}
	if err != nil {
		return err
	}
	return nil
}
