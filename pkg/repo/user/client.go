package repo_user

import (
	"errors"
	"net/http"
)

type UserRepo struct {
	HTTPClient *http.Client
}

var mock = []*User{
	{
		ID: "1",
		Username: "Bob",
		Email: "bob@gmail.com",
		Birthdate: "2000-01-15",
	},
	{
		ID: "2",
		Username: "Jon",
		Email: "jon@gmail.com",
		Birthdate: "1999-12-10",
	},
}

func (c UserRepo) FetchAll() []*User {
	return mock
}

func (c UserRepo) GetById(id string) (*User, error) {
	user := new(User)

	for _, u := range mock {
		if u.ID == id {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("User ID does not exists")
	}

	return user, nil
}