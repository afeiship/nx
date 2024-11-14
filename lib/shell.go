package nx

import (
	"github.com/google/shlex"
	"log"
	"os/exec"
)

type ShellOptions struct {
	Verbose bool
	DryRun  bool
}

func RunShell(command string, options *ShellOptions) (string, error) {
	log.SetPrefix("nx.shell: ")

	cmdArgs, err := shlex.Split(command)
	if err != nil {
		return "", err
	}

	cmdName := cmdArgs[0]
	if !IsCommandExists(cmdName) {
		return "", nil
	}

	if options.Verbose {
		log.Println("Running command: ", command)
	}

	if options.DryRun {
		log.Println("Dry run, not running command: ", command)
		return "", nil
	}

	cmd := exec.Command(cmdName, cmdArgs[1:]...) // 第一个元素是命令，后面的元素是参数
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func IsCommandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
