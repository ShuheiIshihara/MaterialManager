package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// MaterialList : 管理する資材項目
type MaterialList struct {
	Resources Resources
	Materials Materials
	Record    Record
}

// Resources : 資源
type Resources struct {
	Fuel  int // 燃料
	Ammun int // 弾薬
	Steel int // 鋼材
	Baux  int // ボーキサイト
}

// Materials : 資材
type Materials struct {
	Buck  int // 高速修復材
	Dmat  int // 開発資材
	Screw int // 改修資材
	Bann  int // 高速建造材
}

// Record : 戦績
type Record struct {
	WinSo int // 出撃の勝数
	DefSo int // 出撃の敗数
	Expe  int // 遠征の回数
	SuEx  int // 遠征の成功数
	WinEx int // 演習の勝数
	DefEx int // 演習の敗数
	Veter int // 戦果
	Rank  int // 順位
}

// Insert : 新規データ登録
func Insert(data MaterialList, date string) error {
	log.Println("Insert Function")

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
				fuel, 
				ammunition, 
				steel, 
				bauxite, 
				target_date)
		VALUES
			(?, ?, ?, ?, ?)
		`
	if _, err = tx.Exec(
		insertQuery,
		data.Resources.Fuel,
		data.Resources.Ammun,
		data.Resources.Steel,
		data.Resources.Baux,
		date); err != nil {
		log.Println("資源登録エラー")
		tx.Rollback()
		return err
	}

	//　開発資材の登録
	insertQuery = `
		INSERT INTO
			materials(
				bucket, 
				dMaterial, 
				screw, 
				banner, 
				target_date)
		VALUES
			(?, ?, ?, ?, ?)
		`
	if _, err = tx.Exec(
		insertQuery,
		data.Materials.Buck,
		data.Materials.Dmat,
		data.Materials.Screw,
		data.Materials.Bann,
		date); err != nil {
		log.Println("開発資材登録エラー")
		tx.Rollback()
		return err
	}

	insertQuery = `
		Insert INTO
			record(
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
			(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	if _, err = tx.Exec(
		insertQuery,
		data.Record.WinSo,
		data.Record.DefSo,
		data.Record.Expe,
		data.Record.SuEx,
		data.Record.WinEx,
		data.Record.DefEx,
		data.Record.Veter,
		data.Record.Rank,
		date); err != nil {
		log.Println("戦績登録エラー")
		tx.Rollback()
		return err
	}

	log.Println("ここまできたよ")

	// トランザクションの終了
	tx.Commit()

	return nil
}
