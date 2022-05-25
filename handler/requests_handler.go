package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"senior-sidecar-poc/model"
	"senior-sidecar-poc/util"
	"time"
)

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func HandleRequests() {
	http.HandleFunc("/getUsers", returnGetUsers)
	http.HandleFunc("/des", returnDeserialized)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnGetUsers(w http.ResponseWriter, r *http.Request) {
	var users = []model.User{
		{
			Name:        "Alison",
			Description: "Dev",
			Email:       "alison.budag@senior.com.br",
			Locale:      "pt-br",
			AuthType:    model.G7,
			Birthdate:   time.Date(1994, time.July, 29, 10, 00, 00, 00, time.Local),
			Photo:       util.GetUserPhoto(),
			Blocked:     false,
			Roles: []model.Role{
				{
					Name:        "admin",
					Description: "Administrador do Sistema",
				},
				{
					Name:        "Go-POC",
					Description: "Golang POC",
				},
			},
		},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func returnDeserialized(w http.ResponseWriter, r *http.Request) {
	c := httpClient()
	var users []model.User
	err := requestGetUsers(c, "http://localhost:10000/getUsers", &users)
	if err != nil {
		panic(err)
	}

	user := (users)[0]
	fmt.Println("Usuário deserializado: ", user)
}

func requestGetUsers(client *http.Client, url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
