package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/satyajitnayk/uptime_monitor/internal/auth"
	"github.com/satyajitnayk/uptime_monitor/internal/models"
	"github.com/satyajitnayk/uptime_monitor/repositories"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterResponse structure for JSON response
type RegisterResponse struct {
	Message      string `json:"message"`
	ErrorMessage string `json:"error,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var request RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := RegisterResponse{ErrorMessage: "Invalid request payload"}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if request.Email == "" || request.Password == "" {
		response := RegisterResponse{ErrorMessage: "Email and password are required"}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		response := RegisterResponse{ErrorMessage: "Failed to hash password"}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	user := &models.User{
		Email:    request.Email,
		Password: hashedPassword,
	}

	userRepo := repositories.NewUserRepository(db)

	if err := userRepo.Create(user); err != nil {
		response := RegisterResponse{ErrorMessage: "Failed to create user"}
		sendJSONResponse(w, http.StatusInternalServerError, response)
		return
	}

	response := RegisterResponse{Message: "User registered successfully"}
	sendJSONResponse(w, http.StatusCreated, response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var request LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response := LoginResponse{ErrorMessage: "Invalid login payload"}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	if request.Email == "" || request.Password == "" {
		response := LoginResponse{ErrorMessage: "Email and password are required"}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	userRepo := repositories.NewUserRepository(db)

	user, err := userRepo.FindByEmail(request.Email)
	if err != nil {
		response := LoginResponse{ErrorMessage: "Failed to login"}
		sendJSONResponse(w, http.StatusBadRequest, response)
		return
	}

	// comapre password
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) == nil {
		// create jwt
		jwt, err := auth.GenerateJWT(request.Email)
		if err != nil {
			response := LoginResponse{ErrorMessage: "unable to generate token"}
			sendJSONResponse(w, http.StatusInternalServerError, response)
			return
		} else {
			response := LoginResponse{Token: jwt}
			sendJSONResponse(w, http.StatusOK, response)
			return
		}
	} else {
		response := LoginResponse{ErrorMessage: "Failed to Login"}
		sendJSONResponse(w, http.StatusUnauthorized, response)
		return
	}
}
