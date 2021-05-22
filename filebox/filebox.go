package filebox

import (
	"embed"
	"fmt"
	"log"
	"net/http"
)

var _EmbedFiles embed.FS

func AddFiles(files embed.FS)  {
	_EmbedFiles = files
}

func Exists(name string) bool {
	f, err := _EmbedFiles.Open(fmt.Sprintf("statics/%s", name))
	if err != nil {
		return false
	}
	defer f.Close()

	return true
}

func GetFile(name string) ([]byte, string) {
	blob, err := _EmbedFiles.ReadFile(fmt.Sprintf("statics/%s", name))
	if err != nil {
		log.Fatal(err)
	}

	contentType := http.DetectContentType(blob)

	return blob, contentType
}
