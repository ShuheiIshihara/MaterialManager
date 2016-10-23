package db

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
