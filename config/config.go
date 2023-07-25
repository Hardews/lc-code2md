/**
 * @Author: Hardews
 * @Date: 2023/7/20 12:31
 * @Description:
**/

package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var (
	Cookie           string
	ReplaceNormalMap map[string]string
)

func ReloadCookie() {
	f, err := os.Open("./config/config.txt")
	if err != nil {
		log.Println("reload cookie file failed,err:", err)
		return
	}

	var res []byte
	for {
		var temp = make([]byte, 5*1024)
		n, err := f.Read(temp)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("read file failed,err:", err)
			return
		}
		res = append(res, temp[:n]...)
	}

	Cookie = string(res)
}

func ReloadReplaceMap() {
	ReplaceNormalMap = make(map[string]string)
	f, err := os.Open("./config/replace.json")
	if err != nil {
		log.Println("reload replace file failed,err:", err)
		return
	}

	var res []byte
	for {
		var temp = make([]byte, 5*1024)
		n, err := f.Read(temp)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("read file failed,err:", err)
			return
		}
		res = append(res, temp[:n]...)
	}

	json.Unmarshal(res, &ReplaceNormalMap)
}
