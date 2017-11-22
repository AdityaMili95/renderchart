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
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.ListenAndServe(addr, nil)
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
	ChartType := ""
	Data := ""

	tempUserId := r.Form["xyz"]
	tempRoomId := r.Form["yyz"]
	tempGroupId := r.Form["zyz"]
	tempDay := r.Form["day"]
	tempMonth := r.Form["month"]
	tempYear := r.Form["year"]
	tempPeriod := r.Form["period"]
	tempChartType := r.Form["chartType"]
	tempData := r.Form["tok"]

	if tempPeriod == nil || tempDay == nil || tempMonth == nil || tempYear == nil || tempChartType == nil || tempData == nil {
		return
	}

	day = tempDay[0]
	month = tempMonth[0]
	year = tempYear[0]
	period = tempPeriod[0]
	ChartType = tempChartType[0]
	Data = tempData[0]

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
		"ID":        ID,
		"MsgType":   MsgType,
		"Day":       day,
		"Month":     month,
		"Year":      year,
		"Period":    period,
		"ChartType": ChartType,
		"Data":      Data,
	})
}
