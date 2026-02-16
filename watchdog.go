package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

const (
	// MaxMemoryMB: 50MB for quick test
	MaxMemoryMB = 50.0
	// AppCommand: Restart command
	AppCommand = "python3 app.py"
)

func main() {
	fmt.Println("SRE Watchdog started. Monitoring for app.py...")

	for {
		processes, err := process.Processes()
		if err != nil {
			fmt.Printf("Error fetching processes: %v\n", err)
			continue
		}

		for _, p := range processes {
			// Cmdline() returns the full command as a string or slice depending on version
			// convert it to a string to check for script name
			cmdline, _ := p.Cmdline()

			// Check if "app.py" is in the command line
			if strings.Contains(cmdline, "app.py") {
				memInfo, err := p.MemoryInfo()
				if err != nil {
					continue
				}

				// Calculate memory in MB
				memMB := float64(memInfo.RSS) / 1024 / 1024
				fmt.Printf("Monitoring PID %d: Memory Usage %.2f MB\n", p.Pid, memMB)

				// Self-Healing Trigger
				if memMB > MaxMemoryMB {
					fmt.Printf("CRITICAL: Memory %.2f MB exceeded limit. Restarting...\n", memMB)
					p.Kill()
					restartApp()
				}
			}
		}
		time.Sleep(2 * time.Second)
	}
}

func restartApp() {
	cmd := exec.Command("sh", "-c", AppCommand)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to restart app: %v\n", err)
	} else {
		fmt.Println("App restarted successfully. System Healed.")
	}
}