package main

import "fmt"

// placeholder for sql
func main() {
	where := "col1='a' and col2='b'"
	groupBy := "group by co1, col2"

	str := "select * from table1 where %s group by %s"
	str1 := fmt.Sprintf(str, where, groupBy)
	fmt.Println(str1)
}
