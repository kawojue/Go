package stuff

import "strings"

func ToTitle(name string) string {
	nameArr := []string{}

	for _, n := range strings.Split(name, " ") {
		bytesName := strings.Split(n, "")
		initial := strings.ToUpper(bytesName[0])
		others := strings.ToLower(strings.Join(bytesName[1:], ""))
		title := initial + others
		nameArr = append(nameArr, title)
	}

	return strings.Trim(strings.Join(nameArr, " "), "")
}
