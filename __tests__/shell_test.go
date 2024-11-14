/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-11-14 22:07:52
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-11-14 22:09:00
 * @FilePath: __tests__/shell_test.go
 * @Description: 这是默认设置,可以在设置》工具》File Description中进行配置
 */
package nx_test

import (
	nx "github.com/afeiship/nx/lib"
	"testing"
)

func TestIsCommandExists(f *testing.T) {
	res := nx.IsCommandExists("ls")
	if res == false {
		f.Error("ls command not exists")
	}

	res = nx.IsCommandExists("not_exists_command")
	if res == true {
		f.Error("not_exists_command command exists")
	}
}
