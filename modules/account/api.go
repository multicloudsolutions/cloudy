package account

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create Create entry for new account
func Create(w http.ResponseWriter, r *http.Request) {

	accountDecoder := json.NewDecoder(r.Body)
	var accData Account
	err := accountDecoder.Decode(&accData)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatalln("error:", err)
	}
	accData.CreateAccount()
	fmt.Fprintf(w, "Account added successfully")
}

// GetDetails of account by name
func GetDetails(w http.ResponseWriter, r *http.Request) {
	accountDetails := GetAccount(r.URL.Query().Get("name"))
	fmt.Fprintf(w, accountDetails)
}

// Update account details
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Account updated successfully")
}

//Remove account
func Remove(w http.ResponseWriter, r *http.Request) {
	status := RemoveAccount(r.URL.Query().Get("name"))
	fmt.Fprintf(w, status)
}
