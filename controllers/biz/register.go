package biz

import (
	"fmt"
	"log"
	"strconv"

	"MaterialManager/models/db"
)

// 登録データ一覧
const (
	ID    = "dataID"               // データID
	Level = "level"                // 提督レベル
	Fuel  = "fuel"                 // 燃料
	Ammun = "ammunition"           // 弾薬
	Steel = "steel"                // 鋼材
	Baux  = "bauxite"              // ボーキサイト
	Buck  = "bucket"               // 高速修復材
	DMat  = "dMaterial"            // 開発資材
	Bann  = "banner"               // 高速建造材
	Screw = "screw"                // 改修資材
	WinSo = "winning_sortie"       // 出撃の勝数
	DefSo = "defeatting_sortie"    // 出撃の敗数
	Expe  = "expedition"           // 遠征の回数
	SuEx  = "successs_expedition"  // 遠征の成功数
	WinEx = "winning_exercises"    // 演習の勝数
	DefEx = "defeatting_exercises" // 演習の敗数
	Veter = "veterans"             // 戦果
	Rank  = "ranking"              // 順位
	Date  = "date"                 // 対象日
)

// Regist : 登録情報の設定
func Regist(data func(string) string) error {
	log.Println("Regist Function")

	dbRes, err := configureResources(data)
	if err != nil {
		return err
	}
	log.Println(dbRes)

	date := data(Date)

	// DBインサート可能かを判定
	if err = db.CountColumn(date); err != nil {
		return err
	}

	// DBインサート
	if err = db.Insert(dbRes, date); err != nil {
		return err
	}

	// 返り値はエラー型
	// 登録処理に問題がなければnilを返す
	return nil
}

// Update : 更新情報の設定
func Update(data func(string) string) error {
	log.Println("Update Function")

	dbRes, err := configureResources(data)
	if err != nil {
		return err
	}
	log.Println(dbRes)

	dataID, err := strconv.Atoi(data(ID))
	if err != nil {
		err := fmt.Errorf("Error: %s", "更新IDがありません")
		log.Println(err)
		return err
	}
	log.Println("更新ID", dataID)

	date := data(Date)

	// DBインサート
	if err = db.Update(dataID, dbRes, date); err != nil {
		log.Println("更新エラー", err)
		return err
	}
	// 登録処理に問題がなければnilを返す
	return nil
}

// configureResources : 資材情報の設定
func configureResources(data func(string) string) (db.Resources, error) {
	resources, err := checkForm(data, Fuel, Ammun, Steel, Baux, Buck, DMat, Bann, Screw, Level, WinSo, DefSo, Expe, SuEx, WinEx, DefEx, Veter, Rank)
	if err != nil {
		return db.Resources{}, err
	}
	dbRes := db.Resources{
		Fuel:  resources[Fuel],
		Ammun: resources[Ammun],
		Steel: resources[Steel],
		Baux:  resources[Baux],
		Buck:  resources[Buck],
		Dmat:  resources[DMat],
		Screw: resources[Screw],
		Bann:  resources[Bann],
		Level: resources[Level],
		WinSo: resources[WinSo],
		DefSo: resources[DefSo],
		Expe:  resources[Expe],
		SuEx:  resources[SuEx],
		WinEx: resources[WinEx],
		DefEx: resources[DefEx],
		Veter: resources[Veter],
		Rank:  resources[Rank],
	}
	return dbRes, nil
}

// checkForm:入力項目を走査し、全て存在すれば結果を返す。1つでも不足していればerrを返す。
func checkForm(formValue func(string) string, checks ...string) (map[string]int, error) {

	result := map[string]int{}
	for _, check := range checks {
		// 入力内容を格納し、errの返り値で判定する
		value := formValue(check)
		if value == "" {
			err := fmt.Errorf("Error: %s: %s", "入力チェックエラー", check)
			log.Println(err)
			return map[string]int{}, err
		}
		numValue, err := strconv.Atoi(value)
		if err != nil {
			log.Println(err)
			return map[string]int{}, err
		}
		result[check] = numValue
	}
	log.Println("入力チェック成功: ", result)
	return result, nil
}
