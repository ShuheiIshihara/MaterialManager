package biz

import (
	"encoding/json"

	"MaterialManager/models/db"
)

// InsertResult : データ登録用返り値
type InsertResult struct {
	Result int    `json:"result"`
	Reason string `json:"reason"`
}

// SelectResult : データ取得用返り値
type SelectResult struct {
	Result    int
	Reason    string
	Resources []db.Resources
}

// GenerateInsertData : データ登録結果を格納するJSON生成
func GenerateInsertData(err error) []byte {
	var result int
	var reason string
	if err != nil {
		result = -1
		reason = err.Error()
	} else {
		result = 0
		reason = ""
	}
	insertResult := InsertResult{
		Result: result,
		Reason: reason,
	}
	jsonData, _ := json.Marshal(insertResult)
	return jsonData
}

// GenerateSelectData : データ取得結果を格納するJSON生成
func GenerateSelectData(resources []db.Resources, err error) SelectResult {
	var result int
	var reason string
	if err != nil {
		result = -1
		reason = err.Error()
	} else {
		result = 0
		reason = ""
	}
	selectResult := SelectResult{
		Result:    result,
		Reason:    reason,
		Resources: resources,
	}
	// jsonData, _ := json.Marshal(selectResult)
	return selectResult
}
