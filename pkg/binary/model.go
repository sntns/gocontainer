package binary

type OS string
type Architecture string

const (
	OSLinux  OS = "linux"
	OSDarwin OS = "darwin"
)

const (
	ArchAMD64 Architecture = "amd64"
	ArchARM64 Architecture = "arm64"
)

type Platform struct {
	Architecture Architecture
	OS           OS
}

func (pf Platform) String() string {
	return string(pf.OS) + "/" + string(pf.Architecture)
}
