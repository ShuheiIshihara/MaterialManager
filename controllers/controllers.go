package controllers

import (
	"html/template"
	"log"
	"net/http"
	"os"

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
	}
}

// Create : 入力画面ページ
func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == GetMethod {
		// 入力画面的な何か
		docRoot, _ := os.Getwd()
		data := map[string]string{
			"root": docRoot,
		}

		// トップページ的な何か
		tmpl := template.Must(template.ParseFiles("/Users/sishihara/go/src/MaterialManager/public/html/index.tpl"))
		w.Header().Set("Content-Type", "text/html")

		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}

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

// Update : データ更新
func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Update!!!", r.Method)
	if r.Method == PostMethod {
		formValue := r.PostFormValue // POSTデータ
		if err := biz.Update(formValue); err != nil {
			log.Println("登録エラー: ", err)
		}
		//		db.Insert(formValue)
	} else {
		log.Println("それはおかしい")
	}
}
