package db

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// CountColumn : データの存在判定
func CountColumn(date string) error {
	log.Println("Select Function")

	db, err := sql.Open("mysql", "kankore:kongo@/material")
	if err != nil {
		return err
	}
	defer db.Close()

	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return nil
	}

	// 資源の登録
	selectQuery := `
		SELECT
      COUNT(1) CNT
    FROM
      resources
    WHERE
      target_date = ?
		`
	rows, err := tx.Query(selectQuery, date)
	if err != nil {
		return err
	}
	defer rows.Close()

	var cnt int
	for rows.Next() {
		rows.Scan(&cnt)
		log.Println(cnt)
	}

	if cnt > 0 {
		err = fmt.Errorf(date)
		return err
	}

	log.Println("対象のカラム数は", cnt, "件です")

	// トランザクションの終了
	tx.Commit()

	return nil
}
