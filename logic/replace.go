/**
 * @Author: Hardews
 * @Date: 2023/7/20 12:28
 * @Description: 将得到的数据进行替换
**/

package logic

import (
	"fmt"
	"io"
	"io/ioutil"
	"lc-code2md/config"
	"lc-code2md/request"
	"log"
	"os"
	"regexp"
	"strings"
)

func FindAndReplace() {
	files, err := ioutil.ReadDir("./file")
	if err != nil {
		log.Fatalln("read dir failed,err:", err)
	}

	for _, file := range files {
		f, err := os.Open("./file/" + file.Name())
		if err != nil {
			log.Printf("open file failed,err:%s,file name:%s", err.Error(), file.Name())
			return
		}

		var fileContent []byte
		for {
			var temp = make([]byte, 1024)
			n, err := f.Read(temp)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("read file failed,err:", err)
				return
			}
			fileContent = append(fileContent, temp[:n]...)
		}

		fileContent = codeGetAndReplace(fileContent)
		// 去掉不必要的字段
		fileContent = []byte(frameReplace(string(fileContent)))

		fileContent = []byte(normalWordReplace(string(fileContent)))

		fileContent = []byte(quotationReplace(string(fileContent)))

		// 创建结果文件
		ansFile, err := os.Create("./answer/ans_" + file.Name())
		if err != nil {
			log.Printf("create file failed,err:%s", err.Error())
			return
		}

		ansFile.Write(fileContent)
	}
}

// 替换 frame 标签的东西
func codeGetAndReplace(fileContent []byte) []byte {
	// 定义正则表达式
	iframeRegex := regexp.MustCompile(`<iframe src=""https://leetcode\.com/playground/(.*?)/shared""(.*?)></iframe>`)

	// 在文件内容中查找匹配的内容
	matches := iframeRegex.FindAllStringSubmatch(string(fileContent), -1)

	// 处理匹配结果

	for _, match := range matches {
		if len(match) == 3 {
			slug := match[1] // 获取到的 slug
			cj := request.GetCode(slug)
			fmt.Println(cj)

			var res string
			for _, code := range cj.Data.AllPlaygroundCodes {
				res += "```" + langReplace(code.LangSlug) + " [solution]\n"
				res += code.Code
				res += "\n```\n"
			}
			res += "\n"

			/*
			 <iframe src=""https://leetcode.com/playground/eAZhmD7R/shared""
			*/
			replaceInput := "<iframe src=\"\"https://leetcode.com/playground/" + slug + "/shared\"\""
			regex := regexp.MustCompile(replaceInput)

			// 使用正则表达式进行替换
			result := regex.ReplaceAllString(string(fileContent), res)
			fileContent = []byte(result)
		}
	}

	return fileContent
}

// 清除 frame 替换后的剩余字
func frameReplace(fileContent string) string {
	/*
		frameBorder=""0"" width=""100%"" height=""344"" name=""eAZhmD7R""></iframe>
	*/
	var length = len("frameBorder=\"\"0\"\" width=\"\"100%\"\" height=\"\"344\"\" name=\"\"eAZhmD7R\"\"></iframe>")
	for true {
		frameIndex := strings.Index(fileContent, "frameBorder")
		if frameIndex == -1 {
			break
		}
		fileContent = strings.ReplaceAll(fileContent, fileContent[frameIndex-1:frameIndex+length], "")
	}

	return fileContent
}

func quotationReplace(content string) string {
	return strings.TrimSuffix(strings.TrimPrefix(content, "\""), "\"")
}

// 常规翻译问题的转换
func normalWordReplace(content string) string {
	for word, replaceWord := range config.ReplaceNormalMap {
		temp := strings.ReplaceAll(content, word, replaceWord)
		content = temp
	}

	return content
}

func langReplace(lang string) string {
	if l, ok := map[string]string{
		"javascript": "JavaScript",
		"java":       "Java",
		"typescript": "TypeScript",
		"python":     "Python",
		"python3":    "Python3",
		"cpp":        "C++",
		"c":          "C",
		"golang":     "Golang",
	}[lang]; ok {
		return l
	}
	return lang
}
