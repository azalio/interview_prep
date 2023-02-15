package short_way

import (
	"fmt"
	"strings"
)

//дан корректный путь для файла или директории,
//содержащий в том числе конструкции вида ./ ../,
//нужно упростить путь, насколько возможно
//Пример "/a/b/c/./../../d" -> "/a/d"

func shortWay(s string) string {

	var sResult []string

	ss := strings.Split(s, "/")

	for _, c := range ss {
		if c == "." {
			continue
		}
		if c == ".." {
			sResult = sResult[:len(sResult)-1]
			fmt.Println(sResult)
			continue
		}
		sResult = append(sResult, c)
	}

	result := strings.Join(sResult, "/")
	return result

}
