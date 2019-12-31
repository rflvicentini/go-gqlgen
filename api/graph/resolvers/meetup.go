package resolvers

import (
	"context"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	"github.com/rflvicentini/go-gqlgen/api/graph/models"
)

type meetupResolver struct{ *Resolver}

func (r *Resolver) Meetup() generated.MeetupResolver {
	return &meetupResolver{r}
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	repoUser, err := r.UserRepo.GetById(obj.UserID)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID: repoUser.ID,
		Username: repoUser.Username,
		Email: repoUser.Email,
	}

	return &user, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	panic("implement me")
}
