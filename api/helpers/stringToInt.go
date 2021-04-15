package helpers

import (
	"fmt"
	"strconv"
)

// ConvertStringToInt convert a string to int and return int
func ConvertStringToInt(routeParam string) uint64 {
	var err error
	var id uint64
	id, err = strconv.ParseUint(routeParam, 10, 32)

	fmt.Println(err)
	return id
}
