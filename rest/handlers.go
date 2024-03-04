package rest

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
)

// handleSomething обрабатывает веб-запрос
func handleSomething() http.Handler {
	// thing := prepareThing()
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// используем нечто для обработки запроса
			// log.Info(r.Context(), "msg", "handleSomething")
		},
	)
}

// server func
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	if atomic.LoadInt32(&healthy) == 1 {
		_ = encode(w, r, http.StatusOK, map[string]string{"status": "OK"})
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func handleTemplate(files ...string) http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplerr = template.ParseFiles(files...)
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}

		// используем tpl
		_ = tpl.Copy()
	}
}

// receiving file from user
func HandleFileReceiver() http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseMultipartForm(300 << 20) // limit your max input length! 300MB TODO: config
			if err != nil {
				log.Print(err)
				//w.WriteHeader(http.StatusRequestEntityTooLarge)
				//return
			}

			err = r.ParseForm()
			if err != nil {
				log.Print(err)
				//	w.WriteHeader(http.StatusRequestEntityTooLarge)
				//	return
			}
			var buf bytes.Buffer
			// in your case file would be fileupload
			file, header, err := r.FormFile("file")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)

				return
			}

			defer file.Close()

			fmt.Printf("Uploaded File: %+v\n", header.Filename)
			fmt.Printf("File Size: %+v\n", header.Size)
			fmt.Printf("MIME Header: %+v\n", header.Header)

			name := strings.Split(header.Filename, ".")
			fmt.Printf("File name %s\n", name[0])
			// Copy the file data to my buffer
			io.Copy(&buf, file)
			// do something with the contents...
			// I normally have a struct defined and unmarshal into a struct, but this will
			// work as an example
			contents := buf.String()
			//fmt.Println(contents)
			// I reset the buffer in case I want to use it again
			// reduces memory allocations in more intense projects
			_ = contents
			buf.Reset()
			// do something else
			// etc write header
			// Create a temporary file within our temp-images directory that follows
			// a particular naming pattern
			// tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
			// if err != nil {
			//     fmt.Println(err)
			// }
			// defer tempFile.Close()

			// // read all of the contents of our uploaded file into a
			// // byte array
			// fileBytes, err := ioutil.ReadAll(file)
			// if err != nil {
			//     fmt.Println(err)
			// }
			// // write this byte array to our temporary file
			// tempFile.Write(fileBytes)
			// // return that we have successfully uploaded our file!
			// fmt.Fprintf(w, "Successfully Uploaded File\n")
		},
	)
}
