package resolvers

import (
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	"github.com/rflvicentini/go-gqlgen/api/graph/models"
	repo_user "github.com/rflvicentini/go-gqlgen/pkg/repo/user"
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

type Resolver struct{
	UserRepo    *repo_user.UserRepo
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
