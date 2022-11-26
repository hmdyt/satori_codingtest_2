package satori_codingtest_2

import (
	"io"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("hello", helloWorld)
}

func helloWorld(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Hello World")
}
