package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api/config"
	"api/internal/interface/middleware"
	"api/internal/interface/route"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize DB
	db := config.InitDB()

	// Initialize router
	router := mux.NewRouter()

	// Add middleware
	router.Use(middleware.LoggingMiddleware)

	// Static files (commented out as directories might not exist yet)
	// router.PathPrefix("/posters/").Handler(http.StripPrefix("/posters/", http.FileServer(http.Dir("./posters/"))))
	// router.PathPrefix("/trailers/").Handler(http.StripPrefix("/trailers/", http.FileServer(http.Dir("trailers"))))
	// router.PathPrefix("/qrcode/").Handler(http.StripPrefix("/qrcode/", http.FileServer(http.Dir("public/qrcode/"))))

	// Setup routes
	route.SetupRoutesUser(router, db)
	route.SetupRoutesAbsensi(router, db)
	route.SetupRoutesCuti(router, db)
	route.SetupRoutesAuth(router, db)
	// route.SetupRoutesVisitor(router, db)

	// Konfigurasi server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Server dijalankan dalam goroutine terpisah
	go func() {
		fmt.Printf("=========================================\n")
		fmt.Printf("Server berjalan pada port %s...\n", port)
		fmt.Printf("Endpoint yang tersedia:\n")
		fmt.Printf("GET: http://localhost:%s/api/cuti\n", port)
		fmt.Printf("POST: http://localhost:%s/api/cuti\n", port)
		fmt.Printf("POST: http://localhost:%s/api/auth/google\n", port)
		// fmt.Printf("GET: http://localhost:%s/api/\n", port) // Update with actual routes
		fmt.Printf("=========================================\n")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error menjalankan server: %s\n", err)
		}
	}()

	// Channel untuk menangkap signal shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	log.Println("Server berhasil shutdown")
}
