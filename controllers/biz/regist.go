package biz

import (
	"log"
	"strconv"

	"MaterialManager/controllers/db"
)

// 登録データ一覧
const (
	Fuel  = "fuel"       // 燃料
	Ammun = "ammunition" // 弾薬
	Steel = "steel"      // 鋼材
	Baux  = "bauxite"    // ボーキサイト
	Buck  = "bucket"     // 高速修復材
	DMat  = "dMaterial"  // 開発資材
	Screw = "screw"      // 改修資材
	Date  = "date"       // 対象日
)

// Regist Functon
func Regist(data func(string) string) error {
	log.Println("Regist Function")

	// DBインサート用の項目をリストにセット(入力チェックも含む)
	insertData, date, err := configureInput(data)
	if err != nil {
		return err
	}
	log.Println(insertData)

	// DBインサート
	if err = db.Insert(insertData, date); err != nil {
		return err
	}

	// 返り値はエラー型
	// 登録処理に問題がなければnilを返す
	return nil
}

// configureInput : POSTメソッドの入力値をリストへ格納する。
func configureInput(data func(string) string) (db.Material, string, error) {
	log.Println("configureInput Function")
	// 入力チェック
	// TODO: あまりにも冗長なので直す必要あり
	fuel, err := strconv.Atoi(data(Fuel))
	if err != nil {
		return db.Material{}, "", err
	}
	ammu, err := strconv.Atoi(data(Ammun))
	if err != nil {
		return db.Material{}, "", err
	}
	steel, _ := strconv.Atoi(data(Steel))
	if err != nil {
		return db.Material{}, "", err
	}
	baux, _ := strconv.Atoi(data(Baux))
	if err != nil {
		return db.Material{}, "", err
	}
	buck, _ := strconv.Atoi(data(Buck))
	if err != nil {
		return db.Material{}, "", err
	}
	dMat, _ := strconv.Atoi(data(DMat))
	if err != nil {
		return db.Material{}, "", err
	}
	screw, _ := strconv.Atoi(data(Screw))
	if err != nil {
		return db.Material{}, "", err
	}
	date := data(Date)

	return db.Material{
		Fuel:  fuel,
		Ammun: ammu,
		Steel: steel,
		Baux:  baux,
		Buck:  buck,
		Dmat:  dMat,
		Screw: screw,
	}, date, nil
}
