package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	Sort       string `json:"sort"`
	TotalPages int    `json:"totalpages"`
	TotalItems int64  `json:"totalitems"`
	Filter     string `json:"filter"`
}

func GeneratePaginationFromRequest(c *gin.Context) Pagination {
	// Initializing default
	limit := 10
	page := 1
	sort := ""
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break

		}
	}
	return Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
