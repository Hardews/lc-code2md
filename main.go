/**
 * @Author: Hardews
 * @Date: 2023/7/20 12:26
 * @Description:
**/

package main

import (
	"lc-code2md/config"
	"lc-code2md/logic"
)

func main() {
	config.ReloadCookie()
	config.ReloadReplaceMap()
	logic.FindAndReplace()
}
