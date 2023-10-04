package address

import (
	"encoding/json"
	"fmt"
	"jztpracticing/database"
	"jztpracticing/helpers"
	"net/http"
)

var UID string

func Add_Adress(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	house := r.FormValue("house")
	locality := r.FormValue("locality")
	city := r.FormValue("city")
	district := r.FormValue("district")
	state := r.FormValue("state")
	pin := r.FormValue("pin")

	fmt.Println(house)
	fmt.Println(district)

	var errr error

	UID, errr = helpers.Generate_Uid()

	if errr != nil {
		panic("couldnt create unique id")
	}

	var responce = make(map[string]string)

	address := database.Address{
		Addressoff: UID,
		House:      house,
		Locality:   locality,
		City:       city,
		District:   district,
		State:      state,
		Pin:        pin,
	}

	database.DB.Create(&address)

	responce["message"] = "added address successfully"

	jsondta, err := json.Marshal(responce)

	fmt.Println("addressof")
	fmt.Println(UID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsondta)

}
