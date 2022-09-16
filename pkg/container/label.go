package container

import "strings"

func (c *Container) SetLabels(labels ...string) error {
	for _, label := range labels {
		tokens := strings.Split(label, "=")
		if len(tokens) != 2 {
			return ErrorInvalidLabelFormat
		}
		key, value := tokens[0], tokens[1]
		c.Labels[key] = value
	}
	return nil
}
