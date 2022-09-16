package container

import (
	"crypto/sha256"
	"os"
	"path/filepath"

	"github.com/opencontainers/go-digest"
)

type Blob struct {
	Sha256 digest.Digest
	Data   []byte
}

func (c *Container) addBlob(data []byte) (digest.Digest, int64, error) {
	h := sha256.Sum256(data)
	dig := digest.NewDigestFromBytes(digest.SHA256, h[:])

	c.Blobs = append(c.Blobs, Blob{
		Sha256: dig,
		Data:   data,
	})

	return dig, int64(len(data)), nil
}

func (c *Container) WriteBlobs(dirname string) error {
	dir := filepath.Join(dirname, "blobs", "sha256")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	for _, blob := range c.Blobs {
		err := os.WriteFile(filepath.Join(dir, blob.Sha256.Encoded()), blob.Data, 0644)
		if err != nil {
			return err
		}

	}
	return nil
}
