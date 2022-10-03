package container

import (
	"encoding/json"
	"time"

	"github.com/opencontainers/go-digest"
	ocischemav1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sntns/gocontainer/pkg/binary"
)

type (
	// Image is copied from ocischemav1 (except the Config)
	Image struct {
		Created      *time.Time            `json:"created,omitempty"`
		Author       string                `json:"author,omitempty"`
		Architecture string                `json:"architecture"`
		OS           string                `json:"os"`
		Config       ImageConfig           `json:"config,omitempty"`
		RootFS       ocischemav1.RootFS    `json:"rootfs"`
		History      []ocischemav1.History `json:"history,omitempty"`
	}

	// ImageConfig is copied from ocischemav1.ImageConfig, only selecting
	// used item and adding:
	//   - healthcheck
	ImageConfig struct {
		// ImageConfig defines the execution parameters which should be used as a base when running a container using an image.
		Entrypoint []string `json:"Entrypoint,omitempty"`
		WorkingDir string   `json:"WorkingDir,omitempty"`

		Healthcheck *HealthcheckConfig  `json:"Healthcheck,omitempty"`
		Volumes     map[string]struct{} `json:"Volumes,omitempty"`
		Labels      map[string]string   `json:"Labels,omitempty"`
	}
)

func (c *Container) createConfig(entrypoint string, diffIDs []digest.Digest, pf binary.Platform) (ocischemav1.Descriptor, error) {
	now := time.Now()
	config := Image{
		Architecture: string(pf.Architecture),
		OS:           string(pf.OS),
		Created:      &now,
		Config: ImageConfig{
			Entrypoint: []string{
				entrypoint,
			},
			Healthcheck: c.Healthcheck,
			WorkingDir:  "/",
			Labels:      c.Labels,
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
