package sell

import (
	"encoding/json"
	"fmt"
	"io"
	"jztpracticing/controllers/address"
	"jztpracticing/database"
	"net/http"
	"strconv"
	"strings"
)

func Selling(w http.ResponseWriter, r *http.Request) {

	//in net/http since query parameters wont work that why i had to do it

	parts := strings.Split(r.URL.Path, "/")

	name := parts[2]

	//limiting the size of the uploaded files in the total size of files is more than 20 mb then returns an error msg
	//this is done to prevent increased memmory usage and DoS attacks

	err := r.ParseMultipartForm(20 << 20)

	fmt.Println("seeeeeeeeeeeeeeeeeeeeelllllllllllllllllllllllllllllllllll")
	fmt.Println(name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["images"]
	var imageData [][]byte

	for _, file := range files {

		filee, err := file.Open()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer filee.Close()

		data, err := io.ReadAll(filee)

		if err != nil {
			http.Error(w, "Unable to read file content", http.StatusInternalServerError)
			return
		}

		imageData = append(imageData, data)

	}

	product := r.FormValue("productname")
	productdesc := r.FormValue("productdesc")
	price, errr := strconv.Atoi(r.FormValue("price"))

	if errr != nil {
		panic("the price cannot be converted to integer")
	}

	fmt.Println("Nooo problems till nowww...")

	// uid, err := helpers.Generate_Uid()

	fmt.Println("productid")
	fmt.Println(address.UID)

	data := database.Products{
		Productid:   address.UID,
		Soldby:      name,
		Productname: product,
		Productdesc: productdesc,
		Price:       price,
	}

	database.DB.Create(&data)

	for _, v := range imageData {

		images := database.Images{
			Productid: address.UID,
			Image:     v,
		}
		database.DB.Create(&images)

	}

	res := struct {
		Message string `json:"message"`
	}{
		Message: "Uploaded successfully",
	}

	jsondta, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsondta)

}
