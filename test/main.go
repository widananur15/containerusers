package main

import (
	"fmt"
	"github.com/widananur15/containerusers/queryhelper"
)

func main()  {
	queryhelper.Add(" SELECT * ")
	queryhelper.Add(" FROM t_user ")
	query := queryhelper.ToSQL()

	fmt.Println("Query String: ", query)

}
