package controller

import (
	"quoteapp/model"
	"quoteapp/view"
	"io"
	"net/http"
	"quoteapp/db"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"quoteapp/authentication"
)


type Users2 struct {
	users *model.Users

}

type SavePassKey struct {
	Token string `json`
	Email string `json:"email"`
}

type Passkey struct {
	Password string `json:"password"`
	Email string `json:"email"`
}

func userStructController(u *model.Users) *Users2 {
	return &Users2{users: u}
}

func NewAuthController(u *model.Users) *Users2 {
	return &Users2{users: u}
}


func (u *Users2) SigningIn(w http.ResponseWriter, r *http.Request) {
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}	
	dataUser := db.Users{}
	err = json.Unmarshal(body, &dataUser)
	user, err := u.users.FindByEmail(dataUser.Email)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	newData := Passkey{}
	newPassword :=  newData.Password
	err = comparePassword(user.Password, newPassword)


	validToken, err := authentication.GenerateJWT(user.Email)

	if err != nil {
		view.ErrorRespond(w, err)
		return
	}
	
	data := map[string]string{
		"token": "Bearer " + string(validToken),
		"email": user.Email,
		"valid_trough": "30 minutes",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})

	
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}



