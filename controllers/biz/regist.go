package biz

import (
	"fmt"
	"log"
	"strconv"

	"MaterialManager/controllers/db"
)

// 登録データ一覧
const (
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

// Regist Functon
func Regist(data func(string) string) error {
	log.Println("Regist Function")

	// 入力チェック
	resources, err := checkForm(data, Fuel, Ammun, Steel, Baux)
	if err != nil {
		return err
	}
	dbRes := db.Resources{
		Fuel:  resources[Fuel],
		Ammun: resources[Ammun],
		Steel: resources[Steel],
		Baux:  resources[Baux],
	}
	log.Println(dbRes)

	materials, err := checkForm(data, Buck, DMat, Bann, Screw)
	if err != nil {
		return err
	}
	dbMat := db.Materials{
		Buck:  materials[Buck],
		Dmat:  materials[DMat],
		Screw: materials[Screw],
		Bann:  materials[Bann],
	}
	log.Println(dbMat)

	record, err := checkForm(data, WinSo, DefSo, Expe, SuEx, WinEx, DefEx, Veter, Rank)
	if err != nil {
		return err
	}
	dbRec := db.Record{
		WinSo: record[WinSo],
		DefSo: record[DefSo],
		Expe:  record[Expe],
		SuEx:  record[SuEx],
		WinEx: record[WinEx],
		DefEx: record[DefEx],
		Veter: record[Veter],
		Rank:  record[Rank],
	}
	log.Println(dbRec)
	insertData := db.MaterialList{
		Resources: dbRes,
		Materials: dbMat,
		Record:    dbRec,
	}
	log.Println(insertData)

	date := data(Date)

	// DBインサート
	if err = db.Insert(insertData, date); err != nil {
		return err
	}

	// 返り値はエラー型
	// 登録処理に問題がなければnilを返す
	return nil
}

// checkForm:入力項目を走査し、全て存在すれば結果を返す。1つでも不足していればerrを返す。
func checkForm(formValue func(string) string, checks ...string) (map[string]int, error) {

	result := map[string]int{}
	for _, check := range checks {
		// 入力内容を格納し、errの返り値で判定する
		value := formValue(check)
		if value == "" {
			err := fmt.Errorf("Error: %s", "入力チェックエラー")
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
