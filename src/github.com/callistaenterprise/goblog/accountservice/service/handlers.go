package service

import (
	"github.com/callistaenterprise/goblog/accountservice/dbclient"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"github.com/callistaenterprise/goblog/accountservice/model"
	"io/ioutil"
)

var DBClient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Read the 'accountId' path parameter from the mux map
	var accountId = mux.Vars(r)["accountId"]
	// Read the account struct BoltDB
	account, err := DBClient.QueryAccount(accountId)
	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Read the 'accountId' path parameter from the mux map
	var accountId = mux.Vars(r)["accountId"]
	var accountName = mux.Vars(r)["accountName"]
	// Read the account struct BoltDB
	account, err := DBClient.CreateAccount(accountId, accountName)
	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func JsonCreateAccount(w http.ResponseWriter, r *http.Request) {

	var account model.Account

	body, err := ioutil.ReadAll(r.Body)
	if err !=nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	if err := json.Unmarshal(body, &account); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	my_account, err := DBClient.CreateAccount(account.Id, account.Name)
	// If err, return a 404
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If found, marshal into JSON, write headers and content
	data, _ := json.Marshal(my_account)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
