package container

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHealthcheck(t *testing.T) {
	r := require.New(t)
	c, err := parseHealthcheckArg(`--retries 2 --interval 10s CMD [ "testcommand" "--doit" ]`)
	r.NoError(err)
	r.Equal(2, c.Retries)
	r.Equal(10*time.Second, *c.Interval)
	r.Nil(c.StartPeriod)
	r.Equal([]string{"CMD", "testcommand", "--doit"}, c.Test)
}
