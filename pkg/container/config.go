package container

import (
	"encoding/json"
	"time"

	"github.com/opencontainers/go-digest"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sntns/gocontainer/pkg/binary"
)

func (c *Container) createConfig(entrypoint string, diffIDs []digest.Digest, pf binary.Platform) (ocischemav1.Descriptor, error) {
	now := time.Now()
	config := ocischemav1.Image{
		Architecture: string(pf.Architecture),
		OS:           string(pf.OS),
		Created:      &now,
		Config: ocischemav1.ImageConfig{
			Entrypoint: []string{
				entrypoint,
			},
			WorkingDir: "/",
			Labels:     c.Labels,
		},
		RootFS: ocischemav1.RootFS{
			Type:    "layers",
			DiffIDs: diffIDs,
		},
	}

	b, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		return ocischemav1.Descriptor{}, nil
	}

	dig, size, err := c.addBlob(b)
	if err != nil {
		return ocischemav1.Descriptor{}, err
	}

	return ocischemav1.Descriptor{
		MediaType: ocischemav1.MediaTypeImageConfig,
		Digest:    dig,
		Size:      size,
	}, nil
}
