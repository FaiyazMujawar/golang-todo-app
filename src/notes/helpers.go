package notes

import "strings"

func extractObjectKeyFromUrl(url string) string {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}
