package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Insert : 新規データ登録
func Insert(data Resources, date string) error {
	log.Println("Insert Function")

	log.Println(data)
	log.Println(date)

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
	insertQuery := `
		INSERT INTO
			resources(
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
				target_date)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`
	if _, err = tx.Exec(
		insertQuery,
		data.Level,
		data.Fuel,
		data.Ammun,
		data.Steel,
		data.Baux,
		data.Buck,
		data.Dmat,
		data.Screw,
		data.Bann,
		data.WinSo,
		data.DefSo,
		data.Expe,
		data.SuEx,
		data.WinEx,
		data.DefEx,
		data.Veter,
		data.Rank,
		date); err != nil {
		log.Println("資源登録エラー")
		tx.Rollback()
		return err
	}

	// //　開発資材の登録
	// insertQuery = `
	// 	INSERT INTO
	// 		materials(
	// 			bucket,
	// 			dMaterial,
	// 			screw,
	// 			banner,
	// 			target_date)
	// 	VALUES
	// 		(?, ?, ?, ?, ?)
	// 	`
	// if _, err = tx.Exec(
	// 	insertQuery,
	// 	data.Materials.Buck,
	// 	data.Materials.Dmat,
	// 	data.Materials.Screw,
	// 	data.Materials.Bann,
	// 	date); err != nil {
	// 	log.Println("開発資材登録エラー")
	// 	tx.Rollback()
	// 	return err
	// }

	// insertQuery = `
	// 	Insert INTO
	// 		record(
	// 			winning_sortie,
	// 			defeatting_sortie,
	// 			expedition,
	// 			successs_expedition,
	// 			winning_exercises,
	// 			defeatting_exercises,
	// 			veterans,
	// 			ranking,
	// 			target_date)
	// 	VALUES
	// 		(?, ?, ?, ?, ?, ?, ?, ?, ?)
	// `
	// if _, err = tx.Exec(
	// 	insertQuery,
	// 	data.Record.WinSo,
	// 	data.Record.DefSo,
	// 	data.Record.Expe,
	// 	data.Record.SuEx,
	// 	data.Record.WinEx,
	// 	data.Record.DefEx,
	// 	data.Record.Veter,
	// 	data.Record.Rank,
	// 	date); err != nil {
	// 	log.Println("戦績登録エラー")
	// 	tx.Rollback()
	// 	return err
	// }

	log.Println("ここまできたよ")

	// トランザクションの終了
	tx.Commit()

	return nil
}
