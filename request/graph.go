/**
 * @Author: Hardews
 * @Date: 2023/7/20 12:27
 * @Description: 发送 graphQL 请求，拿到数据
**/

package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lc-code2md/config"
	"net/http"
	"strings"
	"time"
)

type CodeJson struct {
	Data struct {
		AllPlaygroundCodes []struct {
			Code     string `json:"code"`
			LangSlug string `json:"langSlug"`
		} `json:"allPlaygroundCodes"`
	} `json:"data"`
}

func GetCode(slug string) (cj CodeJson) {
	url := "https://leetcode.com/graphql"
	method := "GET"

	payload := strings.NewReader("{\"query\":\"query fetchPlayground {\\r\\n  playground(uuid: \\\"" + slug + "\\\") {\\r\\n    testcaseInput\\r\\n    name\\r\\n    isUserOwner\\r\\n    isLive\\r\\n    showRunCode\\r\\n    showOpenInPlayground\\r\\n    selectedLangSlug\\r\\n    isShared\\r\\n    __typename\\r\\n  }\\r\\n  allPlaygroundCodes(uuid: \\\"" + slug + "\\\") {\\r\\n    code\\r\\n    langSlug\\r\\n    __typename\\r\\n  }\\r\\n}\",\"variables\":{}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	cookie := config.Cookie
	req.Header.Add("Cookie", cookie)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1823.82")
	req.Header.Add("Referer", "https://leetcode.com/playground/"+slug+"/shared")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.Unmarshal(body, &cj)
	time.Sleep(2 * time.Second)
	return
}
