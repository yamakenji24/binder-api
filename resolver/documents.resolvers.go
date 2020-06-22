package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/url"
	"strconv"
	"time"

	"github.com/yamakenji24/binder-api/graph/generated"
	"github.com/yamakenji24/binder-api/graph/model"
	"github.com/yamakenji24/binder-api/minio"
	"github.com/yamakenji24/binder-api/repository"
)

func (r *mutationResolver) CreateDocument(ctx context.Context, input model.DocumentInput) (*model.GraphDocument, error) {
	userID, _ := ctx.Value("userId").(int)
	fp, err := decodeFile(strconv.Itoa(userID), input.Title, input.File)
	if err != nil {
		return &model.GraphDocument{}, err
	}

	doc, err := repository.CreateNewDocument(userID, input.Title, input.Description, fp)

	return &model.GraphDocument{
		ID:          strconv.FormatUint(uint64(doc.ID), 10),
		Title:       doc.Title,
		Description: doc.Description,
		File:        input.File,
	}, nil
}

func (r *queryResolver) Docs(ctx context.Context) ([]*model.GraphDocument, error) {
	docs, err := repository.GetAllDocument()
	if err != nil {
		return []*model.GraphDocument{}, err
	}

	graphDocs := make([]*model.GraphDocument, len(docs))
	for i, doc := range docs {
		docID := strconv.FormatUint(uint64(doc.ID), 10)
		file := encodeFile(doc.FilePath, docID)

		graphDocs[i] = &model.GraphDocument{
			ID:          docID,
			Title:       doc.Title,
			Description: doc.Description,
			File:        file.String(),
		}
	}

	return graphDocs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func encodeFile(file_path string, docID string) *url.URL {
	geneURL, _ := minio.GenerateURL(file_path)

	return geneURL
}

func decodeFile(userid string, title string, inputfile string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(inputfile)
	if err != nil {
		return "", err
	}

	fp := userid + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
	fr := bytes.NewReader(data)

	if err := minio.MinioUploader(fp, fr); err != nil {
		return "", err
	}

	return fp, nil
}
