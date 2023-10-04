package home

import (
	"encoding/json"
	"fmt"
	"jztpracticing/database"
	"net/http"
)

func Gethome(w http.ResponseWriter, r *http.Request) {

	var allproducts []database.Products
	var allimages []database.Images
	var alladress []database.Address

	database.DB.Find(&allproducts)
	database.DB.Find(&allimages)
	database.DB.Find(&alladress)

	var results []struct {
		Products database.Products `json:"product"`
		Address  database.Address  `json:"address"`
		Images   []database.Images `json:"images"`
	}

	for i, product := range allproducts {
		element := struct {
			Products database.Products `json:"product"`
			Address  database.Address  `json:"address"`
			Images   []database.Images `json:"images"`
		}{}
		results = append(results, element)
		results[i].Products = product
		for _, address := range alladress {
			if product.Productid == address.Addressoff {
				results[i].Address = address
			}
		}
		for _, image := range allimages {
			if image.Productid == product.Productid {
				results[i].Images = append(results[i].Images, image)
			}
		}
	}

	fmt.Println(results[1].Products.Soldby)

	// responce := struct {
	// 	Username  string              `json:"username"`
	// 	Products  []database.Products `json:"product"`
	// 	Images    []database.Images   `json:"images"`
	// 	Addresses []database.Address  `json:"addresses"`
	// }{
	// 	Username:  username,
	// 	Products:  allproducts,
	// 	Images:    allimages,
	// 	Addresses: alladress,
	// }

	jsonData, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)

}
