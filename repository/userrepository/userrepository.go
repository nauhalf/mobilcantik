package userrepository

import (
	"database/sql"

	"github.com/nauhalf/mobilcantik/repository/model"
)

func GetUserByUsername(db *sql.DB, username string) (*model.User, error) {

	user := new(model.User)
	err := db.QueryRow("SELECT u.intUserId, u.szUsername, u.szFullname, u.szEmail, u.szPassword, u.isDefault FROM SYS_User u WHERE szUsername = ?", username).
		Scan(&user.UserId, &user.Username, &user.Fullname, &user.Email, &user.Password, &user.IsDefault)

	if err != nil {
		return nil, err
	}

	return user, nil
}
