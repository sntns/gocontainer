package container

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTarFromIgnoresTheBuildMachine(t *testing.T) {
	r := require.New(t)
	dir := t.TempDir()
	bin := filepath.Join(dir, "svc")
	conf := filepath.Join(dir, "config.yaml")
	r.NoError(os.WriteFile(bin, []byte("binary"), 0o755))
	r.NoError(os.WriteFile(conf, []byte("key: value"), 0o644))
	items := []TarItem{{File: bin, Destination: "/svc"}, {File: conf, Destination: "/config.yaml"}}

	_, first, err := createTarFrom(items...)
	r.NoError(err)

	later := time.Now().Add(time.Hour)
	r.NoError(os.Chtimes(bin, later, later))
	r.NoError(os.Chmod(bin, 0o775))
	r.NoError(os.Chmod(conf, 0o664))
	_, second, err := createTarFrom(items...)
	r.NoError(err)
	r.Equal(first, second, "times and group bits of the files")

	r.NoError(os.Chmod(bin, 0o644))
	_, notExecutable, err := createTarFrom(items...)
	r.NoError(err)
	r.NotEqual(first, notExecutable, "the executable bit")
}

func TestEpoch(t *testing.T) {
	r := require.New(t)
	t.Setenv("SOURCE_DATE_EPOCH", "")
	r.Equal(time.Unix(0, 0).UTC(), Epoch())
	t.Setenv("SOURCE_DATE_EPOCH", "1700000000")
	r.Equal(time.Unix(1700000000, 0).UTC(), Epoch())
}
