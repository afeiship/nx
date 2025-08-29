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
	"testing"

	nx "github.com/afeiship/nx/lib"
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

func TestRunShell_MultiCommands_AND(t *testing.T) {
	opts := &nx.ShellOptions{
		Commands: []string{"echo hi", "ls"},
		Mode:     nx.ShellModeAnd,
	}
	out, err := nx.RunShell(opts)
	if err != nil {
		t.Errorf("and模式下所有命令都应成功, got err: %v", err)
	}
	if out == "" {
		t.Error("and模式下应返回最后一个命令输出")
	}
}

func TestRunShell_MultiCommands_OR(t *testing.T) {
	opts := &nx.ShellOptions{
		Commands: []string{"not_exists_command", "echo ok"},
		Mode:     nx.ShellModeOr,
	}
	out, err := nx.RunShell(opts)
	if err != nil {
		t.Errorf("or模式下只要有一个命令成功就应通过, got err: %v", err)
	}
	if out == "" || out != "ok\n" {
		t.Errorf("or模式下应返回第一个成功命令输出, got: %q", out)
	}
}

func TestRunShell_MultiCommands_AllFail(t *testing.T) {
	opts := &nx.ShellOptions{
		Commands: []string{"not_exists_command", "also_bad_cmd"},
		Mode:     nx.ShellModeOr,
	}
	out, err := nx.RunShell(opts)
	if err == nil {
		t.Error("全部失败时应返回错误")
	}
	if out != "" {
		t.Error("全部失败时输出应为空")
	}
}
