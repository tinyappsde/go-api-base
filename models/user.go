package models

import (
	"errors"
	"tinyapps/api-base/handlers"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	Id				string			`json:"id"`
	InternalId		int				`json:"-"`
	Username		null.String		`json:"username"`
	Email			null.String		`json:"email"`
	FirstName		string			`json:"firstName"`
	LastName		null.String		`json:"lastName"`
	PasswordHash	null.String		`json:"-"`
}

func DummyUser() User {
	return User{
		Id: "abc123",
		Username: null.StringFrom("johndoe123"),
		Email: null.StringFrom("john.doe@example.com"),
		FirstName: "John",
	}
}

func GetUser(env *handlers.Env, userId string) (User, error) {
	var user User

	results, err := env.DB.Query("SELECT `id`, `public_id`, `username`, `email`, `first_name`, `last_name` FROM `users` WHERE `public_id` = ?", userId)
	if err != nil {
		return user, err
	}

	if (results.Next()) {
		err = results.Scan(
			&user.InternalId,
			&user.Id,
			&user.Username,
			&user.Email,
			&user.FirstName,
			&user.LastName,
		)

		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, errors.New("not found")
}

func GetUserByEmail(env *handlers.Env, email string) (User, error) {
	var user User

	results, err := env.DB.Query("SELECT `id`, `public_id`, `username`, `email`, `first_name`, `last_name`, `password_hash` FROM `users` WHERE `email` = ?", email)
	if err != nil {
		return user, err
	}

	if (results.Next()) {
		err = results.Scan(
			&user.InternalId,
			&user.Id,
			&user.Username,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.PasswordHash,
		)

		if err != nil {
			return user, err
		}

		return user, nil
	}

	return user, errors.New("not found")
}
