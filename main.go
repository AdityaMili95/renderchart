package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/process", testong)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func testong(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tempt, err := template.New("html_capture.html").ParseFiles("html_capture.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tempt.Execute(w, map[string]interface{}{
		"token": "hahahah",
	})
}
