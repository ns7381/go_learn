package main

import (
	"fmt"
	"net/http"
	"html/template"
	"os"
	"io"
)

func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/file", file)
	http.ListenAndServe("localhost:8080", nil)

}

func index(rw http.ResponseWriter, req *http.Request)  {
	t, _ := template.ParseFiles("/Users/nathan/go/go_path/src/go_learn/src/app/backend/index.html")
	t.Execute(rw, nil)
}

func login(rw http.ResponseWriter, req *http.Request)  {
	req.ParseForm()
	if len(req.Form["name"])> 0 {
		fmt.Fprint(rw, "hello", req.Form["name"][0])
	} else {
		fmt.Fprint(rw, "hello world")
	}

}
func file(rw http.ResponseWriter, req *http.Request)  {
	if req.Method == "POST" {
		req.ParseMultipartForm(32 << 20)
		file, header, e := req.FormFile("file")
		if e != nil {
			fmt.Println(e)
			return
		}
		defer file.Close()
		openFile, _ := os.OpenFile("./"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer openFile.Close()
		io.Copy(openFile, file)
	} else {
		fmt.Fprint(rw, "hello world")
	}

}

