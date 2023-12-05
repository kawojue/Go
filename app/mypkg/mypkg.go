package stuff

import "strconv"

var Name string = "Kawojue"

func IntArrToStrArr(intArr ...int) []string {
	strArr := []string{}

	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}

	return strArr
}
