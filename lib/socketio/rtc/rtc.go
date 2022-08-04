package rtc

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DeleteChannel(sn string) {
	log.Println("close RTC")
	queryList := make(map[string]string)
	queryList["Action"] = "DeleteChannel"
	queryList["AppId"] = appID
	queryList["ChannelId"] = sn
	GET("https://rtc.aliyuncs.com", nil, queryList)
}

func GET(host string, headerList, queryList map[string]string) ([]byte, error) {
	da := []byte{}
	client := &http.Client{}
	req, err := http.NewRequest("GET", host, nil)
	if err != nil {
		return da, err
	}
	for k, v := range headerList {
		req.Header.Add(k, v)
	}
	q := req.URL.Query()
	for k, v := range queryList {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return da, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return da, err
	}
	return b, nil
}
