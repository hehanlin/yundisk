package main

import (
	"net/http"

	"github.com/hehanlin/fileupload/handler"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadFile)
	//http.HandleFunc("/file/upload/success", nil)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
