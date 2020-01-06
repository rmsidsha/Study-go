package test_package

import "fmt"

var pop map[string]string

func init() {
	pop = make(map[string]string)
	pop["seo"] = "hello"
	pop["go"] = "start"
	pop["hard"] = "everything"
}

func getString(keyword string) string {
	return pop[keyword]
}

func getKeys() {
	for _, key := range pop {
		fmt.Println(key)
	}
}
