package api

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DmitriiKumancev/refactor-project/apperrors"
	"github.com/DmitriiKumancev/refactor-project/models"
	"github.com/DmitriiKumancev/refactor-project/pkg/logger"
	"github.com/DmitriiKumancev/refactor-project/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		logger.GetLogger().Error("Error reading user store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		logger.GetLogger().Error("Error unmarshalling user data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	render.JSON(w, r, s.List)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		logger.GetLogger().Error("Error reading user store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		logger.GetLogger().Error("Error unmarshalling user data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	request := models.CreateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		logger.GetLogger().Error("Invalid request data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	s.Increment++
	u := models.User{
		CreatedAt:   time.Now().String(),
		DisplayName: request.DisplayName,
		Email:       request.Email,
	}

	id := strconv.Itoa(s.Increment)
	s.List[id] = u

	err = storage.SaveStore(s)
	if err != nil {
		logger.GetLogger().Error("Error saving user to store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	logger.GetLogger().Info("User created", zap.String("user_id", id))
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		logger.GetLogger().Error("Error reading user store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		logger.GetLogger().Error("Error unmarshalling user data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	id := chi.URLParam(r, "id")
	log.Printf("Received ID from URL: %s\n", id)

	user, ok := s.List[id]
	if !ok {
		logger.GetLogger().Error("User not found", zap.String("user_id", id))
		render.Render(w, r, apperrors.ErrUserNotFound)
		return
	}

	render.JSON(w, r, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		logger.GetLogger().Error("Error reading user store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		logger.GetLogger().Error("Error unmarshalling user data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	request := models.UpdateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		logger.GetLogger().Error("Invalid request data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	id := chi.URLParam(r, "id")

	user, ok := s.List[id]
	if !ok {
		logger.GetLogger().Error("User not found", zap.String("user_id", id))
		render.Render(w, r, apperrors.ErrUserNotFound)
		return
	}

	user.DisplayName = request.DisplayName

	if request.Email != "" {
		user.Email = request.Email
	}

	s.List[id] = user

	err = storage.SaveStore(s)
	if err != nil {
		logger.GetLogger().Error("Error saving user to store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	logger.GetLogger().Info("User updated", zap.String("user_id", id))
	render.Status(r, http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		logger.GetLogger().Error("Error reading user store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		logger.GetLogger().Error("Error unmarshalling user data", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	id := chi.URLParam(r, "id")
	log.Printf("Received ID from URL: %s\n", id)

	_, ok := s.List[id]
	if !ok {
		logger.GetLogger().Error("User not found", zap.String("user_id", id))
		render.Render(w, r, apperrors.ErrUserNotFound)
		return
	}

	delete(s.List, id)

	err = storage.SaveStore(s)
	if err != nil {
		logger.GetLogger().Error("Error saving user store", zap.Error(err))
		render.Render(w, r, apperrors.NewInvalidRequestError(err))
		return
	}

	logger.GetLogger().Info("User deleted", zap.String("user_id", id))
	render.Status(r, http.StatusNoContent)
}
