package main

import (
	"net/http"
)

// StudentHandler untuk halaman student
func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

// RequestMethodGet adalah middleware untuk membatasi akses hanya pada metode GET
func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Menggunakan middleware RequestMethodGet untuk endpoint /student
	http.Handle("/student", RequestMethodGet(StudentHandler()))

	http.ListenAndServe("localhost:8080", nil)
}
