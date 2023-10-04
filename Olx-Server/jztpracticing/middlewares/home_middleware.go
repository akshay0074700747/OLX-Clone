package middlewares

import (
	"jztpracticing/token"
	"net/http"
	"strings"
)

func Home_Middleware(next http.HandlerFunc, secret []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie := r.Header.Get("Authorization")
		var jtoken string

		if cookie != "" && strings.HasPrefix(cookie, "Bearer ") {
			jtoken = strings.TrimPrefix(cookie, "Bearer ")
		}

		val, err := token.ValidateToken(jtoken, secret)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//i cant append any value to the r.Body sice its an io.ReadCloser  interface which is read only
		//so i have to read the existing request body, create a new buffer, write the existing data and the additional data to the buffer, and then assign the buffer as the new request body.

		// existing, err := io.ReadAll(r.Body)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// buf := bytes.NewBuffer(existing)

		// //cannot convert val (variable of type map[string]string) to type []bytecompilerInvalidConversion so i am converting it

		// jsonData, err := json.Marshal(val)

		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// buf.Write([]byte(jsonData))

		// r.Body = io.NopCloser(buf)

		for key, value := range val {
			r.Header.Add(key, value)
		}

		next(w, r)

	}
}
