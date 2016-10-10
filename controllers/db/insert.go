package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Material : 管理する資材項目
type Material struct {
	Fuel  int // 燃料
	Ammun int // 弾薬
	Steel int // 鋼材
	Baux  int // ボーキサイト
	Buck  int // 高速修復材
	Dmat  int // 開発資材
	Screw int // 改修資材
}

// Insert : 新規データ登録
func Insert(data Material, date string) error {
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
		resources(fuel, ammunition, steel, bauxite, target_date)
	VALUES
		(?, ?, ?, ?, ?)
	`
	if _, err = tx.Exec(insertQuery, data.Fuel, data.Ammun, data.Steel, data.Baux, date); err != nil {
		log.Println("資源登録エラー")
		tx.Rollback()
		return err
	}

	//　開発資材の登録
	insertQuery = `
	INSERT INTO
		materials(bucket, dMaterial, screw, target_date)
	VALUES
		(?, ?, ?, ?)
	`
	if _, err = tx.Exec(insertQuery, data.Buck, data.Dmat, data.Screw, date); err != nil {
		log.Println("開発資材登録エラー")
		tx.Rollback()
		return err
	}

	log.Println("ここまできたよ")

	// トランザクションの終了
	tx.Commit()

	return nil
}
