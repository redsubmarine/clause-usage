package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/getlantern/systray"
)

var testMode bool

func main() {
	// Check if test mode
	if len(os.Args) > 1 && os.Args[1] == "test" {
		testMode = true
		runTest()
		return
	}


	// Check if already running as daemon
	if len(os.Args) > 1 && os.Args[1] == "--daemon" {
		// Disable logging for daemon operation
		log.SetOutput(io.Discard)
		systray.Run(onReady, onExit)
		return
	}

	// Fork to background
	daemonize()
}


func daemonize() {
	// Get current executable path
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	// Start daemon process
	cmd1 := exec.Command(executable, "--daemon")
	cmd1.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
	
	// Redirect all I/O to /dev/null
	devNull, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf("Failed to open /dev/null: %v", err)
	}
	
	cmd1.Stdin = devNull
	cmd1.Stdout = devNull
	cmd1.Stderr = devNull

	if err := cmd1.Start(); err != nil {
		devNull.Close()
		log.Fatalf("Failed to start daemon process: %v", err)
	}

	// Release the process immediately so it becomes orphaned
	if err := cmd1.Process.Release(); err != nil {
		log.Printf("Warning: failed to release process: %v", err)
	}
	
	devNull.Close()

	fmt.Printf("Claude Usage started in background\n")
	fmt.Println("Use 'claude-usage test' to test the application")
	fmt.Println("To stop: pkill claude-usage")
	
	// Exit parent immediately
	os.Exit(0)
}

func runTest() {
	fmt.Println("Testing ccusage data...")

	// Test daily data
	fmt.Println("\n=== Daily Data ===")
	dailyResp, err := getDailyUsage()
	if err != nil {
		log.Fatalf("Error getting daily usage: %v", err)
	}

	todayData, err := getTodayData(dailyResp)
	if err != nil {
		log.Printf("Error getting today's data: %v", err)
	} else {
		fmt.Printf("Today (%s): %s - %s\n",
			todayData.Date,
			formatTokens(todayData.TotalTokens),
			formatCost(todayData.TotalCost))
	}

	// Test daily table
	fmt.Println("\n=== Daily Table ===")
	dailyTable := generateDailyTable(dailyResp)
	fmt.Println(dailyTable)

	// Test monthly data
	fmt.Println("\n=== Monthly Data ===")
	monthlyResp, err := getMonthlyUsage()
	if err != nil {
		log.Fatalf("Error getting monthly usage: %v", err)
	}

	table := generateMonthlyTable(monthlyResp)
	fmt.Println(table)
}

func onReady() {
	// Set title and tooltip (skip icon for now)
	systray.SetTitle("Claude Usage")
	systray.SetTooltip("Claude Usage Monitor")

	// Add menu items
	mShowDaily := systray.AddMenuItem("Show Daily Data", "Show daily usage table")
	mShowMonthly := systray.AddMenuItem("Show Monthly Data", "Show monthly usage table")
	systray.AddSeparator()
	mRefresh := systray.AddMenuItem("Refresh", "Refresh data")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Initial data load
	go updateMenuBar()

	// Handle menu events
	go func() {
		for {
			select {
			case <-mShowDaily.ClickedCh:
				go showDailyData()
			case <-mShowMonthly.ClickedCh:
				go showMonthlyData()
			case <-mRefresh.ClickedCh:
				go updateMenuBar()
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

	// Auto refresh every 5 minutes
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				go updateMenuBar()
			}
		}
	}()
}

func onExit() {
	// Cleanup code here
}

func updateMenuBar() {
	dailyResp, err := getDailyUsage()
	if err != nil {
		if testMode {
			log.Printf("Error getting daily usage: %v", err)
		}
		systray.SetTitle("Claude Usage")
		return
	}

	todayData, err := getTodayData(dailyResp)
	if err != nil {
		if testMode {
			log.Printf("Error getting today's data: %v", err)
		}
		systray.SetTitle("Claude Usage")
		return
	}

	// Format title for menu bar
	title := fmt.Sprintf("%s - %s",
		formatTokens(todayData.TotalTokens),
		formatCost(todayData.TotalCost))

	systray.SetTitle(title)
}

func showDailyData() {
	showDailyPopover()
}

func showMonthlyData() {
	showPopover()
}

// getIcon returns a simple icon (you can replace with actual icon file)
// func getIcon() []byte {
// 	// Return a simple 16x16 icon data
// 	// For now, return empty to use default icon
// 	return []byte{}
// }
