package delivery

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"example.local/go-pilot/internal/models"
	"example.local/go-pilot/internal/user"
)

// UserHandler represents http handler for user
type UserHandler struct {
	UserUsecase user.Usecase
}

// NewUserHandler will initialize /api/users endpoint
func NewUserHandler(us user.Usecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	http.HandleFunc("/api/users", handler.HandleUsers)
}

// HandleUsers will handle request based on request method
func (u *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		u.FetchUsers(w, r)
	case http.MethodPost:
		u.CreateUser(w, r)
	case http.MethodPut:
		u.UpdateUser(w, r)
	case http.MethodDelete:
		u.DeleteUser(w, r)
	}
}

func parseBody(body io.ReadCloser) (*models.User, error) {
	user := models.User{}

	jsn, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsn, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FetchUsers will fetch users based on given params
func (u *UserHandler) FetchUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.UserUsecase.Fetch()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	out, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// CreateUser will create new user
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := parseBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.UserUsecase.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UpdateUser will update user fields
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(r.URL.Query().Get("uid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := parseBody(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.UserUsecase.Update(uid, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteUser will remove user
func (u *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uid, err := strconv.Atoi(r.URL.Query().Get("uid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.UserUsecase.Delete(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
