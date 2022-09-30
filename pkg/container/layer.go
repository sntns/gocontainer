package container

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/opencontainers/go-digest"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func (c *Container) createLayer(items ...TarItem) (ocischemav1.Descriptor, digest.Digest, error) {
	b, diffid, err := createTarFrom(items...)
	if err != nil {
		return ocischemav1.Descriptor{}, digest.Digest(""), err
	}

	dig, size, err := c.addBlob(b)
	if err != nil {
		return ocischemav1.Descriptor{}, digest.Digest(""), err
	}

	return ocischemav1.Descriptor{
		MediaType: ocischemav1.MediaTypeImageLayerGzip,
		Digest:    dig,
		Size:      size,
	}, diffid, nil
}

func createTarItemFrom(target string, fnames ...string) ([]TarItem, error) {
	items := []TarItem{}
	for _, fname := range fnames {
		info, err := os.Stat(fname)
		if err != nil {
			return nil, err
		}

		if info.IsDir() {
			err := filepath.Walk(fname, func(path string, info fs.FileInfo, err error) error {
				if len(info.Name()) > 2 && info.Name()[0] == '.' {
					// skip hidden files
					return filepath.SkipDir
				}
				if !info.IsDir() {
					items = append(items, TarItem{
						File:        path,
						Destination: strings.Replace(path, fname, target, 1),
					})
				}
				return nil
			})
			if err != nil {
				return nil, err
			}
		} else {
			items = append(items, TarItem{
				File:        fname,
				Destination: filepath.Join(target, filepath.Base(fname)),
			})
		}
	}

	return items, nil
}

func (c *Container) createLayerFrom(target string, fnames ...string) (ocischemav1.Descriptor, digest.Digest, error) {
	items, err := createTarItemFrom(target, fnames...)
	if err != nil {
		return ocischemav1.Descriptor{}, digest.Digest(""), err
	}
	return c.createLayer(items...)
}
