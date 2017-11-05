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

	token := ""
	userId := ""
	roomId := ""
	groupId := ""
	valid := false

	tempToken := r.Form["token"]
	tempUserId := r.Form["userId"]
	tempRoomId := r.Form["roomId"]
	tempGroupId := r.Form["groupId"]

	if tempToken != nil {
		token = tempToken[0]
	}

	if tempUserId != nil {
		userId = tempUserId[0]
		valid = true
	}

	if tempRoomId != nil {
		roomId = tempRoomId[0]
		valid = true
	}

	if tempGroupId != nil {
		groupId = tempGroupId[0]
		valid = true
	}

	if !valid {
		return
	}

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
