package account

import (
	"fmt"
	"log"
	"net/http"
        "encoding/json"
)

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

func GetDetails(w http.ResponseWriter, r *http.Request) {
    account_details := GetAccount(r.URL.Query().Get("name"))
    fmt.Fprintf(w, account_details)
}

func Update(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Account updated successfully")
}

func Remove(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Account deleted successfully")
}
