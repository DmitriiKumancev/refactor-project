package api

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DmitriiKumancev/refactor-project/models"
	"github.com/DmitriiKumancev/refactor-project/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var (
	ErrUserNotFound = errors.New("user_not_found")
)

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, s.List)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	request := models.CreateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
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
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]interface{}{
		"user_id": id,
	})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")
	log.Printf("Received ID from URL: %s\n", id)

	user, ok := s.List[id]
	if !ok {
		_ = render.Render(w, r, models.ErrInvalidRequest(ErrUserNotFound))
		return
	}

	render.JSON(w, r, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	request := models.UpdateUserRequest{}
	if err := render.Bind(r, &request); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")

	user, ok := s.List[id]
	if !ok {
		_ = render.Render(w, r, models.ErrInvalidRequest(ErrUserNotFound))
		return
	}

	user.DisplayName = request.DisplayName

	if request.Email != "" {
		user.Email = request.Email
	}

	s.List[id] = user

	err = storage.SaveStore(s)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	f, err := os.ReadFile(storage.Store)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	s := models.UserStore{}
	if err := json.Unmarshal(f, &s); err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	id := chi.URLParam(r, "id")
	log.Printf("Received ID from URL: %s\n", id)


	_, ok := s.List[id]
	if !ok {
		_ = render.Render(w, r, models.ErrInvalidRequest(ErrUserNotFound))
		return
	}

	delete(s.List, id)

	err = storage.SaveStore(s)
	if err != nil {
		_ = render.Render(w, r, models.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusNoContent)
}
