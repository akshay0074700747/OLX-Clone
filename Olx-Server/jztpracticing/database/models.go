package database

import "fmt"

//this interface is used to group all the tables in one interface so that i can write only one migrator function to all of the tables

type Group_tables interface {
	Its_a_table()
}

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

func (user *Users) Its_a_table() {
	fmt.Println("Yeah... this is a table")
}

type Products struct {
	Productid   string `json:"productid"`
	Soldby      string `json:"soldby"`
	Productname string `json:"productname"`
	Productdesc string `json:"productdesc"`
	Price       int    `json:"price"`
}

func (product *Products) Its_a_table() {
	fmt.Println("Yeah... this is a table")
}

type Address struct {
	Addressoff string `json:"addressoff"`
	House      string `json:"house"`
	Locality   string `json:"locality"`
	City       string `json:"city"`
	District   string `json:"district"`
	State      string `json:"state"`
	Pin        string `json:"pin"`
}

func (addr *Address) Its_a_table() {
	fmt.Println("Yeah... this is a table")
}

type Images struct {
	Productid string `json:"productid"`
	Image     []byte `json:"image"`
}

func (imgs *Images) Its_a_table() {
	fmt.Println("Yeah... this is a table")
}
