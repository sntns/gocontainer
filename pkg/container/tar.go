package container

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"

	"github.com/opencontainers/go-digest"
)

type TarItem struct {
	File        string
	Destination string
}

func addToArchive(tw *tar.Writer, filename string, destname string) error {
	// Open the file which will be written into the archive
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get FileInfo about our file providing file size, mode, etc.
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Create a tar Header from the FileInfo data
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	// Use full path as name (FileInfoHeader only takes the basename)
	// If we don't do this the directory strucuture would
	// not be preserved
	// https://golang.org/src/archive/tar/common.go?#L626
	header.Name = destname

	// Write file header to the tar archive
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	// Copy file content to tar archive
	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}

func createTarFrom(items ...TarItem) ([]byte, digest.Digest, error) {
	buf := bytes.NewBuffer([]byte{})

	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Iterate over files and add them to the tar archive
	for _, item := range items {
		err := addToArchive(tw, item.File, item.Destination)
		if err != nil {
			return nil, digest.Digest(""), err
		}
	}

	err := tw.Flush()
	if err != nil {
		return nil, digest.Digest(""), err
	}
	tw.Close()

	b := buf.Bytes()
	diffid := digest.FromBytes(b)

	comp := bytes.NewBuffer([]byte{})
	gw := gzip.NewWriter(comp)
	if _, err := io.Copy(gw, buf); err != nil {
		return nil, digest.Digest(""), err
	}
	gw.Close()

	return comp.Bytes(), diffid, nil
}
