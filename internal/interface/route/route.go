package route

import (
    "github.com/gorilla/mux"
    "api/internal/interface/controller"
    "gorm.io/gorm"
)

func SetupRoutesUser(router *mux.Router, db *gorm.DB) {
    userController := controller.NewUserController(db)

    // User routes
    router.HandleFunc("/api/user", userController.GetAllUsers).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/user/insert", userController.CreateUser).Methods("POST","OPTIONS")
    router.HandleFunc("/api/user/{id_user}", userController.GetUserByID).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/user/update/{id_user}", userController.UpdateUser).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/user/delete/{id_user}", userController.DeleteUser).Methods("DELETE", "OPTIONS")
} 

func SetupRoutesAbsensi(router *mux.Router, db *gorm.DB) {
    absensiController := controller.NewAbsensiController(db)

    // Absensi routes
    router.HandleFunc("/api/absensi", absensiController.GetAllAbsensis).Methods("GET", "OPTIONS")
    // Generic create (optional, maybe keep for admin?)
    router.HandleFunc("/api/absensi", absensiController.CreateAbsensi).Methods("POST", "OPTIONS")
    
    // Check-in and Check-out routes
    router.HandleFunc("/api/absensi/checkin", absensiController.CheckIn).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/absensi/checkout", absensiController.CheckOut).Methods("POST", "OPTIONS")

    router.HandleFunc("/api/absensi/{id}", absensiController.GetAbsensiByID).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/absensi/user/{user_id}", absensiController.GetAbsensiByUserID).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/absensi/update/{id}", absensiController.UpdateAbsensi).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/absensi/delete/{id}", absensiController.DeleteAbsensi).Methods("DELETE", "OPTIONS")
}

func SetupRoutesCuti(router *mux.Router, db *gorm.DB) {
    cutiController := controller.NewCutiController(db)

    // Cuti routes
    router.HandleFunc("/api/cuti", cutiController.GetAllCutis).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/cuti", cutiController.CreateCuti).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/cuti/{id}", cutiController.GetCutiByID).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/cuti/user/{id_user}", cutiController.GetCutiByUserID).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/cuti/update/{id}", cutiController.UpdateCuti).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/cuti/delete/{id}", cutiController.DeleteCuti).Methods("DELETE", "OPTIONS")
} 

func SetupRoutesAuth(router *mux.Router, db *gorm.DB) {
    authController := controller.NewAuthController(db)

    // Auth routes
	router.HandleFunc("/api/auth/google", authController.GoogleAuth).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/register", authController.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/auth/login", authController.Login).Methods("POST", "OPTIONS")
}
