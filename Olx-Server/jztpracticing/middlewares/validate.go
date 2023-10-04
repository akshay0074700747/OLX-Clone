package middlewares

import (
	"encoding/json"
	"fmt"
	"jztpracticing/token"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var Store = sessions.NewCookieStore([]byte("this-is-my-secret-key"))

func Validation_middleware(w http.ResponseWriter, r *http.Request) {

	if err := godotenv.Load(); err != nil {
		panic("couldnt load env files...")
	}

	secret := os.Getenv("SECRET_KEY")

	// var inputdata map[string]string

	// buf := new(bytes.Buffer)
	// buf.ReadFrom(r.Body)
	// bodyStr := buf.String()

	// err := json.Unmarshal([]byte(bodyStr), &inputdata)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// r.Body = io.NopCloser(bytes.NewBufferString(bodyStr))

	// username := inputdata["username"]
	// password := inputdata["password"]
	// email := inputdata["email"]

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")

	// var user database.Users
	// query := "select * from users where username = ?"

	// database.Query(&user, query, database.DB, username)

	// if user.Username == "" {
	// 	http.Error(w, "user doesnt exist", http.StatusInternalServerError)
	// 	return
	// }

	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	jwt_token, err := token.Generatejwt(username, email, []byte(secret))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(jwt_token)

	// sending the jwt token as a  JSON response to the reacr frontend server

	response := map[string]string{"token": jwt_token}

	jsonResponse, errr := json.Marshal(response)

	if errr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "jwtToken",
	// 	Value: jwt_token,
	// 	Path:  "/",
	// 	// Secure:   true, // Use HTTPS
	// 	HttpOnly: true, // Prevent client-side JavaScript access
	// 	SameSite: http.SameSiteStrictMode,
	// 	Domain:   "localhost",
	// })

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)

}
