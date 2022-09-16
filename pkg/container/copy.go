package container

import (
	"os"
	"strings"
)

func (c *Container) Copy(names ...string) error {
	for _, name := range names {
		tokens := strings.Split(name, ":")
		var fname, destination string
		switch len(tokens) {
		case 1:
			fname = tokens[0]
			destination = tokens[0]
		case 2:
			fname = tokens[0]
			destination = tokens[1]
		default:
			return ErrInvalidCopyName
		}

		info, err := os.Stat(fname)
		if err != nil {
			return err
		}

		if info.IsDir() {
			d, diffid, err := c.createLayerFromDir(destination, fname)
			if err != nil {
				return err
			}
			c.Shared = append(c.Shared, d)
			c.SharedDiffIDs = append(c.SharedDiffIDs, diffid)
			continue
		}

		d, diffid, err := c.createLayerFromFiles(destination, fname)
		if err != nil {
			return err
		}
		c.Shared = append(c.Shared, d)
		c.SharedDiffIDs = append(c.SharedDiffIDs, diffid)

	}
	return nil
}
