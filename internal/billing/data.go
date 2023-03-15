package billings

import (
	"os"
	"strconv"
)

func GetData(path string) (BillingData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return BillingData{}, err
	}

	var statuses []bool

	for i := len(data) - 1; i >= 0; i-- {
		status, err := strconv.ParseBool(string(data[i]))
		if err != nil {
			return BillingData{}, err
		}
		statuses = append(statuses, status)
	}

	billing := BillingData{
		statuses[0],
		statuses[1],
		statuses[2],
		statuses[3],
		statuses[4],
		statuses[5],
	}

	return billing, nil
}