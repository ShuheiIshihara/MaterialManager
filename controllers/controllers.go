package controllers

import (
	"log"
	"net/http"

	"MaterialManager/controllers/biz"
)

// リクエスト一覧
const (
	GetMethod  = "GET"
	PostMethod = "POST"
)

// Index : トップページ
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == GetMethod {
		// トップページ的な何か
	}
}

// Create : 入力画面ページ
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == GetMethod {
		// 入力画面的な何か
	}
}

// Store : 新規データ作成
func Store(w http.ResponseWriter, r *http.Request) {
	if r.Method == PostMethod {
		formValue := r.PostFormValue // POSTデータ
		if err := biz.Regist(formValue); err != nil {
			log.Println("登録エラー: ", err)
		}
		//		db.Insert(formValue)
	} else {
		log.Println("それはおかしい")
	}
}
