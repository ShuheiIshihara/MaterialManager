package controllers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

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
		layout := "2006-01-"
		currentMonth := time.Now().Format(layout)
		// トップページ的な何か

		funcMap := template.FuncMap{
			"day": func(text string) string { return text[8:10] },
		}

		// tmpl := template.Must(template.ParseFiles("/Users/sishihara/go/src/MaterialManager/public/html/index.html")).Funcs(funcMap)
		tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("/Users/sishihara/go/src/MaterialManager/public/html/index.html")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		// 当月のデータ取得
		// test, _ := biz.Select(currentMonth + "01")
		data := biz.GenerateSelectData(biz.Select(currentMonth + "01"))
		// log.Println("=================")
		// log.Println(data)
		// log.Println("=================")
		if err := tmpl.Execute(w, data); err != nil {
			log.Println(err)
		}
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

		tmpl := template.Must(template.ParseFiles("/Users/sishihara/go/src/MaterialManager/public/html/create.html"))
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
		var err error
		if err = biz.Regist(formValue); err != nil {
			log.Println("登録エラー: ", err)
		}
		// ヘッダにjsonとアクセスコントロールの設定を追加
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(biz.GenerateInsertData(err))
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
