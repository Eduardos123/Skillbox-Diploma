package smss

import (
	"os"
	"strings"

	"github.com/DwarfWizzard/skillbox-diplom-go/internal/codes"
)

func GetData(path string) ([]SMSData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	content := strings.Split(string(data), "\n")

	var result []SMSData
	for _, record := range content {
		cortege := strings.Split(record, ";")
		if len(cortege) != 4 || !codes.IsExist(cortege[0]) ||
			(cortege[3] != "Topolo" && cortege[3] != "Rond" && cortege[3] != "Kildy") {
			continue
		}

		result = append(result, SMSData{
			Country:      cortege[0],
			Bandwidth:    cortege[1],
			ResponseTime: cortege[2],
			Provider:     cortege[3],
		})
	}

	return result, nil
}
