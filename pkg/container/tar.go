package container

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"

	digest "github.com/opencontainers/go-digest"
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

	// Only the content and the executable bit: no owner, time nor umask of the machine that builds.
	mode := int64(0o644)
	if info.Mode()&0o111 != 0 {
		mode = 0o755
	}
	header := &tar.Header{Typeflag: tar.TypeReg, Name: destname, Mode: mode, Size: info.Size(), ModTime: Epoch()}

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
