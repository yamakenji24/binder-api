package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
	"github.com/yamakenji24/binder-api/repository"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	user, err := repository.GetUserByName(input.Username)
	if err != nil {
		return "", err
	}
	if !compareHashedPassword(user.Password, input.Password) {
		return "", fmt.Errorf("Invalid username and password combination!")
	}
	// TODO: create jwt token
	return "jwt token will be returned", nil
}

func (r *queryResolver) User(ctx context.Context, username string) (*model.GraphUser, error) {
	user, err := repository.GetUserByName(username)
	if err != nil {
		return &model.GraphUser{}, err
	}
	return &model.GraphUser{
		ID:       strconv.FormatUint(uint64(user.ID), 10),
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

func compareHashedPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}
