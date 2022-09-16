package binary

type OS string
type Architecture string

const (
	OS_LINUX  OS = "linux"
	OS_DARWIN OS = "darwin"
)

const (
	ARCH_AMD64 Architecture = "amd64"
	ARCH_ARM64 Architecture = "arm64"
)

type Platform struct {
	Architecture Architecture
	OS           OS
}

func (pf Platform) String() string {
	return string(pf.OS) + "/" + string(pf.Architecture)
}
