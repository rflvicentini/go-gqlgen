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
	return users, nil
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
