package app

import (
	"net/http"
	"regexp"
	"strconv"
	"fmt"
)

func GetRange(name string, r *http.Request) (int, int, error) {
	if _, ok := r.Header["Range"]; !ok {
		return -1, -1, nil
	}
	pattern := "^" + name + "=(\\d+)\\.\\.(\\d+)$"
	re := regexp.MustCompile(pattern)
	rangeValue := r.Header["Range"][0]
	if !re.MatchString(rangeValue) {
		return -1, -1, fmt.Errorf("The Range header value '%s' is invalid.", rangeValue)
	}
	matches := re.FindStringSubmatch(rangeValue)
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])
	return from, to, nil
}

