package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Update : 指定されたdataIdのカラムを更新する
func Update(dataID int, data Resources, date string) error {
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
	// 	data.Buck,
	// 	data.Dmat,
	// 	data.Screw,
	// 	data.Bann,
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
	// 	data.WinSo,
	// 	data.DefSo,
	// 	data.Expe,
	// 	data.SuEx,
	// 	data.WinEx,
	// 	data.DefEx,
	// 	data.Veter,
	// 	data.Rank,
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
