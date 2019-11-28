package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type User struct {
}

func Create(user User) (User, error) {
	db.Create(user)
	return user, nil
}