/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:06:33
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-11-15 09:20:08
 */

// go test -v __tests__/gitinfo_test.go -run TestGitInfo
package nx

import (
	"encoding/json"
	"fmt"
	"testing"

	nx "github.com/afeiship/nx/lib"
)

func TestGitInfo(f *testing.T) {
	info := nx.GitInfo()
	jsonString, _ := json.Marshal(info)
	fmt.Println(string(jsonString))
	fmt.Println(info.CurrentBranch)
	fmt.Println(info.LatestVersion)
}
