/**
 * @Author: aric 1290657123@qq.com
 * @Date: 2024-10-15 21:09:48
 * @LastEditors: aric 1290657123@qq.com
 * @LastEditTime: 2024-10-15 21:11:13
 */
package nx

import "os/exec"

func RunShell(cmd string) (string, error) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
