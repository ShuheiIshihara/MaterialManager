package main

import (
	"log"
	"net/http"

	"MaterialManager/controllers"
)

// Routes : ルーティング用関数
func Routes() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public/assets"))))
	http.HandleFunc("/material", controllers.Index)         // トップページ
	http.HandleFunc("/material/create", controllers.Create) // 入力画面
	http.HandleFunc("/material/store", controllers.Store)   // 新規登録
	http.HandleFunc("/material/update", controllers.Update)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("接続エラー: ", err)
	}
}
