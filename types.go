package main

// DailyResponse represents the response from ccusage daily --json
type DailyResponse struct {
	Daily  []DailyData  `json:"daily"`
	Totals DailySummary `json:"totals"`
}

// DailyData represents a single day's usage data
type DailyData struct {
	Date                string   `json:"date"`
	ModelsUsed          []string `json:"modelsUsed"`
	InputTokens         int      `json:"inputTokens"`
	OutputTokens        int      `json:"outputTokens"`
	CacheCreationTokens int      `json:"cacheCreationTokens"`
	CacheReadTokens     int      `json:"cacheReadTokens"`
	TotalTokens         int      `json:"totalTokens"`
	TotalCost           float64  `json:"totalCost"`
}

// DailySummary represents the summary of daily data
type DailySummary struct {
	TotalInputTokens         int     `json:"totalInputTokens"`
	TotalOutputTokens        int     `json:"totalOutputTokens"`
	TotalCacheCreationTokens int     `json:"totalCacheCreationTokens"`
	TotalCacheReadTokens     int     `json:"totalCacheReadTokens"`
	TotalTokens              int     `json:"totalTokens"`
	TotalCostUSD             float64 `json:"totalCostUSD"`
}

// MonthlyResponse represents the response from ccusage monthly --json
type MonthlyResponse struct {
	Monthly []MonthlyData `json:"monthly"`
	Totals  MonthlyTotals `json:"totals"`
}

// MonthlyData represents a single month's usage data
type MonthlyData struct {
	Month               string           `json:"month"`
	InputTokens         int              `json:"inputTokens"`
	OutputTokens        int              `json:"outputTokens"`
	CacheCreationTokens int              `json:"cacheCreationTokens"`
	CacheReadTokens     int              `json:"cacheReadTokens"`
	TotalTokens         int              `json:"totalTokens"`
	TotalCost           float64          `json:"totalCost"`
	ModelsUsed          []string         `json:"modelsUsed"`
	ModelBreakdowns     []ModelBreakdown `json:"modelBreakdowns"`
}

// ModelBreakdown represents usage breakdown by model
type ModelBreakdown struct {
	ModelName           string  `json:"modelName"`
	InputTokens         int     `json:"inputTokens"`
	OutputTokens        int     `json:"outputTokens"`
	CacheCreationTokens int     `json:"cacheCreationTokens"`
	CacheReadTokens     int     `json:"cacheReadTokens"`
	Cost                float64 `json:"cost"`
}

// MonthlyTotals represents the totals of monthly data
type MonthlyTotals struct {
	InputTokens         int     `json:"inputTokens"`
	OutputTokens        int     `json:"outputTokens"`
	CacheCreationTokens int     `json:"cacheCreationTokens"`
	CacheReadTokens     int     `json:"cacheReadTokens"`
	TotalCost           float64 `json:"totalCost"`
	TotalTokens         int     `json:"totalTokens"`
}
