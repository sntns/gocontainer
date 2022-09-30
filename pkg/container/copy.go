package container

import (
	"strings"
)

func (c *Container) Copy(names ...string) error {
	items := []TarItem{}
	for _, name := range names {
		tokens := strings.Split(name, ":")
		var fname, target string
		switch len(tokens) {
		case 1:
			fname = tokens[0]
			target = tokens[0]
		case 2:
			fname = tokens[0]
			target = tokens[1]
		default:
			return ErrInvalidCopyName
		}

		its, err := createTarItemFrom(target, fname)
		if err != nil {
			return err
		}
		items = append(items, its...)
	}

	d, diffid, err := c.createLayer(items...)
	if err != nil {
		return err
	}
	c.Shared = append(c.Shared, d)
	c.SharedDiffIDs = append(c.SharedDiffIDs, diffid)

	return nil
}
