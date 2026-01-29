// +build darwin

package overseer

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

var (
	supported = true
	uid       = syscall.Getuid()
	gid       = syscall.Getgid()
	SIGUSR1   = syscall.SIGUSR1
	SIGUSR2   = syscall.SIGUSR2
	SIGTERM   = syscall.SIGTERM
)

func move(dst, src string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}
	
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}
	
	if err := os.Rename(src, dst); err == nil {
		return nil
	}
	
	cmd := exec.Command("mv", "-f", src, dst)
	if err := cmd.Run(); err != nil {
		return err
	}
	
	return syncCmd().Run()
}

func syncCmd() *exec.Cmd {
	return exec.Command("sync")
}

func chmod(f *os.File, perms os.FileMode) error {
	return f.Chmod(perms)
}

func chown(f *os.File, uid, gid int) error {
	return f.Chown(uid, gid)
}
