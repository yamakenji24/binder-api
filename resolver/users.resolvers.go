package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
	"github.com/yamakenji24/binder-api/repository"
)

func (r *queryResolver) User(ctx context.Context, username string) (*model.GraphUser, error) {
	userid := ctx.Value("userId")
	fmt.Println(userid)
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
