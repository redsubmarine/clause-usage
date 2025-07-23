package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// showPopover displays monthly data in a popover window
func showPopover() {
	monthlyResp, err := getMonthlyUsage()
	if err != nil {
		log.Printf("Error getting monthly usage: %v", err)
		return
	}

	table := generateMonthlyTable(monthlyResp)

	// Create a temporary file with the table content
	tmpFile, err := os.CreateTemp("", "claude-usage-*.txt")
	if err != nil {
		log.Printf("Error creating temp file: %v", err)
		return
	}
	defer os.Remove(tmpFile.Name())

	// Write table to temp file
	if _, err := tmpFile.WriteString(table); err != nil {
		log.Printf("Error writing to temp file: %v", err)
		return
	}
	tmpFile.Close()

	// Open the file with TextEdit in a small window
	openPopover(tmpFile.Name())
}

// openPopover opens a small window with the table content
func openPopover(filename string) {
	if runtime.GOOS == "darwin" {
		// Use AppleScript to create a small window with TextEdit
		script := fmt.Sprintf(`
			tell application "TextEdit"
				activate
				open POSIX file "%s"
				set bounds of front window to {100, 100, 800, 600}
				set visible of front window to true
			end tell
		`, filename)

		cmd := exec.Command("osascript", "-e", script)
		if err := cmd.Run(); err != nil {
			log.Printf("Error opening popover: %v", err)
			// Fallback: just open the file normally
			exec.Command("open", filename).Run()
		}
	} else {
		// For other platforms, just open the file
		exec.Command("open", filename).Run()
	}
}
