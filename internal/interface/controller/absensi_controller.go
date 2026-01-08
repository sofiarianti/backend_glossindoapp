package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/internal/entity"
	"api/internal/usecase"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AbsensiController struct {
	absensiUsecase usecase.AbsensiUsecase
}

func NewAbsensiController(db *gorm.DB) *AbsensiController {
	return &AbsensiController{
		absensiUsecase: usecase.NewAbsensiUsecase(db),
	}
}

func (c *AbsensiController) GetAllAbsensis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	absensis, err := c.absensiUsecase.GetAllAbsensis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(absensis)
}

func (c *AbsensiController) GetAbsensiByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid absensi ID", http.StatusBadRequest)
		return
	}

	absensi, err := c.absensiUsecase.GetAbsensiByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(absensi)
}

func (c *AbsensiController) GetAbsensiByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["user_id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	absensis, err := c.absensiUsecase.GetAbsensiByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(absensis)
}

func (c *AbsensiController) CreateAbsensi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var absensi entity.Absensi
	if err := json.NewDecoder(r.Body).Decode(&absensi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.absensiUsecase.CreateAbsensi(&absensi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(absensi)
}

func (c *AbsensiController) UpdateAbsensi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid absensi ID", http.StatusBadRequest)
		return
	}

	var absensi entity.Absensi
	if err := json.NewDecoder(r.Body).Decode(&absensi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.absensiUsecase.UpdateAbsensi(uint(id), &absensi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(absensi)
}

func (c *AbsensiController) DeleteAbsensi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid absensi ID", http.StatusBadRequest)
		return
	}

	err = c.absensiUsecase.DeleteAbsensi(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Absensi deleted successfully"})
}

// CheckIn handler
func (c *AbsensiController) CheckIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var absensi entity.Absensi
	if err := json.NewDecoder(r.Body).Decode(&absensi); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.absensiUsecase.CheckIn(&absensi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(absensi)
}

// CheckOut handler
func (c *AbsensiController) CheckOut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		UserID uint `json:"user_id"`
		// Include other optional fields if needed for body
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Alamat string `json:"alamat"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a temporary entity to pass data
	checkoutData := &entity.Absensi{
		Latitude: req.Latitude,
		Longitude: req.Longitude,
		Alamat: req.Alamat,
	}

	err := c.absensiUsecase.CheckOut(req.UserID, checkoutData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Checkout successful"})
}
