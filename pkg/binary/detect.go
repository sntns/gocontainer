package binary

import (
	"debug/elf"
	"debug/macho"
	"errors"
)

type infoFunction func(name string) (Platform, error)

var (
	ErrOsNotSupported           = errors.New("OS not supported")
	ErrArchitectureNotSupported = errors.New("Architecture not supported")
)

func GetInfo(name string) (Platform, error) {
	var err error
	var pf Platform
	for _, fn := range []infoFunction{
		machoInfo,
		elfInfo,
	} {
		pf, err = fn(name)
		if err != nil {
			continue
		}
		return pf, nil
	}
	return pf, err
}

func elfInfo(name string) (Platform, error) {
	f, err := elf.Open(name)
	if err != nil {
		return Platform{}, err
	}
	defer f.Close()

	mmachine := map[elf.Machine]Architecture{
		elf.EM_AARCH64: ARCH_ARM64,
		elf.EM_X86_64:  ARCH_AMD64,
	}
	arch, ok := mmachine[f.Machine]
	if !ok {
		return Platform{}, ErrArchitectureNotSupported
	}

	return Platform{
		OS:           OS_LINUX,
		Architecture: arch,
	}, nil
}

func machoInfo(name string) (Platform, error) {
	f, err := macho.Open(name)
	if err != nil {
		return Platform{}, err
	}
	mcpu := map[macho.Cpu]Architecture{
		macho.CpuArm64: ARCH_ARM64,
		macho.CpuAmd64: ARCH_AMD64,
	}
	arch, ok := mcpu[f.Cpu]
	if !ok {
		return Platform{}, ErrArchitectureNotSupported
	}
	return Platform{
		OS:           "darwin",
		Architecture: arch,
	}, nil
}
