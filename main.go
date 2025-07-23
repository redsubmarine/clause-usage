package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
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
