//go:build !windows

package platform

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func ApplyMetadata(srcPath, dstPath string, srcInfo os.FileInfo, preserveSecurity, preserveOwner bool) error {
	if err := os.Chmod(dstPath, srcInfo.Mode()); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	if preserveOwner {
		if stat, ok := srcInfo.Sys().(*syscall.Stat_t); ok {
			if err := os.Chown(dstPath, int(stat.Uid), int(stat.Gid)); err != nil {
				return fmt.Errorf("chown: %w", err)
			}
		}
	}

	mod := srcInfo.ModTime()
	if err := os.Chtimes(dstPath, time.Now(), mod); err != nil {
		return fmt.Errorf("chtimes: %w", err)
	}

	return nil
}
