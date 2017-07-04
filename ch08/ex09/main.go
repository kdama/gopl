// ch08/ex09 は、root ディレクトリのそれぞれに対して個別の合計を計算して定期的に表示する du です。
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"time"

	"github.com/kdama/gopl/ch08/ex01/client/console"
)

type Root struct {
	path  string
	files int64
	bytes int64
}

func main() {
	// Determine the initial directories.
	var roots []Root
	for _, rootPath := range os.Args[1:] {
		roots = append(roots, Root{rootPath, 0, 0})
	}
	if len(roots) == 0 {
		roots = []Root{Root{".", 0, 0}}
	}

	go func() {
		for {
			var n sync.WaitGroup
			for i := range roots {
				n.Add(1)
				i := i
				go func() {
					// Traverse each root of the file tree in parallel.
					fileSizes := make(chan int64)
					var m sync.WaitGroup
					m.Add(1)
					go func() {
						walkDir(roots[i].path, &m, fileSizes)
					}()
					go func() {
						m.Wait()
						close(fileSizes)
						n.Done()
					}()
					var nfiles, nbytes int64
					for {
						size, ok := <-fileSizes
						if !ok {
							break
						}
						nfiles++
						nbytes += size
					}
					roots[i].bytes = nbytes
					roots[i].files = nfiles
				}()
			}
			n.Wait()
			<-time.After(500 * time.Millisecond)
		}
	}()

	for {
		<-time.After(500 * time.Millisecond)
		printDiskUsage(roots) // final totals
	}
}

func printDiskUsage(roots []Root) {
	var rows [][]string

	for _, root := range roots {
		rows = append(rows, []string{
			root.path,
			fmt.Sprintf("%d files", root.files),
			fmt.Sprintf("%.1f GB", float64(root.bytes)/1e9),
		})
	}

	console.Clear()
	fmt.Println(console.SprintTable(rows))
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
