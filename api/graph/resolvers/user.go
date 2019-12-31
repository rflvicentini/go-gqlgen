package resolvers

import (
	"context"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	"github.com/rflvicentini/go-gqlgen/api/graph/models"
)

type userResolver struct{ *Resolver}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	repoUsers := r.UserRepo.FetchAll()
	
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

	for _, m := range meetups {
		if m.UserID == obj.ID {
			userMeetups = append(userMeetups, m)
		}
	}

	return userMeetups, nil
}
