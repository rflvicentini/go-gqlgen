package resolvers

import (
	"context"
	"errors"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	"github.com/rflvicentini/go-gqlgen/api/graph/models"
)

type userResolver struct{ *Resolver}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	repoUsers, err := r.UserRepo.FetchAll()
	if err != nil {
		return nil, errors.New("Unknown error when fetching data")
	}
	
	result := []*models.User{}
	for _, u := range repoUsers {
		resultObj := models.User{
			ID: u.ID,
			Username: u.Username,
			Email: u.Email,
		}
		result = append(result, &resultObj)
	}

	return result, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var userMeetups []*models.Meetup

	repoMeetups, err := r.MeetupRepo.FetchAll()
	if err != nil {
		return nil, errors.New("Unknown error when fetching data")
	}

	for _, m := range repoMeetups {
		if m.UserID == obj.ID {
			meetup := models.Meetup{
				ID: m.ID,
				Name: m.Name,
				Description: m.Description,
				UserID: m.UserID,
			}
			userMeetups = append(userMeetups, &meetup)
		}
	}

	return userMeetups, nil
}
