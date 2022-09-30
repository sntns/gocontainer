package container

import (
	"encoding/json"
	"path/filepath"

	"github.com/opencontainers/go-digest"
	ocischema "github.com/opencontainers/image-spec/specs-go"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/sntns/gocontainer/pkg/binary"
)

type Binary struct {
	Platform ocischemav1.Platform
	File     string
}

func (c *Container) createImageFromBinary(fname string, pf binary.Platform) (digest.Digest, int64, error) {

	rootfs, diffid, err := c.createLayerFrom("/", fname)
	if err != nil {
		return digest.Digest(""), 0, err
	}

	config, err := c.createConfig(
		filepath.Join(
			"/",
			filepath.Base(fname),
		),
		append(c.SharedDiffIDs, diffid),
		pf,
	)
	if err != nil {
		return digest.Digest(""), 0, err
	}

	m := ocischemav1.Manifest{
		MediaType: ocischemav1.MediaTypeImageManifest,
		Versioned: ocischema.Versioned{
			SchemaVersion: 2,
		},
		Config: config,
		Layers: append(
			c.Shared,
			rootfs,
		),
	}

	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return digest.Digest(""), 0, err
	}

	return c.addBlob(b)
}

func (c *Container) addBinary(fname string) (ocischemav1.Descriptor, error) {
	pf, err := binary.GetInfo(fname)
	if err != nil {
		return ocischemav1.Descriptor{}, err
	}

	dig, size, err := c.createImageFromBinary(fname, pf)
	if err != nil {
		return ocischemav1.Descriptor{}, err
	}

	return ocischemav1.Descriptor{
		MediaType: ocischemav1.MediaTypeImageManifest,
		Digest:    dig,
		Size:      size,
		Platform: &ocischemav1.Platform{
			Architecture: string(pf.Architecture),
			OS:           string(pf.OS),
		},
	}, nil
}

func (c *Container) AddBinary(fnames []string) error {
	index := ocischemav1.Index{
		MediaType: ocischemav1.MediaTypeImageIndex,
		Versioned: ocischema.Versioned{
			SchemaVersion: 2,
		},
		Manifests: []ocischemav1.Descriptor{},
	}

	for _, fname := range fnames {
		desc, err := c.addBinary(fname)
		if err != nil {
			return err
		}
		index.Manifests = append(index.Manifests, desc)
	}

	b, err := json.MarshalIndent(index, "", "    ")
	if err != nil {
		return err
	}

	dig, size, err := c.addBlob(b)
	if err != nil {
		return err
	}
	c.Index.Manifests = append(c.Index.Manifests, ocischemav1.Descriptor{
		MediaType: ocischemav1.MediaTypeImageIndex,
		Digest:    dig,
		Size:      size,
	})
	return nil
}
