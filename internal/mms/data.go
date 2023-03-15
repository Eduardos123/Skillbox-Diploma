package mmss

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DwarfWizzard/skillbox-diplom-go/internal/codes"
)

func GetData(url string) ([]MMSData, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	respone, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer respone.Body.Close()

	if err != nil {
		return nil, err
	}
	if respone.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d", respone.StatusCode)
	}

	bytes, err := io.ReadAll(respone.Body)
	if err != nil {
		return nil, err
	}

	var data []MMSData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	var result []MMSData
	for _, mms := range data {
		if !codes.IsExist(mms.Country) ||
			(mms.Provider != "Topolo" && mms.Provider != "Rond" && mms.Provider != "Kildy") {
			continue
		}

		result = append(result, mms)
	}

	return result, nil
}
