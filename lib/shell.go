package nx

import (
	"github.com/google/shlex"
	"log"
	"os/exec"
)

type ShellOptions struct {
	Command string
	Verbose bool
	DryRun  bool
  OnBeforeRun func(cmdArgs []string) error
  OnAfterRun func(cmdArgs []string, output string, err error) error
}

func RunShell(options *ShellOptions) (string, error) {
	log.SetPrefix("nx.shell: ")

	if options == nil {
		options = &ShellOptions{
			DryRun:  false,
			Verbose: false,
		}
	}

	command := options.Command
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

  if options.OnBeforeRun!= nil {
    if err := options.OnBeforeRun(cmdArgs); err!= nil {
      return "", err
    }
  }

	cmd := exec.Command(cmdName, cmdArgs[1:]...) // 第一个元素是命令，后面的元素是参数
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

  if options.OnAfterRun!= nil {
    if err := options.OnAfterRun(cmdArgs, string(out), err); err!= nil {
      return "", err
    }
  }

	return string(out), nil
}

func IsCommandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
