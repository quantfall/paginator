package paginator

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Pagination[T any] struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Sort   string `json:"sort,omitempty"`
	Total  int64  `json:"total"`
	Rows   []T    `json:"rows"`
	Offset int    `json:"-"`
}

func New[T any](c *gin.Context) Pagination[T] {
	p := Pagination[T]{}
	p.Page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	p.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10"))
	p.Offset = (p.Page - 1) * p.Limit
	sort := c.QueryArray("sort")

	if len(sort) > 0 {
		srt := make([]string, 0)
		for _, s := range sort {
			srt = append(srt, strings.Replace(s, ",", " ", 1))
		}
		p.Sort = strings.Join(srt, ",")
	}
	return p
}
