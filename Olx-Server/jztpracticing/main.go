package main

import (
	home "jztpracticing/controllers/Home"
	"jztpracticing/controllers/address"
	"jztpracticing/controllers/sell"
	"jztpracticing/database"
	"jztpracticing/middlewares"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// func CORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

// 		/*
// 			When a web page makes an API request to a different domain the browser will perform a preflight request to determine if the server allows the actual request to be made. This preflight request serves as a safety check to prevent potentially harmful cross-origin requests a preflight request is a request which is made by the browser before actually making the request the preflight request is a HTTPOPTIONS request here it checks that if the req is a preflight req ie HTTPOPTIONS request then sent a status code of http.StatusOK.
// 		*/
// 		if r.Method == http.MethodOptions {
// 			fmt.Println("this called")
// 			w.WriteHeader(http.StatusOK)
// 			return
// 		}

// 		// Call the next handler
// 		fmt.Println("this called")
// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {

	if err := godotenv.Load(); err != nil {
		panic("couldnt load env files...")
	}

	db_addr := os.Getenv("DB_ADDR")
	secret := os.Getenv("SECRET_KEY")

	database.Connect_to(db_addr)

	mux := http.NewServeMux()

	// http.Handle("/",CORS(mux))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Replace with your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	database.Migrte_all(&database.Users{}, &database.Products{}, &database.Address{}, &database.Images{})

	mux.Handle("/signup", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			middlewares.Signup_Middleware(http.HandlerFunc(middlewares.Validation_middleware)).ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}))

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			middlewares.Home_Middleware(http.HandlerFunc(home.Gethome), []byte(secret)).ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}))

	mux.Handle("/sell/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			http.HandlerFunc(sell.Selling).ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}))

	mux.Handle("/address", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			http.HandlerFunc(address.Add_Adress).ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}))

	mux.Handle("/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			http.HandlerFunc(middlewares.LoginNow).ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

	}))

	http.ListenAndServe(":8080", handler)

}
