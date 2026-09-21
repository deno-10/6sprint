package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	indexFile   = "index.html"
	formFileKey = "myFile"
	maxMemory   = 10 << 20
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, indexFile)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile(formFileKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := saveResult(result, header.Filename); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := io.WriteString(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func saveResult(result, srcName string) (err error) {
	stamp := strings.NewReplacer(" ", "_", ":", "-").Replace(time.Now().UTC().String())
	name := stamp + filepath.Ext(srcName)

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); err == nil {
			err = cerr
		}
	}()

	_, err = f.WriteString(result)

	return err
}
