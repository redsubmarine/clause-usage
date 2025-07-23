package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// getDailyUsage executes ccusage daily --json and returns parsed data
func getDailyUsage() (*DailyResponse, error) {
	cmd := exec.Command("ccusage", "daily", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute ccusage daily: %w", err)
	}

	var response DailyResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return nil, fmt.Errorf("failed to parse daily JSON: %w", err)
	}

	return &response, nil
}

// getMonthlyUsage executes ccusage monthly --json and returns parsed data
func getMonthlyUsage() (*MonthlyResponse, error) {
	cmd := exec.Command("ccusage", "monthly", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute ccusage monthly: %w", err)
	}

	var response MonthlyResponse
	if err := json.Unmarshal(output, &response); err != nil {
		return nil, fmt.Errorf("failed to parse monthly JSON: %w", err)
	}

	return &response, nil
}

// getTodayData returns today's usage data from daily response
func getTodayData(dailyResp *DailyResponse) (*DailyData, error) {
	today := time.Now().Format("2006-01-02")

	for _, data := range dailyResp.Data {
		if data.Date == today {
			return &data, nil
		}
	}

	return nil, fmt.Errorf("no data found for today (%s)", today)
}

// formatTokens formats token count with K/M suffix
func formatTokens(tokens int) string {
	if tokens >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(tokens)/1000000)
	} else if tokens >= 1000 {
		return fmt.Sprintf("%.1fK", float64(tokens)/1000)
	}
	return fmt.Sprintf("%d", tokens)
}

// formatCost formats cost with dollar sign
func formatCost(cost float64) string {
	return fmt.Sprintf("$%.2f", cost)
}

// formatModelName shortens model name for display
func formatModelName(modelName string) string {
	if strings.Contains(modelName, "claude-sonnet") {
		return "sonnet-4"
	} else if strings.Contains(modelName, "claude-opus") {
		return "opus-4"
	} else if strings.Contains(modelName, "claude-haiku") {
		return "haiku-4"
	}
	return modelName
}
