package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	// Check if test mode
	if len(os.Args) > 1 && os.Args[1] == "test" {
		runTest()
		return
	}

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
			formatCost(todayData.CostUSD))
	}

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
	// Set icon (you can use any icon file)
	systray.SetIcon(getIcon())
	systray.SetTitle("Claude Usage")
	systray.SetTooltip("Claude Usage Monitor")

	// Add menu items
	mRefresh := systray.AddMenuItem("Refresh", "Refresh data")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Initial data load
	go updateMenuBar()

	// Handle menu events
	go func() {
		for {
			select {
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
		log.Printf("Error getting daily usage: %v", err)
		systray.SetTitle("Error")
		return
	}

	todayData, err := getTodayData(dailyResp)
	if err != nil {
		log.Printf("Error getting today's data: %v", err)
		systray.SetTitle("No data")
		return
	}

	// Format title for menu bar
	title := fmt.Sprintf("%s - %s",
		formatTokens(todayData.TotalTokens),
		formatCost(todayData.CostUSD))

	systray.SetTitle(title)
}

// getIcon returns a simple icon (you can replace with actual icon file)
func getIcon() []byte {
	// Return a simple 16x16 icon data
	// For now, return empty to use default icon
	return []byte{}
}
