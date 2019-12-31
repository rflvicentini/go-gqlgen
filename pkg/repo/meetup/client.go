package repo_meetup

import (
	"errors"
	"net/http"
)

type MeetupRepo struct {
	HTTPClient *http.Client
}

var mock = []*Meetup{
	{
		ID: "1",
		Name: "Meetup: Go",
		Description: "Golang",
		UserID: "1",
	},
	{
		ID: "2",
		Name: "Meetup: GQLGen",
		Description: "GQLGen",
		UserID: "2",
	},
}

func (c MeetupRepo) FetchAll() ([]*Meetup, error) {
	return mock, nil
}

func (c MeetupRepo) GetById(id string) (*Meetup, error) {
	meetup := new(Meetup)

	for _, m := range mock {
		if m.ID == id {
			meetup = m
			break
		}
	}

	if meetup == nil {
		return nil, errors.New("User ID does not exists")
	}

	return meetup, nil
}