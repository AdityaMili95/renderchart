package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/render_chart", testong)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func testong(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}

	token := r.Form["token"][0]
	userId := r.Form["userId"][0]
	roomId := r.Form["roomId"][0]
	groupId := r.Form["groupId"][0]

	tempt, err := template.New("html_capture.html").ParseFiles("html_capture.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tempt.Execute(w, map[string]interface{}{
		"ReqToken": token,
		"ReqUser":  userId,
		"ReqRoom":  roomId,
		"ReqGroup": groupId,
	})
}
