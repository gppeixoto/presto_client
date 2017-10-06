package presto_client

import "time"

type queryResult struct {
	NextURI          string `json:"nextUri"`
	InfoURI          string `json:"infoUri"`
	PartialCancelURI string `json:"partialCancelUri"`
	Stats            struct {
		Scheduled      bool `json:"scheduled"`
		TotalSplits    int  `json:"totalSplits"`
		RunningSplits  int  `json:"runningSplits"`
		CPUTimeMillis  int  `json:"cpuTimeMillis"`
		QueuedSplits   int  `json:"queuedSplits"`
		ProcessedBytes int  `json:"processedBytes"`
		RootStage      struct {
			TotalSplits   int `json:"totalSplits"`
			RunningSplits int `json:"runningSplits"`
			SubStages     []struct {
				TotalSplits     int           `json:"totalSplits"`
				RunningSplits   int           `json:"runningSplits"`
				SubStages       []interface{} `json:"subStages"`
				CPUTimeMillis   int           `json:"cpuTimeMillis"`
				QueuedSplits    int           `json:"queuedSplits"`
				ProcessedBytes  int           `json:"processedBytes"`
				UserTimeMillis  int           `json:"userTimeMillis"`
				State           string        `json:"state"`
				CompletedSplits int           `json:"completedSplits"`
				Done            bool          `json:"done"`
				StageID         string        `json:"stageId"`
				Nodes           int           `json:"nodes"`
				ProcessedRows   int           `json:"processedRows"`
				WallTimeMillis  int           `json:"wallTimeMillis"`
			} `json:"subStages"`
			CPUTimeMillis   int    `json:"cpuTimeMillis"`
			QueuedSplits    int    `json:"queuedSplits"`
			ProcessedBytes  int    `json:"processedBytes"`
			UserTimeMillis  int    `json:"userTimeMillis"`
			State           string `json:"state"`
			CompletedSplits int    `json:"completedSplits"`
			Done            bool   `json:"done"`
			StageID         string `json:"stageId"`
			Nodes           int    `json:"nodes"`
			ProcessedRows   int    `json:"processedRows"`
			WallTimeMillis  int    `json:"wallTimeMillis"`
		} `json:"rootStage"`
		UserTimeMillis  int    `json:"userTimeMillis"`
		State           string `json:"state"`
		CompletedSplits int    `json:"completedSplits"`
		Nodes           int    `json:"nodes"`
		Queued          bool   `json:"queued"`
		ProcessedRows   int    `json:"processedRows"`
		WallTimeMillis  int    `json:"wallTimeMillis"`
	} `json:"stats"`
	Data    [][]string `json:"data"`
	ID      string     `json:"id"`
	Columns []struct {
		TypeSignature struct {
			Arguments []struct {
				Kind  string `json:"kind"`
				Value int64  `json:"value"`
			} `json:"arguments"`
			TypeArguments    []interface{} `json:"typeArguments"`
			RawType          string        `json:"rawType"`
			LiteralArguments []interface{} `json:"literalArguments"`
		} `json:"typeSignature"`
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"columns"`
}

type apiResponse struct {
	ID      string `json:"id"`
	InfoURI string `json:"infoUri"`
	NextURI string `json:"nextUri"`
	Stats   struct {
		State           string `json:"state"`
		Queued          bool   `json:"queued"`
		Scheduled       bool   `json:"scheduled"`
		Nodes           int    `json:"nodes"`
		TotalSplits     int    `json:"totalSplits"`
		QueuedSplits    int    `json:"queuedSplits"`
		RunningSplits   int    `json:"runningSplits"`
		CompletedSplits int    `json:"completedSplits"`
		UserTimeMillis  int    `json:"userTimeMillis"`
		CPUTimeMillis   int    `json:"cpuTimeMillis"`
		WallTimeMillis  int    `json:"wallTimeMillis"`
		ProcessedRows   int    `json:"processedRows"`
		ProcessedBytes  int    `json:"processedBytes"`
	} `json:"stats"`
}

type nodeResponse []struct {
	Age                  string    `json:"age"`
	LastRequestTime      time.Time `json:"lastRequestTime"`
	LastResponseTime     time.Time `json:"lastResponseTime"`
	RecentFailureRatio   float64   `json:"recentFailureRatio"`
	RecentFailures       float64   `json:"recentFailures"`
	RecentFailuresByType struct {
	} `json:"recentFailuresByType"`
	RecentRequests  float64 `json:"recentRequests"`
	RecentSuccesses float64 `json:"recentSuccesses"`
	URI             string  `json:"uri"`
}
