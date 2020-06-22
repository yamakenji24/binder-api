package repository

import "github.com/yamakenji24/binder-api/model"

func CreateNewDocument(userID int, title string, description string, filepath string) (*model.Document, error) {
	db := NewSQLHandler()
	doc := model.Document{
		UserID:      userID,
		Title:       title,
		Description: description,
		FilePath:    filepath,
	}
	if err := db.Create(&doc).Error; err != nil {
		return nil, err
	}
	return &doc, nil
}

func GetAllDocument() (docs []*model.Document, err error) {
	db := NewSQLHandler()
	if err := db.Find(&docs).Error; err != nil {
		return nil, err
	}
	return docs, nil
}
