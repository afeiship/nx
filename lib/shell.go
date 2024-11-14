package nx

import (
	"github.com/google/shlex"
	"os/exec"
)

func RunShell(command string) (string, error) {
	cmdArgs, err := shlex.Split(command)
	if err != nil {
		return "", err
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...) // 第一个元素是命令，后面的元素是参数
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
