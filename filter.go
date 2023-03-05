package paginator

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
)

type Operator string

const (
	Is         Operator = "Is"
	IsNot      Operator = "IsNot"
	Contain    Operator = "Contain"
	NotContain Operator = "NotContain"
	In         Operator = "In"
	NotIn      Operator = "NotIn"
)

type Filter struct {
	Field string
	Op    Operator
	Value interface{}
}

func NewF(c *gin.Context) []Filter {
	filters := make([]Filter, 0)
	query := c.Request.URL.Query()

	for k, v := range query {
		switch k {
		case "limit", "sort", "page":
			continue
		}
		filter := strings.Split(k, ".")
		filters = append(filters, Filter{
			Field: filter[0],
			Op:    resolveOperator(filter[1]),
			Value: v[0]})
	}
	return filters
}

func resolveOperator(op string) Operator {
	var operator Operator
	switch op {
	case "Is":
		operator = Is
	case "IsNot":
		operator = IsNot
	case "Contain":
		operator = Contain
	case "NotContain":
		operator = NotContain
	case "In":
		operator = In
	case "NotIn":
		operator = NotIn
	}
	return operator
}

func FilterScope(filters []Filter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, filter := range filters {
			switch filter.Op {
			case Is:
				db.Where(filter.Field+" = ?", filter.Value)
			case IsNot:
				db.Where(filter.Field+" <> ?", filter.Value)
			case Contain:
				db.Where(filter.Field+" like %?%", filter.Value)
			case NotContain:
				db.Where(filter.Field+" not like ?", filter.Value)
			case In:
				db.Where(filter.Field+" in (?)", filter.Value)
			case NotIn:
				db.Where(filter.Field+" not in (?)", filter.Value)
			}
		}
		return db
	}
}
