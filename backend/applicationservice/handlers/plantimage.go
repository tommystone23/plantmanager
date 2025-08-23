package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
)

func PlantImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
    	return
	}

	// get image from form
	file, _, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
    	return
	}
	defer file.Close()

	// detect content type
	buf := make([]byte, 512)
    _, _ = file.Read(buf)
    contentType := http.DetectContentType(buf)
    file.Seek(0, io.SeekStart)

	// read image data
	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
    	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
    	return
	}

	// encode image data and return to client
	encoded := base64.StdEncoding.EncodeToString(data)
	json.NewEncoder(w).Encode(map[string]string{
		"image": "data:" + contentType + ";base64," + encoded,
	})
}