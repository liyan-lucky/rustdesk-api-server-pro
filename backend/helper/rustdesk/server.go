package rustdesk

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

var (
	serverBinDir = GetRustdeskServerBinDir()

	hbbrPidFile = filepath.Join(serverBinDir, "hbbr.pid")
	hbbsPidFile = filepath.Join(serverBinDir, "hbbs.pid")
)

func GetRustdeskServerBinDir() string {
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, "rustdesk-server")
}

func GetRustdeskServerBin() (hbbr, hbbs string) {
	dir := GetRustdeskServerBinDir()
	switch runtime.GOOS {
	case "windows":
		hbbr = filepath.Join(dir, "hbbr.exe")
		hbbs = filepath.Join(dir, "hbbs.exe")
	case "linux":
		hbbr = filepath.Join(dir, "hbbr")
		hbbs = filepath.Join(dir, "hbbs")
	default:
		// now, rustdesk-server only support windows and linux.
	}

	return hbbr, hbbs
}

func StartServer() (bool, error) {
	hbbr, hbbs := GetRustdeskServerBin()
	if strings.TrimSpace(hbbr) == "" || strings.TrimSpace(hbbs) == "" {
		return false, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
	if err := os.MkdirAll(serverBinDir, 0755); err != nil {
		return false, err
	}

	pHbbr := exec.Command(hbbr)
	pHbbr.Dir = serverBinDir
	if err := pHbbr.Start(); err != nil {
		return false, fmt.Errorf("hbbr start error: %w", err)
	}
	if err := writePidFile(hbbrPidFile, pHbbr.Process.Pid); err != nil {
		_ = pHbbr.Process.Kill()
		return false, fmt.Errorf("write hbbr pid file error: %w", err)
	}

	pHbbs := exec.Command(hbbs)
	pHbbs.Dir = serverBinDir
	if err := pHbbs.Start(); err != nil {
		_ = pHbbr.Process.Kill()
		_ = os.Remove(hbbrPidFile)
		return false, fmt.Errorf("hbbs start error: %w", err)
	}
	if err := writePidFile(hbbsPidFile, pHbbs.Process.Pid); err != nil {
		_ = pHbbr.Process.Kill()
		_ = pHbbs.Process.Kill()
		_ = os.Remove(hbbrPidFile)
		return false, fmt.Errorf("write hbbs pid file error: %w", err)
	}
	return true, nil
}

func writePidFile(pidFile string, pid int) error {
	return os.WriteFile(pidFile, []byte(strconv.Itoa(pid)), 0644)
}

func readServerPID() (hbbrPid, hbbsPid int) {
	hbbrPidBytes, err := os.ReadFile(hbbrPidFile)
	if err != nil {
		return -1, -1
	}
	hbbrPid, err = strconv.Atoi(strings.TrimSpace(string(hbbrPidBytes)))
	if err != nil {
		return -1, -1
	}

	hbbsPidBytes, err := os.ReadFile(hbbsPidFile)
	if err != nil {
		return -1, -1
	}

	hbbsPid, err = strconv.Atoi(strings.TrimSpace(string(hbbsPidBytes)))
	if err != nil {
		return -1, -1
	}

	return hbbrPid, hbbsPid
}

func StopServer() bool {
	hbbrPid, hbbsPid := readServerPID()
	stoppedHbbr := killProcessByPID(hbbrPid)
	stoppedHbbs := killProcessByPID(hbbsPid)
	_ = os.Remove(hbbrPidFile)
	_ = os.Remove(hbbsPidFile)
	return stoppedHbbr || stoppedHbbs
}

func killProcessByPID(pid int) bool {
	if pid <= 0 {
		return false
	}
	proc, err := process.NewProcess(int32(pid))
	if err != nil || proc == nil {
		return false
	}
	if err = proc.Kill(); err != nil {
		return false
	}
	return true
}

func Status() (hbbrIsRunning, hbbsIsRunning bool) {
	hbbrPid, hbbsPid := readServerPID()
	hbbrIsRunning = isProcessRunning(hbbrPid)
	hbbsIsRunning = isProcessRunning(hbbsPid)
	return hbbrIsRunning, hbbsIsRunning
}

func isProcessRunning(pid int) bool {
	if pid <= 0 {
		return false
	}
	proc, err := process.NewProcess(int32(pid))
	if err != nil || proc == nil {
		return false
	}
	running, err := proc.IsRunning()
	return err == nil && running
}

func Keys() (public, private string) {
	publicKeysBytes, err := os.ReadFile(filepath.Join(serverBinDir, "id_ed25519.pub"))
	if err != nil {
		return "", ""
	}

	privateKeysBytes, err := os.ReadFile(filepath.Join(serverBinDir, "id_ed25519"))
	if err != nil {
		return "", ""
	}

	public = string(publicKeysBytes)

	private = string(privateKeysBytes)

	return public, private
}
