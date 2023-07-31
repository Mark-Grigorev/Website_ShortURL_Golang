package main

func authorization(login string, pass string) error {

	db, err := ConnectToPostgreSQL()
	if err != nil {
		return err
	}

	pass, _ = HelperHashPass(pass)

	err1 := AuthorizationUserDB(db, login, pass)

	if err1 != nil {
		return err1
	}
	return nil
}
