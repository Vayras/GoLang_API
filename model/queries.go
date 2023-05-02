package model

import (
	"errors"
	"fmt"
)

func GetAll() ([]Person, error) {
	var people []Person // Rename "Person" to "people" to follow naming conventions
	if err := DB.Find(&people).Error; err != nil {
		return nil, err
	}
	return people, nil
}

func CreatePerson(p Person) error {
	if err := DB.Create(&p).Error; err != nil {
		return err
	}
	return nil
}

func DeletePersonByID(id int) error {
	user := Person{Id: id}
	result := DB.Unscoped().Delete(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	fmt.Println("Deleted user", id)
	return nil

}
