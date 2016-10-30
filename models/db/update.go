package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Update : 指定されたdataIdのカラムを更新する
func Update(dataID int, data MaterialList, date string) error {
	log.Println("Update Function")

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
	updateQuery := `
		UPDATE
			resources
    SET
		  level = ?,
			fuel = ?, 
			ammunition = ?, 
			steel = ?, 
			bauxite =?, 
	    bucket = ?,
	  	dMaterial = ?,
    	screw = ?,
      banner = ?,
			winning_sortie = ?, 
			defeatting_sortie = ?, 
			expedition = ?, 
			successs_expedition = ?,
			winning_exercises = ?,
			defeatting_exercises = ?,
			veterans = ?,
			ranking = ?,
      update_date = current_time
		WHERE
          dataId = ?
		`
	if _, err = tx.Exec(
		updateQuery,
		data.Record.Level,
		data.Resources.Fuel,
		data.Resources.Ammun,
		data.Resources.Steel,
		data.Resources.Baux,
		data.Materials.Buck,
		data.Materials.Dmat,
		data.Materials.Screw,
		data.Materials.Bann,
		data.Record.WinSo,
		data.Record.DefSo,
		data.Record.Expe,
		data.Record.SuEx,
		data.Record.WinEx,
		data.Record.DefEx,
		data.Record.Veter,
		data.Record.Rank,
		dataID); err != nil {
		log.Println("資源登録エラー")
		tx.Rollback()
		return err
	}

	// //　開発資材の登録
	// updateQuery = `
	// 	UPDATE
	//       	materials
	//       SET
	//         bucket = ?,
	//         dMaterial = ?,
	//   	    screw = ?,
	//       	banner = ?,
	//   	    target_date = ?,
	//           update_date = current_time
	//       WHERE
	//           dataId = ?
	// 	`
	// if _, err = tx.Exec(
	// 	updateQuery,
	// 	data.Materials.Buck,
	// 	data.Materials.Dmat,
	// 	data.Materials.Screw,
	// 	data.Materials.Bann,
	// 	date,
	// 	dataID); err != nil {
	// 	log.Println("開発資材登録エラー")
	// 	tx.Rollback()
	// 	return err
	// }

	// updateQuery = `
	// 	UPDATE
	// 		record
	// 	SET
	// 		winning_sortie = ?,
	// 		defeatting_sortie = ?,
	// 		expedition = ?,
	// 		successs_expedition = ?,
	// 		winning_exercises = ?,
	// 		defeatting_exercises = ?,
	// 		veterans = ?,
	// 		ranking = ?,
	// 		target_date = ?,
	//           update_date = current_time
	//       WHERE
	//           dataId = ?
	// `
	// if _, err = tx.Exec(
	// 	updateQuery,
	// 	data.Record.WinSo,
	// 	data.Record.DefSo,
	// 	data.Record.Expe,
	// 	data.Record.SuEx,
	// 	data.Record.WinEx,
	// 	data.Record.DefEx,
	// 	data.Record.Veter,
	// 	data.Record.Rank,
	// 	date,
	// 	dataID); err != nil {
	// 	log.Println("戦績登録エラー")
	// 	tx.Rollback()
	// 	return err
	// }

	log.Println("ここまできたよ")

	// トランザクションの終了
	tx.Commit()

	return nil
}
