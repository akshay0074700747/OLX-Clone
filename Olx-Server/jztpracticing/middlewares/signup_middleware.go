package middlewares

import (
	"fmt"
	"jztpracticing/database"
	"jztpracticing/helpers"
	"net/http"
)

func Signup_Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// var inputdata map[string]string

		// // i read the request body and stored it in a buffer because on using json.NewDecoder(r.Body).Decode(&inputdata) the entire data in the r.Body() is getting consumed completly by the inputdata
		// //so that later i can assign tihs buffered data back into the r.Body

		// buf := new(bytes.Buffer)
		// buf.ReadFrom(r.Body)
		// bodyStr := buf.String()

		// // Create a new reader from the buffer and decode the JSON data

		// err := json.Unmarshal([]byte(bodyStr), &inputdata)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// // Restoring the request body with the original data which was stored inj the buffer

		// r.Body = io.NopCloser(bytes.NewBufferString(bodyStr))

		err := r.ParseForm()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")

		var user database.Users
		query := "select * from users where username = ?"

		//i am querying like this to make the Search() function loosely coupled to any structs

		database.Query(&user, query, database.DB, username)

		if user.Username != "" {
			http.Error(w, "User already exists", http.StatusNotAcceptable)
			return
		}

		hashed := helpers.Hash_pass(password)

		user = database.Users{
			Username: username,
			Password: hashed,
			Email:    email,
			Mobile:   mobile,
		}

		record := database.DB.Create(&user)

		if record.Error != nil {
			http.Error(w, "there was a problem while inserting user to the database", http.StatusNotAcceptable)
			return
		}

		fmt.Println("user successfully created")

		next(w, r)
	}
}

// func Signing_up(next http.HandlerFunc) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {

// 		session, _ := Store.Get(r, "my-session")

// 		// Get a value from the session
// 		user, ok := session.Values["jwt"].(string)

// 		if !ok {
// 			http.Error(w, "User not found in session", http.StatusNotFound)
// 			return
// 		}

// 		fmt.Println("hereeeeeeeeeeeeeeeeee we goooo..........")
// 		fmt.Println(user)

// 		var inputdata map[string]string

// 		json.NewDecoder(r.Body).Decode(&inputdata)

// 		data := struct {
// 			Username string `json:"username"`
// 			Password string `json:"password"`
// 			Email    string `json:"email"`
// 			Mobile   string `json:"mobile"`
// 			Cookie   string `json:"cookie"`
// 		}{
// 			Username: inputdata["username"],
// 			Password: inputdata["password"],
// 			Email:    inputdata["email"],
// 			Mobile:   inputdata["mobile"],
// 			Cookie:   user,
// 		}

// 		jsondta, _ := json.Marshal(data)
// 		w.WriteHeader(http.StatusOK)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(jsondta)

// 		// Set the response headers
// 		w.Header().Set("Content-Type", "text/plain")

// 		// Set the status code in the response
// 		w.WriteHeader(http.StatusOK)

// 		// Write the success message as the response body
// 		message := "Success! Your request was processed successfully."
// 		fmt.Fprint(w, message)

// 		http.Redirect(w, r, "/", http.StatusFound)

// 		next(w, r)

// 	}
// }
