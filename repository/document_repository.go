package repository

import "github.com/yamakenji24/binder-api/models"

func CreateNewDocument(userID int, title string, description string, filepath string) (*models.Document, error) {
	db := NewSQLHandler()
	doc := models.Document{
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

func GetAllDocument() (docs []*models.Document, err error) {
	db := NewSQLHandler()
	if err := db.Find(&docs).Error; err != nil {
		return nil, err
	}
	return docs, nil
}
