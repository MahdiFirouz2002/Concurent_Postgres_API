package database

import (
	"encoding/json"
	"fmt"
	"users/common"
)

func InsertUser(user common.User) error {
	if err := checkConnection(); err != nil {
		panic("something went wrong with connection!")
	}

	insertQuery := `INSERT INTO users(id, name, email, phone_number) VALUES($1, $2, $3, $4)`
	_, err := database.Exec(insertQuery, user.Id, user.Name, user.Email, user.Phone_number)
	if err != nil {
		return err
	}

	userAddresses := user.Addresses
	for _, address := range userAddresses {
		insertAddressQuery := `INSERT INTO addresses(userId, street, city, state, zip_code, country) VALUES($1, $2, $3, $4, $5, $6)`
		_, err := database.Exec(insertAddressQuery, user.Id, address.Street, address.City, address.State, address.Zip_code, address.Country)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetUserByID(id string) (common.User, error) {
	if err := checkConnection(); err != nil {
		panic("something went wrong with connection!")
	}

	query := `
		SELECT u.id, u.name, u.email, u.phone_number,
			json_agg(
				json_build_object(
					'street', a.street,
					'city', a.city,
					'state', a.state,
					'zip_code', a.zip_code,
					'country', a.country
				)
			) AS addresses
		FROM users AS u
		LEFT JOIN addresses AS a ON a.userid = u.id
		WHERE u.id = $1
		GROUP BY u.id, u.name, u.email, u.phone_number;
	`

	rows, err := database.Query(query, id)
	if err != nil {
		return common.User{}, err
	}

	defer rows.Close()

	var user common.User
	if rows.Next() {
		var name, email, phone string
		var addressesJSON string

		err = rows.Scan(&user.Id, &name, &email, &phone, &addressesJSON)
		if err != nil {
			return common.User{}, err
		}

		user.Name = name
		user.Email = email
		user.Phone_number = phone

		var addresses []common.Addresses
		json.Unmarshal([]byte(addressesJSON), &addresses)
		user.Addresses = addresses
	} else {
		return common.User{}, fmt.Errorf("no user maches")
	}

	if err := rows.Err(); err != nil {
		return common.User{}, err
	}

	return user, nil
}
