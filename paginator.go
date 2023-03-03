package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Sort   string `json:"sort,omitempty"`
	Offset int    `json:"-"`
}

type Page[T any] struct {
	Total int64 `json:"total"`
	Rows  []T   `json:"rows"`
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

func (p *Pagination) pagingScope(db *gorm.DB) *gorm.DB {
	return db.Limit(p.Limit).Offset(p.Offset).Order(p.Sort)
}
