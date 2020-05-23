package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

type SqlExtension struct {
	OrderBy  string
	AscOrder bool
	Page     int32
	Size     int32
}

//  解析请求中排序分页信息，暂时只支持单个排序
func getPagedList(query *gorm.DB, se SqlExtension) *gorm.DB {
	//fmt.Printf("se val:%#v", se)
	if &se == nil {
		return query
	}
	if se.OrderBy != "" {
		if se.AscOrder {
			se.OrderBy += " asc"
		} else {
			se.OrderBy += " desc"
		}
		query = query.Order(se.OrderBy)
	}

	if se.Page != 0 {
		query = query.Offset((se.Page - 1) * se.Size)
	}
	if se.Size != 0 {
		if se.Size > 50 {
			se.Size = 50
		}
		query = query.Limit(se.Size)
	} else {
		query = query.Limit(10)
	}
	return query
}

// sql build where
// param where: 查询条件写入map结构体，key:dbColumn[ dbOperator], val: queryValue.
// 例：map[string]interface{
// 		"name": "ceshi1",
// 		"id in": []int{20, 19, 18},
// }
// result: where name = 'ceshi1' and id in (20,19,18)
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}

		if whereSQL != "" {
			whereSQL += " AND "
		}
		strings.Join(ks, ",")
		switch len(ks) {
		case 1:
			//fmt.Println(reflect.TypeOf(v))
			switch v := v.(type) {
			case NullType:
				if v == IsNotNull {
					whereSQL += fmt.Sprint(k, " IS NOT NULL")
				} else {
					whereSQL += fmt.Sprint(k, " IS NULL")
				}
			default:
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
			}
			break
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				whereSQL += fmt.Sprint(k, "=?")
				vals = append(vals, v)
				break
			case ">":
				whereSQL += fmt.Sprint(k, ">?")
				vals = append(vals, v)
				break
			case ">=":
				whereSQL += fmt.Sprint(k, ">=?")
				vals = append(vals, v)
				break
			case "<":
				whereSQL += fmt.Sprint(k, "<?")
				vals = append(vals, v)
				break
			case "<=":
				whereSQL += fmt.Sprint(k, "<=?")
				vals = append(vals, v)
				break
			case "!=":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
				break
			case "<>":
				whereSQL += fmt.Sprint(k, "!=?")
				vals = append(vals, v)
				break
			case "in":
				whereSQL += fmt.Sprint(k, " in (?)")
				vals = append(vals, v)
				break
			case "like":
				whereSQL += fmt.Sprint(k, " like ?")
				vals = append(vals, fmt.Sprintf("%v", v)+"%")
			}
			break
		}
	}
	return
}
