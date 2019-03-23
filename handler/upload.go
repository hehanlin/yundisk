package handler

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hehanlin/fileupload/util"
	"github.com/pkg/errors"
)

var (
	debug  = true
	logger = log.New(os.Stdout, "", log.Lshortfile|log.LstdFlags)
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	var err error
	errorLog := func(msg string) {
		err := errors.Wrap(err, msg)
		logger.Panicln(err)
		if debug == true {
			w.Write([]byte(err.Error()))
		}
	}

	if r.Method == "GET" {
		err = util.Template(w, "upload")
		if err != nil {
			errorLog("")
			return
		}
	} else {
		err = util.StoreFile(r, "upload_file")
		if err != nil {
			errorLog("")
			return
		}
		io.WriteString(w, "upload success")
	}
}
