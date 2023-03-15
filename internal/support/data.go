package supports

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetData(url string) ([]SupportData, error) {
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

	var data []SupportData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}