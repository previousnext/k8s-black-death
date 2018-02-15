package retention

import (
	"time"
	"strconv"
)

// Unix returns a unix string based on a retention period.
func Unix(retention string) (string, error) {
	duration, err := time.ParseDuration(retention)
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(time.Now().Local().Add(duration).Unix(), 10), nil
}
