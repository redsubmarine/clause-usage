package main

import (
	"fmt"
	"sort"
	"strings"
)

// generateMonthlyTable generates a formatted table from monthly data
func generateMonthlyTable(monthlyResp *MonthlyResponse) string {
	if len(monthlyResp.Monthly) == 0 {
		return "No monthly data available"
	}

	// Sort monthly data by month (newest first)
	sort.Slice(monthlyResp.Monthly, func(i, j int) bool {
		return monthlyResp.Monthly[i].Month > monthlyResp.Monthly[j].Month
	})

	var result strings.Builder

	// Header
	result.WriteString("┌──────────┬───────────────┬───────────┬───────────┬───────────────┬─────────────┬───────────────┬─────────────┐\n")
	result.WriteString("│ Month    │ Models        │     Input │    Output │  Cache Create │  Cache Read │  Total Tokens │  Cost (USD) │\n")
	result.WriteString("├──────────┼───────────────┼───────────┼───────────┼───────────────┼─────────────┼───────────────┼─────────────┤\n")

	// Monthly data rows
	for _, data := range monthlyResp.Monthly {
		models := formatModelNames(data.ModelsUsed)
		result.WriteString(fmt.Sprintf("│ %-8s │ %-13s │ %9s │ %9s │ %13s │ %11s │ %13s │ %11s │\n",
			data.Month,
			truncateString(models, 13),
			formatNumber(data.InputTokens),
			formatNumber(data.OutputTokens),
			formatNumber(data.CacheCreationTokens),
			formatNumber(data.CacheReadTokens),
			formatNumber(data.TotalTokens),
			formatCost(data.TotalCost),
		))
		result.WriteString("├──────────┼───────────────┼───────────┼───────────┼───────────────┼─────────────┼───────────────┼─────────────┤\n")
	}

	// Totals row
	result.WriteString(fmt.Sprintf("│ Total    │               │ %9s │ %9s │ %13s │ %11s │ %13s │ %11s │\n",
		formatNumber(monthlyResp.Totals.InputTokens),
		formatNumber(monthlyResp.Totals.OutputTokens),
		formatNumber(monthlyResp.Totals.CacheCreationTokens),
		formatNumber(monthlyResp.Totals.CacheReadTokens),
		formatNumber(monthlyResp.Totals.TotalTokens),
		formatCost(monthlyResp.Totals.TotalCost),
	))
	result.WriteString("└──────────┴───────────────┴───────────┴───────────┴───────────────┴─────────────┴───────────────┴─────────────┘\n")

	return result.String()
}

// truncateString truncates string to specified length
func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

// formatNumber formats numbers with comma separators
func formatNumber(num int) string {
	if num >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(num)/1000000)
	} else if num >= 1000 {
		return fmt.Sprintf("%.1fK", float64(num)/1000)
	}
	return fmt.Sprintf("%d", num)
}

// formatModelNames formats model names for display
func formatModelNames(models []string) string {
	if len(models) == 0 {
		return ""
	}

	var formatted []string
	for _, model := range models {
		formatted = append(formatted, formatModelName(model))
	}

	return strings.Join(formatted, ", ")
}
