package biz

import "MaterialManager/models/db"

// Select : 指定した月の資源情報を取得する
func Select(date string) ([]db.Resources, error) {
	result, err := db.Select(date)
	if err != nil {
		return nil, err
	}
	return result, nil
}
