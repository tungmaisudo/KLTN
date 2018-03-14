package Data

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func GetDataByUrl(url string) ([]Data, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var dataList []Data
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &dataList); err != nil {
		return nil, err
	}
	return dataList, nil
}
