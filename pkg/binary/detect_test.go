package binary

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	cmd := exec.Command("make", "all")
	cmd.Dir = filepath.Join("testdata")
	err := cmd.Run()
	if err != nil {
		t.Error(err)
	}

	for k, tc := range []struct {
		desc     string
		platform Platform
	}{
		{
			desc:     "should detect linux/amd64",
			platform: Platform{OS: OS_LINUX, Architecture: ARCH_AMD64},
		},
		{
			desc:     "should detect linux/arm64",
			platform: Platform{OS: OS_LINUX, Architecture: ARCH_ARM64},
		},
		{
			desc:     "should detect darwin/arm64",
			platform: Platform{OS: OS_DARWIN, Architecture: ARCH_ARM64},
		},
	} {
		t.Run(fmt.Sprintf("k=%d/case=%s", k, tc.desc), func(t *testing.T) {
			r := require.New(t)
			fname := filepath.Join("testdata", fmt.Sprintf("helloworld-%s-%s", tc.platform.OS, tc.platform.Architecture))
			pf, err := GetInfo(fname)
			r.NoError(err)
			r.Equal(tc.platform.OS, pf.OS)
			r.Equal(tc.platform.Architecture, pf.Architecture)

		})
	}

}
