package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	_ "strconv"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
	"github.com/yamakenji24/binder-api/repository"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	// mock data
	if input.Username == "yamakenji24" {
		return "jwt token will be returned", nil
	}
	return "", fmt.Errorf("not yamakenji")
}

func (r *queryResolver) User(ctx context.Context, username string) (*model.GraphUser, error) {
	user, _ := repository.GetUserByName(username)
	return &model.GraphUser{
		ID:       "1", //strconv.FormatUint(uint64(user.ID), 10),
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
