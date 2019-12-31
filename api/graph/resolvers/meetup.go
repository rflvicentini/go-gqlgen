package resolvers

import (
	"context"
	"errors"
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	"github.com/rflvicentini/go-gqlgen/api/graph/models"
)

type meetupResolver struct{ *Resolver}

func (r *Resolver) Meetup() generated.MeetupResolver {
	return &meetupResolver{r}
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	repoMeetups, err := r.MeetupRepo.FetchAll()
	if err != nil {
		return nil, errors.New("Unknown error when fetching data")
	}
	
	result := []*models.Meetup{}
	for _, m := range repoMeetups {
		resultObj := models.Meetup{
			ID: m.ID,
			Name: m.Name,
			Description: m.Description,
			UserID: m.UserID,
		}
		result = append(result, &resultObj)
	}

	return result, nil
}

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	repoUser, err := r.UserRepo.GetById(obj.UserID)
	if err != nil {
		return nil, errors.New("Unknown error when fetching user data")
	}
	if repoUser == nil {
		return nil, errors.New("User ID does not exists")
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
