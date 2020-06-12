package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/base64"
	"os"
	"strconv"
	"time"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
	"github.com/yamakenji24/binder-api/repository"
)

func (r *mutationResolver) CreateDocument(ctx context.Context, input model.DocumentInput) (*model.GraphDocument, error) {
	userID, _ := ctx.Value("userId").(int)
	filepath, err := decodeFile(userID, input.Title, input.File)
	if err != nil {
		return &model.GraphDocument{}, err
	}

	// DB 登録とか？
	doc, err := repository.CreateNewDocument(userID, input.Title, input.Description, filepath)

	return &model.GraphDocument{
		ID:          strconv.FormatUint(uint64(doc.ID), 10),
		Title:       doc.Title,
		Description: doc.Description,
		File:        input.File,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func decodeFile(id int, title string, inputfile string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(inputfile)
	if err != nil {
		return "", err
	}
	dir, err := os.Getwd()
	fp := dir + "/document/" + strconv.Itoa(id) + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"

	file, err := os.Create(fp)
	defer file.Close()
	file.Write(data)
	if err != nil {
		return "", err
	}

	return fp, nil
}
