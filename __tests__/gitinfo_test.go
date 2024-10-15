/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:06:33
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-15 21:58:34
 */
package nx

import (
	"fmt"
	nx "github.com/afeiship/nx/lib"
	"testing"
)

func TestGitInfo(f *testing.T) {
	info := nx.GitInfo()
	fmt.Println(info.CurrentBranch)
}
