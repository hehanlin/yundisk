package conf

import "os"

var (
	WORKDIR        = ""
	STORE_FILE_DIR = "/files/"
	HTML_DIR       = "/static/html/"
	HTML_SUFFIX    = ".html"
)

func init() {
	var err error
	WORKDIR, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	STORE_FILE_DIR = WORKDIR + STORE_FILE_DIR
	HTML_DIR = WORKDIR + HTML_DIR
}
