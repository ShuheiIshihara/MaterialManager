package db

import (
	"database/sql"
	"log"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Select : データ取得
func Select(date string) ([]Resources, error) {
	log.Println("Select Function")
	resources := []Resources{}

	db, err := sql.Open("mysql", "kankore:kongo@/material")
	if err != nil {
		log.Println("DBオープンエラー: ", err)
		return nil, err
	}
	defer db.Close()

	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		log.Println("トランザクションエラー: ", err)
		return nil, err
	}

	selectQuery := `
		SELECT
			dataId,
			level,
			fuel,
			ammunition,
			steel,
			bauxite,
			bucket,
			dMaterial,
			screw,
			banner,
			winning_sortie,
			defeatting_sortie,
			expedition,
			successs_expedition,
			winning_exercises,
			defeatting_exercises,
			veterans,
			ranking,
			target_date
		FROM
			resources
		WHERE
			target_date >= ?
		ORDER BY
			target_date`
	rows, err := tx.Query(selectQuery, date)
	if err != nil {
		log.Println("Select実行エラー: ", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var resource Resources
		if err = rows.Scan(
			&resource.DataID,
			&resource.Level,
			&resource.Fuel,
			&resource.Ammun,
			&resource.Steel,
			&resource.Baux,
			&resource.Buck,
			&resource.Dmat,
			&resource.Screw,
			&resource.Bann,
			&resource.WinSo,
			&resource.DefSo,
			&resource.Expe,
			&resource.SuEx,
			&resource.WinEx,
			&resource.DefEx,
			&resource.Veter,
			&resource.Rank,
			&resource.Date,
		); err != nil {
			log.Println("格納エラー: ", err)
			return nil, err
		}
		resources = append(resources, resource)
	}
	// log.Println("--------------------")
	// log.Println(resources)
	// log.Println("--------------------")

	// トランザクションの終了
	tx.Commit()

	return resources, nil
}

// CountColumn : データの存在判定
func CountColumn(date string) error {
	log.Println("CountColumn Function")

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

	// データカウント
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
		errStr := date + "は既に存在しています。"
		err = fmt.Errorf(errStr)
		return err
	}

	log.Println("対象のカラム数は", cnt, "件です")

	// トランザクションの終了
	tx.Commit()

	return nil
}
