package apikeyrepository

import (
	"database/sql"
)

func GetApiKeyById(db *sql.DB, id string) (string, error) {

	var apiKey string

	err := db.QueryRow("SELECT szApiKeyId FROM SYS_ApiKey WHERE szApiKeyId = ?", id).Scan(&apiKey)

	if err != nil {
		return "", err
	}

	return apiKey, nil
}
