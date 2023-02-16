package dummy

import "github.com/k0kubun/pp/v3"

func Dummy(name string) string {
	result := "Dummy " + name
	pp.Println(result)
	return result
}
