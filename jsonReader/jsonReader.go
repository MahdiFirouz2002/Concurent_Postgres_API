package jsonReader

import (
	"encoding/json"
	"os"
	"users/common"
)

func read_Json(path string) (outputData []byte, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func UnmarshalData_from_json() (outputUsers []common.User, err error) {
	data, err := read_Json("users_data.json")
	if err != nil {
		return nil, err
	}

	var users []common.User
	parsingErr := json.Unmarshal(data, &users)
	if parsingErr != nil {
		return nil, parsingErr
	}

	return users, nil
}
