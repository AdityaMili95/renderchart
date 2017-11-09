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
	ID := ""
	MsgType := ""

	tempToken := r.Form["token"]
	tempUserId := r.Form["userId"]
	tempRoomId := r.Form["roomId"]
	tempGroupId := r.Form["groupId"]
	day := r.Form["day"]
	month := r.Form["month"]
	year := r.Form["year"]
	period := r.Form["period"]

	if period == nil || day == nil || month == nil || year == nil {
		return
	}

	if tempToken != nil {
		token = tempToken[0]
	}

	if tempUserId != nil {
		ID = tempUserId[0]
		MsgType = "User"
	} else if tempRoomId != nil {
		ID = tempRoomId[0]
		MsgType = "Room"
	} else if tempGroupId != nil {
		ID = tempGroupId[0]
		MsgType = "Group"
	} else {
		return
	}

	tempt, err := template.New("html_capture.html").ParseFiles("html_capture.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tempt.Execute(w, map[string]interface{}{
		"ReqToken": token,
		"ID":       ID,
		"MsgType":  MsgType,
		"Day":      day,
		"Month":    month,
		"Year":     year,
		"Period":   period,
	})
}
