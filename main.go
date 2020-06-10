// Copyright © 2020. Drew Lee. All rights reserved

package main

import (
	"KarlMalone/internal/controller"
	"html/template"
	"log"
	"net/http"
)

// 实现注册模版
func registerView() {
	tpl, err := template.ParseGlob("./internal/view/**/*")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, v := range tpl.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(writer http.ResponseWriter, request *http.Request) {
			_ = tpl.ExecuteTemplate(writer, tplName, nil)
		})
	}
}

func main() {
	http.HandleFunc("/member/login", controller.MemberLogin)
	http.HandleFunc("/member/register", controller.MemberRegister)
	http.HandleFunc("/contact/add_friend", controller.AddFriend)
	http.HandleFunc("/contact/load_friend", controller.LoadFriend)
	http.HandleFunc("/contact/load_group", controller.LoadGroup)
	http.HandleFunc("/contact/create_group", controller.CreateGroup)
	http.HandleFunc("/contact/join_group", controller.JoinGroup)
	http.HandleFunc("/chat", controller.Chat)
	http.HandleFunc("/attach/upload", controller.FileUpload)

	//提供静态资源目录支持
	http.Handle("/assets/", http.FileServer(http.Dir(".")))
	http.Handle("/resource/", http.FileServer(http.Dir(".")))
	registerView()
	log.Fatal(http.ListenAndServe(":5288", nil))
}
