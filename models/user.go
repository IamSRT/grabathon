package models

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Bundl Technologies Private Ltd.
 */

type User struct {
	PhoneNumber string
	AutoPay     bool
	Name        string
	City        string
	Email       string
	VouchCount  int
}

func CreateUser(user User) (User, error) {
	db.Create(&user)
	return user, nil
}

func GetUser(id string)(User, error){
	var user User
	db.Find(&user, User{PhoneNumber:id})
	return user, nil
}


func GetAllUsers()([]User, error){
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user User)(User, error){
	db.Save(&user)
	return user, nil
}

func DeleteUser(user User) error {
	db.Delete(&user)
	return nil
}