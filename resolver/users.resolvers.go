package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	// mock data
	if input.Username == "yamakenji24" {
		return "jwt token will be returned", nil
	}
	return "", fmt.Errorf("not yamakenji")
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return &model.User{
		ID: "1",
		Username: "yamakenji24",
		Password: "hogehoge",
		Email: "yamakenji24@example.com",
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
