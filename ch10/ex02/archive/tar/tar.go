package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/kdama/gopl/ch10/ex02/archive"
)

func list(f *os.File) ([]archive.FileHeader, error) {
	var headers []archive.FileHeader

	// Open the tar archive for reading.
	tr := tar.NewReader(f)

	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return nil, err
		}
		headers = append(headers, archive.FileHeader{
			Name: hdr.Name,
			Size: uint64(hdr.Size),
		})
	}
	return headers, nil
}

func init() {
	archive.RegisterFormat("tar", "ustar\x0000", 257, list)

	// GNU tar はサポートされません。
	// archive.RegisterFormat("tar", "ustar\x0040\x0050", 257, list)
}
