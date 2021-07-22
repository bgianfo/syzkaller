// Copyright 2021 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package main

import (
	"path/filepath"
	"strings"
        "os"
        "fmt"
	"github.com/google/syzkaller/pkg/compiler"
)

type serenityos struct{}

func (*serenityos) prepare(sourcedir string, build bool, arches []*Arch) error {
	return nil
}

func (*serenityos) prepareArch(arch *Arch) error {
	archName := arch.target.Arch
	// Use the the correct name for SerenityOS/i386
	if archName == "386" {
		archName = "i386"
	}
	if err := os.Symlink(filepath.Join(arch.sourceDir, "sys", archName, "include"),
		filepath.Join(arch.buildDir, "machine")); err != nil {
		return fmt.Errorf("failed to create link: %v", err)
	}
	if err := os.Symlink(filepath.Join(arch.sourceDir, "sys", "x86", "include"),
		filepath.Join(arch.buildDir, "x86")); err != nil {
		return fmt.Errorf("failed to create link: %v", err)
	}
	return nil
}

func (*serenityos) processFile(arch *Arch, info *compiler.ConstInfo) (map[string]uint64, map[string]bool, error) {
	dir := arch.sourceDir
	args := []string{
		"-fmessage-length=0",
		"-D__KERNEL__",
		"-DROS_KERNEL",
		"-I", filepath.Join(dir, "kern", "include"),
		"-I", filepath.Join(dir, "user", "parlib", "include"),
	}
	for _, incdir := range info.Incdirs {
		args = append(args, "-I"+filepath.Join(dir, incdir))
	}
	if arch.includeDirs != "" {
		for _, dir := range strings.Split(arch.includeDirs, ",") {
			args = append(args, "-I"+dir)
		}
	}
	params := &extractParams{
		DeclarePrintf: true,
		TargetEndian:  arch.target.HostEndian,
	}
	return extract(info, "gcc", args, params)
}
