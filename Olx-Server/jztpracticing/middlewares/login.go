package middlewares

import (
	"encoding/json"
	"fmt"
	"jztpracticing/database"
	"jztpracticing/token"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoginNow(w http.ResponseWriter, r *http.Request) {

	if err := godotenv.Load(); err != nil {
		panic("couldnt load env files...")
	}

	secret := os.Getenv("SECRET_KEY")

	// var data map[string]string

	// decoder := json.NewDecoder(r.Body)

	// if err := decoder.Decode(&data); err != nil {
	// 	http.Error(w,err.Error(),http.StatusBadRequest)
	// 	return
	// }

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	query := "select * from users where username = ?"

	var user database.Users

	database.Query(&user, query, database.DB, username)

	if user.Username == "" {
		http.Error(w, "User doesn't exist", http.StatusBadRequest)
		return
	}

	email := user.Email

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		http.Error(w, "username or password doesnt match", http.StatusBadRequest)
		return
	}

	jwt_token, err := token.Generatejwt(username, email, []byte(secret))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(jwt_token)

	response := map[string]string{"token": jwt_token}

	jsonResponse, errr := json.Marshal(response)

	if errr != nil {
		http.Error(w, errr.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)

}
