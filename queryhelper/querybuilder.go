package queryhelper

var toSql string

func Add(queryTxt string)  {
	toSql += queryTxt	
}

func AddIfNotEmpty(obj interface{}, queryTxt string) {
	toSql += queryTxt
}

func ToSQL() string {
	return toSql
}