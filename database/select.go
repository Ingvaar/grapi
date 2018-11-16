package sql

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"grapi/utils"
)

/*Select : Select entry from table
** {GET} /sql/{table}/
**
** {table} in the URL indicate which table to select in (mandatory)
**
**	     |	NAME	|	TYPE	|	DESCRIPTION		|	EXAMPLE
** @apiParam | field	| Array[string] | field to select		| field=id
** @apiParam | ijoin	| Array[string] | inner join tables		| ijoin=table2-id-id
**	-> ijoin={table_to_join}-{field_table_to_join}-{field_current_table}
** @apiParam | ojoin	| Array[string]	| outer join tables		| ojoin=table2-id-id
** @apiParam | order	| Array[string] | fields used to sort		| order=date
** @apiParam | desc	| Bool		| order by descending if true	| desc=true
** @apiParam | where	| Array[string] | select only matching param	| where=id=1
** @apiParam | offset	| Int		| offset (default: 0)		| offset=10
** @apiParam | limit	| Int		| limit (default: 20)		| limit=40
 */
func (db *SQL) Select(w http.ResponseWriter, r *http.Request) {
	tabName := mux.Vars(r)["table"]
	r.ParseForm()

	statement := prepareStatement(r.Form, tabName)
	rows, err := db.DB.Query(statement)
	defer rows.Close()
	colNames, errCol := rows.Columns()
	if err != nil {
		utils.SendError(w, err, http.StatusBadRequest)
	} else if errCol != nil {
		utils.SendError(w, err, http.StatusInternalServerError)
	} else {
		PrintOne(colNames, rows, w)
	}
}

func prepareStatement(form url.Values, tab string) string {
	var statement string

	statement = "SELECT " + getField(form["field"])
	statement += " FROM " + tab
	statement += join(form["ijoin"], tab, " INNER JOIN ")
	statement += join(form["ojoin"], tab, " OUTER JOIN ")
	statement += getOrder(form["order"], form["desc"])
	statement += getWhere(form["where"])
	statement += getLimits(form["offset"], form["limit"])
	return (statement)
}

func getLimits(offset []string, limit []string) string {
	limitStr := " LIMIT "
	if len(offset) > 0 {
		limitStr += checkNumber(offset, "0")
	} else {
		limitStr += "0"
	}
	limitStr += ","
	if len(limit) > 0 {
		limitStr += checkNumber(limit, "20")
	} else {
		limitStr += "20"
	}
	return (limitStr)
}

func getWhere(where []string) string {
	if len(where) > 0 {
		mult := false
		whereStr := " WHERE "
		for _, value := range where {
			if mult {
				whereStr += ", "
			}
			whereStr += value
			mult = true
		}
		return (whereStr)
	}
	return ("")
}

func getOrder(order []string, desc []string) string {
	if len(order) > 0 {
		mult := false
		orderStr := " ORDER BY "
		for _, value := range order {
			if mult {
				orderStr += ", "
			}
			orderStr += value
			mult = true
		}
		if len(desc) > 0 && desc[0] == "true" || desc[0] == "1" {
			orderStr += " DESC"
		}
		return (orderStr)
	}
	return ("")
}

func join(joins []string, tabName string, jType string) string {
	jStr := jType
	tempJGet := ""
	mult := false
	for _, join := range joins {
		tempJGet = getJoin(join, tabName)
		if mult && tempJGet != "" {
			jStr += jType
		}
		jStr += tempJGet
		mult = true
	}
	if len(joins) > 0 {
		return (jStr)
	}
	return ("")
}

func getJoin(join string, tab string) string {
	joinSplit := strings.Split(join, "-")

	if len(joinSplit) == 3 {
		return (joinSplit[0] + " ON " +
			joinSplit[0] + "." + joinSplit[1] +
			"=" + tab + "." + joinSplit[2])
	}
	return ("")
}

func getField(field []string) string {
	if len(field) > 0 {
		mult := false
		fields := ""
		for _, value := range field {
			if mult {
				fields += ", "
			}
			fields += value
			mult = true
		}
		return (fields)
	}
	return "*"
}

func arrToStr(arr []string) string {
	var str string

	for _, value := range arr {
		str += value
	}
	return (str)
}

func checkNumber(arr []string, def string) string {
	_, err := strconv.Atoi(arr[0])

	if err == nil {
		return (arr[0])
	}
	return (def)
}
