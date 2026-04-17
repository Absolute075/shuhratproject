package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type healthResponse struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	distDir := os.Getenv("FRONTEND_DIST")
	if distDir == "" {
		distDir = filepath.FromSlash("../frontend/dist")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(healthResponse{Status: "ok", Time: time.Now().UTC()})
	})

	spa := newSPAHandler(distDir)
	mux.Handle("/", spa)

	h := withCORS(mux)

	addr := ":" + port
	log.Printf("listening on %s", addr)
	log.Printf("serving frontend dist (if exists) from %s", distDir)
	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatal(err)
	}
}

type spaHandler struct {
	distDir string
	fs      http.Handler
}

func newSPAHandler(distDir string) http.Handler {
	return &spaHandler{
		distDir: distDir,
		fs:      http.FileServer(http.Dir(distDir)),
	}
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !dirExists(h.distDir) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("frontend build not found. Run the frontend dev server (Vite) or build it to frontend/dist."))
		return
	}

	path := r.URL.Path
	if path == "/" {
		serveIndex(w, r, h.distDir)
		return
	}

	clean := filepath.Clean(filepath.FromSlash(strings.TrimPrefix(path, "/")))
	candidate := filepath.Join(h.distDir, clean)

	if fileExists(candidate) {
		h.fs.ServeHTTP(w, r)
		return
	}

	serveIndex(w, r, h.distDir)
}

func serveIndex(w http.ResponseWriter, r *http.Request, distDir string) {
	indexPath := filepath.Join(distDir, "index.html")
	if fileExists(indexPath) {
		http.ServeFile(w, r, indexPath)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("index.html not found in frontend dist."))
}

func dirExists(path string) bool {
	st, err := os.Stat(path)
	if err != nil {
		return false
	}
	return st.IsDir()
}

func fileExists(path string) bool {
	st, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !st.IsDir()
}

func withCORS(next http.Handler) http.Handler {
	allowedOrigins := map[string]bool{
		"http://localhost:5173": true,
		"http://127.0.0.1:5173": true,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && allowedOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
