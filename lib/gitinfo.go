/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:05:50
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-15 21:16:13
 */
package nx

func GitInfo() string {
	out, _ := RunShell("git rev-parse --abbrev-ref HEAD")
	return out
}
