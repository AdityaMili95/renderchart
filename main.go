package main

import (
	"fmt"
	"html/template"
	"net/http"
	//"os"
)

func main() {
	http.HandleFunc("/render_chart", testong)
	//port := os.Getenv("PORT")
	//addr := fmt.Sprintf(":%s", port)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(":8007", nil)
}

func testong(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	}

	ID := ""
	MsgType := ""
	day := ""
	month := ""
	year := ""
	period := ""

	tempUserId := r.Form["userId"]
	tempRoomId := r.Form["roomId"]
	tempGroupId := r.Form["groupId"]
	tempDay := r.Form["day"]
	tempMonth := r.Form["month"]
	tempYear := r.Form["year"]
	tempPeriod := r.Form["period"]

	if tempPeriod == nil || tempDay == nil || tempMonth == nil || tempYear == nil {
		return
	}

	day = tempDay[0]
	month = tempMonth[0]
	year = tempYear[0]
	period = tempPeriod[0]

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
		"ID":      ID,
		"MsgType": MsgType,
		"Day":     day,
		"Month":   month,
		"Year":    year,
		"Period":  period,
	})
}
