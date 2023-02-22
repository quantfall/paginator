package paginator

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type Pagination struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Offset int
}

func GenerateFromRequest(c *gin.Context) Pagination {
	p := Pagination{}
	query := c.Request.URL.Query()
	for k, v := range query {
		switch k {
		case "limit":
			p.Limit, _ = strconv.Atoi(v[0])
			if p.Limit <= 0 {
				p.Limit = 10
			}
		case "page":
			p.Page, _ = strconv.Atoi(v[0])
			if p.Page <= 0 {
				p.Page = 1
			}
		case "sort":
			sort := make([]string, 0)
			for _, s := range v {
				sort = append(sort, strings.Replace(s, ",", " ", 1))
			}
			p.Sort = strings.Join(sort, ",")
		}
	}
	p.Offset = (p.Page - 1) * p.Limit
	return p
}
