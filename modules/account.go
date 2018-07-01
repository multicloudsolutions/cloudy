package ac

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
)

// AccountHandler is the db connection
type AccountHandler struct {
	Db *sql.DB
}

func (h *AccountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var accountType, accessKey, secretKey string
	var id int64
	var err error

	log.Print("Adding new account")

	accountType = r.URL.Query().Get("account_type")
	accessKey = r.URL.Query().Get("access_key")
	secretKey = r.URL.Query().Get("secret_key")

	stmt, err := h.Db.Prepare("INSERT test SET account_type=?,access_key=?,secret_key=?")
	checkErr(err)

	res, err := stmt.Exec(accountType, accessKey, secretKey)
	checkErr(err)
	id, err = res.LastInsertId()
	checkErr(err)

	fmt.Fprintf(w, "Account added successfully, id: %d", id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
