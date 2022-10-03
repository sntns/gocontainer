package container

import (
	"os"

	"github.com/opencontainers/go-digest"
	ocischema "github.com/opencontainers/image-spec/specs-go"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type Container struct {
	Blobs         []Blob
	Shared        []ocischemav1.Descriptor
	SharedDiffIDs []digest.Digest
	Layout        ocischemav1.ImageLayout
	Index         ocischemav1.Index
	Labels        map[string]string
	Healthcheck   *HealthcheckConfig
}

type Layer struct {
}

func New() (*Container, error) {
	cont := &Container{
		Index: ocischemav1.Index{
			MediaType: ocischemav1.MediaTypeImageIndex,
			Versioned: ocischema.Versioned{
				SchemaVersion: 2,
			},
			Manifests: []ocischemav1.Descriptor{},
		},
		Layout: ocischemav1.ImageLayout{
			Version: ocischemav1.ImageLayoutVersion,
		},
		Labels: map[string]string{},
	}
	return cont, nil
}

func (c *Container) Save(dirname string) error {
	_, err := os.Stat(dirname)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	} else {
		err := os.RemoveAll(dirname)
		if err != nil {
			return err
		}
	}
	err = os.MkdirAll(dirname, 0755)
	if err != nil {
		return err
	}

	if err := c.WriteIndex(dirname); err != nil {
		return err
	}

	if err := c.WriteLayout(dirname); err != nil {
		return err
	}

	if err := c.WriteBlobs(dirname); err != nil {
		return err
	}

	return nil
}
