package resolvers

import (
	"context"
	"errors"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	"github.com/rflvicentini/go-gqlgen/api/graph/models"
)

var meetups = []*models.Meetup{
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

var users = []*models.User{
	{
		ID: "1",
		Username: "Bob",
		Email: "bob@gmail.com",
	},
	{
		ID: "2",
		Username: "Jon",
		Email: "jon@gmail.com",
	},
}

type Resolver struct{}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }


// ROOT

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// USER

type userResolver struct{ *Resolver}

func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
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

// MEETUP

type meetupResolver struct{ *Resolver}

func (r *Resolver) Meetup() generated.MeetupResolver {
	return &meetupResolver{r}
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("User ID not exists")
	}

	return user, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	panic("implement me")
}
