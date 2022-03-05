package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	pageCtx     = "page"
	categoryCtx = "category"
	idCtx       = "id"
)

// this is meant to be constant! Please don't mutate it!
var availableCategory = []string{"1", "2", "6", "12", "15", "25", "30", "23"}

func getPage(ctx *gin.Context) (string, error) {
	page := ctx.Query(pageCtx)

	_, ok := strconv.ParseUint(page, 10, 32)
	if ok != nil {
		return "", errors.New("page is of invalid type")
	}

	return page, nil
}

func getCategory(ctx *gin.Context) (string, error) {
	category := ctx.Query(categoryCtx)
	for _, n := range availableCategory {
		if category == n {
			return category, nil
		}
	}

	return "", errors.New("category not found")
}

func getId(ctx *gin.Context) (string, error) {
	id := ctx.Query(idCtx)

	intId, ok := strconv.Atoi(id)
	if ok != nil || intId < 0 {
		return "", errors.New("id is of invalid type")
	}

	return id, nil
}
