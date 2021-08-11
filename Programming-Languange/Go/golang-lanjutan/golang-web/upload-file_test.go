package golangweb

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func parseTemplate(namaFile string) *template.Template {
	return template.Must(template.ParseFiles("../template/" + namaFile))
}

func uploadForm(w http.ResponseWriter, r *http.Request) {
	parseTemplate("upload-form.gohtml").Execute(w, map[string]interface{}{
		"name": "ade",
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", uploadForm)
	mux.HandleFunc("/upload", upload)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../resources"))))
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("halo pak eko")
	})

	server := http.Server{
		Addr:    ":9000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	handlePanic(err)

	fmt.Println("akhir")
}

func upload(w http.ResponseWriter, r *http.Request) {
	// akses file yang telah diupload,
	// "file-upload" disesuaikan dengan 'name' di tag input
	file, fileHeader, err := r.FormFile("file-upload")
	handlePanic(err)

	// buat lokasi penyimpanan
	fileDestination, err := os.Create("../resources/" + fileHeader.Filename)
	handlePanic(err)

	// copy
	_, err = io.Copy(fileDestination, file)
	handlePanic(err)

	// "nama-file" disesuaikan dengan 'name' di tag input
	fileName := r.PostFormValue("nama-file")
	parseTemplate("upload-success.gohtml").Execute(w, map[string]interface{}{
		"file": "/static/" + fileHeader.Filename,
		"name": fileName,
	})
}

func handlePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("nama-file", "Ade Suprapto")
	file, err := writer.CreateFormFile("file-upload", "test.jpg")
	handlePanic(err)

	fileByte, err := ioutil.ReadFile("../resources/1.jpg")
	handlePanic(err)

	file.Write(fileByte)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:9000/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	upload(recorder, request)

	bodyResponse, _ := ioutil.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
