package container

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func (c *Container) WriteIndex(dirname string) error {
	b, err := json.Marshal(c.Index)
	if err != nil {
		return err
	}
	return os.WriteFile(
		filepath.Join(dirname, "index.json"),
		b,
		0644,
	)
}
