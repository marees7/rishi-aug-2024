package helpers

import (
	"github.com/marees7/rishi-aug-2024/common/constants"
	"strconv"
)

// converts the string limit and offset into int and calculates the offset
func Pagination(limitStr string, offsetStr string) (int, int, error) {
	offset, err := strconv.Atoi(offsetStr)
	if offsetStr == "" {
		offset = constants.DefaultOffset
	} else if err != nil {
		return 0, 0, err
	}

	limit, err := strconv.Atoi(limitStr)
	if limitStr == "" {
		limit = constants.DefaultLimit
	} else if err != nil {
		return 0, 0, err
	}
	
	offset = (offset - 1) * limit

	return limit, offset, nil
}
