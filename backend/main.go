package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cfg := LoadConfig()
	provider := NewOctoProvider(cfg)

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{cfg.FrontendURL},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
		})

		r.Post("/payments/create", func(w http.ResponseWriter, r *http.Request) {
			var req createPaymentRequest
			if err := decodeJSON(r, &req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
				return
			}

			resp, err := provider.CreatePayment(r.Context(), req)
			if err != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
				return
			}

			writeJSON(w, http.StatusOK, resp)
		})

		r.Post("/applications/submit", func(w http.ResponseWriter, r *http.Request) {
			var req applicationSubmitRequest
			if err := decodeJSON(r, &req); err != nil {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
				return
			}

			if !req.AgreedToTermsAndPrivacyPolicy {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "terms_not_accepted"})
				return
			}

			if strings.TrimSpace(req.FullName) == "" || strings.TrimSpace(req.Email) == "" {
				writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing_required_fields"})
				return
			}

			applicationID := uuid.NewString()
			phone := digitsOnly(req.PhoneCountryDial) + digitsOnly(req.PhoneNumber)

			paymentResp, err := provider.CreatePayment(r.Context(), createPaymentRequest{
				Amount:      cfg.ApplicationFee,
				Currency:    cfg.OctoCurrency,
				Description: "EIMUN 2026 application fee",
				FullName:    req.FullName,
				Email:       req.Email,
				Phone:       phone,
			})
			if err != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
				return
			}

			log.Printf("application submitted id=%s fullName=%s email=%s", applicationID, req.FullName, req.Email)

			writeJSON(w, http.StatusOK, applicationSubmitResponse{
				ApplicationID: applicationID,
				PaymentID:     paymentResp.PaymentID,
				RedirectURL:   paymentResp.RedirectURL,
				Status:        paymentResp.Status,
			})
		})

		r.Post("/payments/notify", func(w http.ResponseWriter, r *http.Request) {
			body, _ := readBodyString(r)
			log.Printf("octo notify: %s", body)
			writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
		})
	})

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("backend listening on %s", cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
