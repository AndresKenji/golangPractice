//go:build windows

package platform

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func ApplyMetadata(srcPath, dstPath string, srcInfo os.FileInfo, preserveSecurity, preserveOwner bool) error {
	if err := os.Chmod(dstPath, srcInfo.Mode()); err != nil {
		return fmt.Errorf("chmod: %w", err)
	}

	mod := srcInfo.ModTime()
	if err := os.Chtimes(dstPath, time.Now(), mod); err != nil {
		return fmt.Errorf("chtimes: %w", err)
	}

	if preserveSecurity {
		if err := copyACLWithPowerShell(srcPath, dstPath); err != nil {
			return fmt.Errorf("copiar ACL: %w", err)
		}
	}

	return nil
}

func copyACLWithPowerShell(srcPath, dstPath string) error {
	src := strings.ReplaceAll(srcPath, "'", "''")
	dst := strings.ReplaceAll(dstPath, "'", "''")
	script := "Import-Module Microsoft.PowerShell.Security -ErrorAction SilentlyContinue; " +
		"$acl = Get-Acl -LiteralPath '" + src + "'; Set-Acl -LiteralPath '" + dst + "' -AclObject $acl"

	commands := [][]string{
		{"powershell", "-NoProfile", "-NonInteractive", "-Command", script},
		{"pwsh", "-NoProfile", "-NonInteractive", "-Command", script},
	}

	var lastErr error
	for _, c := range commands {
		cmd := exec.Command(c[0], c[1:]...)
		if out, err := cmd.CombinedOutput(); err != nil {
			lastErr = fmt.Errorf("%s: %v: %s", c[0], err, strings.TrimSpace(string(out)))
			continue
		}
		return nil
	}

	if lastErr != nil {
		return lastErr
	}
	return nil
}
