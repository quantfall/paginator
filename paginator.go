package paginator

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit  int         `json:"limit,omitempty"`
	Page   int         `json:"page,omitempty"`
	Sort   string      `json:"sort,omitempty"`
	Total  int64       `json:"total"`
	Rows   interface{} `json:"rows"`
	Offset int         `json:"-"`
}

func New(c *gin.Context) Pagination {
	p := Pagination{}
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
