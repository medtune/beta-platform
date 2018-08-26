package api

import "fmt"

func staticPath(demo string, file string) string {
	return fmt.Sprintf("static/demos/%s/images/%s", demo, file)
}
