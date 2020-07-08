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
	"github.com/yamakenji24/binder-api/models"
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

func (r *queryResolver) Docs(ctx context.Context, page *model.PaginationInput) (*model.DocumentConnection, error) {
	docs, err := repository.GetAllDocument()
	if err != nil {
		return &model.DocumentConnection{}, err
	}

	total := len(docs)
	start := 0
	docLength := total

	if page.First != nil {
		docLength = *page.First
	}
	if page.Offset != nil {
		start = *page.Offset
	}

	pageDocs := docs[start : start+docLength]
	graphDocs, endCursor := handleCursor(pageDocs)

	startCursor := encodeToBase64(strconv.FormatUint(uint64(pageDocs[0].ID), 10))
	hasNextPage := (start + *page.First) < total

	return &model.DocumentConnection{
		PageInfo: &model.PageInfo{
			StartCursor: startCursor,
			EndCursor:   endCursor,
			HasNextPage: hasNextPage,
		},
		Edges: graphDocs,
		Total: strconv.Itoa(total),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func handleCursor(pageDocs []*models.Document) ([]*model.DocumentEdge, string) {
	var endCursor string
	graphDocs := make([]*model.DocumentEdge, len(pageDocs))
	for i, doc := range pageDocs {
		docID := strconv.FormatUint(uint64(doc.ID), 10)
		file := getFileURL(doc.FilePath)
		endCursor = encodeToBase64(docID)

		graphDocs[i] = &model.DocumentEdge{
			Cursor: endCursor,
			Node: &model.GraphDocument{
				ID:          docID,
				Title:       doc.Title,
				Description: doc.Description,
				File:        file.String(),
			},
		}
	}
	return graphDocs, endCursor
}
func encodeToBase64(id string) string {
	return base64.StdEncoding.EncodeToString([]byte(id))
}
func getFileURL(file_path string) *url.URL {
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
