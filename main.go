package main

import (
	"image"
	"log"
	"log/slog"
	"net/http"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func main() {
	http.HandleFunc("/v1/compress", compressHandler)

	addr := ":1323"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func compressHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	compressToWebP(w, r)
	duration := time.Since(startTime)

	slog.Info("received request", "path", r.URL.Path, "method", r.Method, "remote_addr", r.RemoteAddr, "duration", duration)
}

func compressToWebP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Limit memory used for parsing form (files beyond this will be stored in temp files)
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "missing file 'image'", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Decode uploaded image
	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "failed to decode image", http.StatusBadRequest)
		return
	}

	// Encode to WebP into an in-memory buffer and return
	w.Header().Set("Content-Type", "image/webp")

	// Encode using default settings (no options struct available)
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}
	if err := webp.Encode(w, img, options); err != nil {
		http.Error(w, "failed to encode webp", http.StatusInternalServerError)
		return
	}
}
