package test

import (
	"aoBlog/log"
	"aoBlog/models"
	"testing"
)

func TestArticleArchive(t *testing.T) {
	list, err := models.Archive{}.GetAll()
	if err != nil {
		log.Logger.Sugar().Error("error -> ", err)
	}
	log.Logger.Sugar().Debug(list)
}
