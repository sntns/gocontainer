package container

import (
	"strings"
)

func (c *Container) Copy(names ...string) error {
	items := []TarItem{}
	for _, name := range names {
		fname, target, found := strings.Cut(name, ":")
		if !found {
			target = fname
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
