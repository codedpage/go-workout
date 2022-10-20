package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"

	//"html/template"

	"net/http"
	"text/template"
)

func SaveKycRequestHandlerOut(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "2")
	http.ServeFile(w, r, "view/index.html")

}

func SaveKycRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("view/fileUp.gtpl")
		t.Execute(w, nil)
	} else {

		// 1. parse input
		r.ParseMultipartForm(10 << 20)

		// 2. retrieve file
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// 3. write temporary file on our server
		tempFile, err := ioutil.TempFile("uploads", "upload-*.csv")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		tempFile.Write(fileBytes)
		fmt.Println(string(fileBytes))

		//-----------------
		// Read file to byte slice
		s := bufio.NewScanner(file)
		for s.Scan() {
			fmt.Println(s.Text())
			str := strings.Split(s.Text(), ",")
			fmt.Println("------------------")

			if len(str) > 10 {
				fmt.Println(str[1], str[6], str[8], str[9], str[10])
			}
			fmt.Println("++++++++++++++++++++++++++++++++")
		}

		//------------------------

		// 4. return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")

	}

}
