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

type CutiController struct {
	cutiUsecase usecase.CutiUsecase
}

func NewCutiController(db *gorm.DB) *CutiController {
	return &CutiController{
		cutiUsecase: usecase.NewCutiUsecase(db),
	}
}

func (c *CutiController) GetAllCutis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cutis, err := c.cutiUsecase.GetAllCutis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cutis)
}

func (c *CutiController) GetCutiByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid cuti ID", http.StatusBadRequest)
		return
	}

	cuti, err := c.cutiUsecase.GetCutiByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cuti)
}

func (c *CutiController) GetCutiByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id_user"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cutis, err := c.cutiUsecase.GetCutiByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cutis)
}

func (c *CutiController) CreateCuti(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cuti entity.Cuti
	if err := json.NewDecoder(r.Body).Decode(&cuti); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := c.cutiUsecase.CreateCuti(&cuti)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cuti)
}

func (c *CutiController) UpdateCuti(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid cuti ID", http.StatusBadRequest)
		return
	}

	var cuti entity.Cuti
	if err := json.NewDecoder(r.Body).Decode(&cuti); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.cutiUsecase.UpdateCuti(uint(id), &cuti)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cuti)
}

func (c *CutiController) DeleteCuti(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid cuti ID", http.StatusBadRequest)
		return
	}

	err = c.cutiUsecase.DeleteCuti(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
