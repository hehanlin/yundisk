package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/hehanlin/fileupload/conf"
)

func Template(w http.ResponseWriter, filename string) error {
	file := conf.HTML_DIR + filename + conf.HTML_SUFFIX
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func StoreFile(r *http.Request, key string) error {
	file, head, err := r.FormFile(key)
	if err != nil {
		return err
	}
	newFile, err := os.Create(conf.STORE_FILE_DIR + head.Filename)
	if err != nil {
		return err
	}
	_, err = io.Copy(newFile, file)
	return err
}
