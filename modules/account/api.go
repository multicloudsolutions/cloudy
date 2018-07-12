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
	var ac Account
	accountName := r.URL.Query().Get("name")
	(&ac).GetAccount(accountName)
	res, _ := json.Marshal(ac)
	fmt.Fprintf(w, string(res))
}

// Update account details
//func Update(w http.ResponseWriter, r *http.Request) {
//	accountDecoder := json.NewDecoder(r.Body)
//	var accData Account
//	err := accountDecoder.Decode(&accData)
//	if err != nil {
//		panic(err)
//	}
//	if err != nil {
//		log.Fatalln("error:", err)
//	}
//	accData.UpdateAccount()
//	fmt.Fprintf(w, "Account updated successfully")
//}
//

//Remove account
func Remove(w http.ResponseWriter, r *http.Request) {
	var ac Account
	accountName := r.URL.Query().Get("name")
	status := (&ac).RemoveAccount(accountName)
	fmt.Fprintf(w, status)
}
