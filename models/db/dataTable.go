package db

// // MaterialList : 管理する資材項目
// type MaterialList struct {
// 	Resources Resources
// 	Materials Materials
// 	Record    Record
// }

// Resources : 資源
type Resources struct {
	DataID int    `json:"dataID"`               // id
	Fuel   int    `json:"fuel"`                 // 燃料
	Ammun  int    `json:"ammunition"`           // 弾薬
	Steel  int    `json:"steel"`                // 鋼材
	Baux   int    `json:"bauxite"`              // ボーキサイト
	Buck   int    `json:"bucket"`               // 高速修復材
	Dmat   int    `json:"dMaterial"`            // 開発資材
	Screw  int    `json:"screw"`                // 改修資材
	Bann   int    `json:"banner"`               // 高速建造材
	Level  int    `json:"level"`                // 提督レベル
	WinSo  int    `json:"winning_sortie"`       // 出撃の勝数
	DefSo  int    `json:"defeatting_sortie"`    // 出撃の敗数
	Expe   int    `json:"expedition"`           // 遠征の回数
	SuEx   int    `json:"successs_expedition"`  // 遠征の成功数
	WinEx  int    `json:"winning_exercises"`    // 演習の勝数
	DefEx  int    `json:"defeatting_exercises"` // 演習の敗数
	Veter  int    `json:"veterans"`             // 戦果
	Rank   int    `json:"ranking"`              // 順位
	Date   string `json:"date"`                 // 日付
}
