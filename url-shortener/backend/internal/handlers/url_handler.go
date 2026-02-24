package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/z66x/url-shortener/internal/db"
)

type URLHandler struct {
	repo *db.URLRepository
}

func NewURLHandler(repo *db.URLRepository) *URLHandler {
	return &URLHandler{repo: repo}
}

func (h *URLHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	var body struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.URL == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	shortCode := generateCode(6)

	url, err := h.repo.Create(r.Context(), body.URL, shortCode)
	if err != nil {
		http.Error(w, "failed to create short url", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"short_url":    os.Getenv("BASE_URL") + "/" + url.ShortCode,
		"original_url": url.OriginalURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "code")

	url, err := h.repo.GetByShortCode(r.Context(), shortCode)
	if err != nil {
		http.Error(w, "url not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusMovedPermanently)
}

func generateCode(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, n)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}