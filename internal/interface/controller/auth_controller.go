package controller

import (
	"encoding/json"
	"net/http"

	"api/internal/entity"
	"api/internal/usecase"
	"gorm.io/gorm"
)

type AuthController struct {
	userUsecase usecase.UserUsecase
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		userUsecase: usecase.NewUserUsecase(db),
	}
}

// Struct untuk parsing response dari Google Token Info
type GoogleTokenInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"` // Bisa boolean atau string "true" tergantung endpoint
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Error         string `json:"error"`
	ErrorDesc     string `json:"error_description"`
}

func (c *AuthController) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req struct {
		IDToken string `json:"id_token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.IDToken == "" {
		http.Error(w, "id_token is required", http.StatusBadRequest)
		return
	}

	// Verifikasi token ke Google
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + req.IDToken)
	if err != nil {
		http.Error(w, "Failed to verify token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Invalid ID Token", http.StatusUnauthorized)
		return
	}

	var tokenInfo GoogleTokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		http.Error(w, "Failed to parse token info", http.StatusInternalServerError)
		return
	}

	if tokenInfo.Error != "" {
		http.Error(w, "Invalid token: "+tokenInfo.ErrorDesc, http.StatusUnauthorized)
		return
	}

	// Cek apakah user sudah ada di database berdasarkan Email
	user, err := c.userUsecase.GetUserByEmail(tokenInfo.Email)
	if err == nil && user.ID != 0 {
		// User sudah ada, update info jika perlu (misal foto atau nama berubah)
		// Opsional: update GoogleID jika belum ada
		if user.GoogleID == "" {
			user.GoogleID = tokenInfo.Sub
			user.PhotoURL = tokenInfo.Picture
			user.Name = tokenInfo.Name
			c.userUsecase.UpdateUser(user.ID, &user)
		}
		
		json.NewEncoder(w).Encode(user)
		return
	}

	// User belum ada, buat user baru
	newUser := entity.User{
		GoogleID: tokenInfo.Sub,
		Email:    tokenInfo.Email,
		Name:     tokenInfo.Name,
		PhotoURL: tokenInfo.Picture,
	}

	if err := c.userUsecase.CreateUser(&newUser); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
