package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/view"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	users *model.Users
}

func NewUsersController(q *model.Users) *Users {
	return &Users{users: q}
}

func (q *Users) Store(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	users := db.Users{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	generated, err := GeneratehashPassword(users.Password)

	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	users.Password = generated
	err = q.users.Create(&users)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"message": "registered successfully, please login",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}


func (q *Users) FindAll(w http.ResponseWriter, r *http.Request) {

	// add code to hit find all model user here
	data, err := q.users.FindAll()
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// add code to hit find by id model user here
	data, err := q.users.FindByID(id)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}


func (q *Users) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	user := db.Users{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	err = q.users.Update(id, &user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"pesan": "berhasil mengupdate user dengan id : " + id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// add code to hit delete model user here
	err := q.users.Delete(id)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"pesan": "berhasil menghapus user dengan id : " + id,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
