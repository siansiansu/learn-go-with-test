package server

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

func GetEmployee(url string) (*Employee, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var emp Employee

	err = json.NewDecoder(resp.Body).Decode(&emp)
	if err != nil {
		return nil, err
	}
	return &emp, nil
}
