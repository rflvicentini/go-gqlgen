package resolvers

import (
	generated "github.com/rflvicentini/go-gqlgen/api/graph/generated"
	repo_user "github.com/rflvicentini/go-gqlgen/pkg/repo/user"
	repo_meetup "github.com/rflvicentini/go-gqlgen/pkg/repo/meetup"
)

type Resolver struct{
	UserRepo    *repo_user.UserRepo
	MeetupRepo    *repo_meetup.MeetupRepo
}

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
