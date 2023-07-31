package main

import (
	"encoding/json"
)

type RegistrationUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func registration(info string) error {
	var RegistrationUser RegistrationUser
	if err := json.Unmarshal([]byte(info), &RegistrationUser); err != nil {
		return err
	}

	hashPassword, err1 := HelperHashPass(RegistrationUser.Password)

	if err1 != nil {
		return err1
	}
	db, err2 := ConnectToPostgreSQL()

	if err2 != nil {
		return err2
	}

	err3 := RegistrationUserDB(db, RegistrationUser.Login, hashPassword, RegistrationUser.Name)

	if err3 != nil {
		return err3
	}
	return nil
}
