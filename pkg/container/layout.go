package container

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (c *Container) WriteLayout(dirname string) error {
	b, err := json.Marshal(c.Layout)
	if err != nil {
		return err
	}
	return os.WriteFile(
		filepath.Join(dirname, "oci-layout"),
		b,
		0644,
	)
}
