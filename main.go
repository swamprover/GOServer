package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func run1c() {
	cmd := exec.Command(`C:\Program Files (x86)\1cv8\8.3.8.2442\bin\1cv8s.exe`)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	err := cmd.Start()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	err = cmd.Wait()

	fmt.Printf("Command finished with error: %v\n", err)
	fmt.Printf("Command finished with output: %v\n", buf.String())
}

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //анализ аргументов,
	fmt.Println(r.Form) // ввод информации о форме на стороне сервера
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Printf("key:", k)
		fmt.Printf("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "<h1> Hello word <h1>")
	run1c()
}

func main() {
	//run1c(`C:\Program Files (x86)\1cv8\8.3.8.2442\bin\1cv8s.exe`)
	http.HandleFunc("/1", hello)
	log.Fatal(http.ListenAndServe(":9111", nil))

	//router := httprouter.New()
	//router.GET("/1", run1c)
	//router.GET("/api/v1/records/:id", getRecord)
	//router.POST("/api/v1/records", addRecord)
	//router.PUT("/api/v1/records/:id", updateRecord)
	//router.DELETE("/api/v1/records/:id", deleteRecord)
	//http.ListenAndServe(":9111", router)
}
