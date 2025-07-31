package main

import (
	"fmt"
	"io"
	"log"
	"os"
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

	// Disable logging for normal operation (run silently in background)
	log.SetOutput(io.Discard)

	systray.Run(onReady, onExit)
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
