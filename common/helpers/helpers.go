package helpers

import (
	"blogs/common/constants"
	"strconv"
)

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

	return limit, offset, nil
}
