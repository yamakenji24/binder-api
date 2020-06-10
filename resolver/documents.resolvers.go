package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
)

func (r *mutationResolver) CreateDocument(ctx context.Context, input model.DocumentInput) (*model.GraphDocument, error) {

	decData := decodeFile(1, input.Title, input.File)
	fmt.Println(decData)

	return &model.GraphDocument{
		ID:          "1",
		Title:       input.Title,
		Description: input.Description,
		File:        input.File,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func decodeFile(id int, title string, file string) []byte {

	data, _ := base64.StdEncoding.DecodeString(file)

	return data
}
