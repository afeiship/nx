package nx

import (
	"log"
	"os/exec"

	"github.com/google/shlex"
)

type ShellMode string

const (
	ShellModeAnd ShellMode = "and" // 所有命令都需要成功
	ShellModeOr  ShellMode = "or"  // 任一命令成功即可
)

type ShellOptions struct {
	Command     string    // 单条命令，兼容旧用法
	Commands    []string  // 多条命令
	Mode        ShellMode // 执行模式：and 或 or
	Verbose     bool
	DryRun      bool
	OnBeforeRun func(cmdArgs []string) error
	OnAfterRun  func(cmdArgs []string, output string, err error) error
}

func RunShell(options *ShellOptions) (string, error) {
	log.SetPrefix("nx.shell: ")

	if options == nil {
		options = &ShellOptions{
			DryRun:  false,
			Verbose: false,
		}
	}

	// 兼容单条命令
	var commands []string
	if len(options.Commands) > 0 {
		commands = options.Commands
	} else if options.Command != "" {
		commands = []string{options.Command}
	} else {
		return "", nil
	}

	mode := options.Mode
	if mode != ShellModeOr && mode != ShellModeAnd {
		mode = ShellModeAnd // 默认 and
	}

	var lastErr error
	var lastOut string
	for idx, command := range commands {
		cmdArgs, err := shlex.Split(command)
		if err != nil {
			lastErr = err
			if mode == ShellModeAnd {
				return "", err
			}
			continue
		}

		cmdName := cmdArgs[0]
		if !IsCommandExists(cmdName) {
			lastErr = exec.ErrNotFound
			if mode == ShellModeAnd {
				return "", lastErr
			}
			continue
		}

		if options.Verbose {
			log.Printf("[%d/%d] Running command: %s\n", idx+1, len(commands), command)
		}

		if options.DryRun {
			log.Printf("[%d/%d] Dry run, not running command: %s\n", idx+1, len(commands), command)
			continue
		}

		if options.OnBeforeRun != nil {
			if err := options.OnBeforeRun(cmdArgs); err != nil {
				lastErr = err
				if mode == ShellModeAnd {
					return "", err
				}
				continue
			}
		}

		cmd := exec.Command(cmdName, cmdArgs[1:]...)
		out, err := cmd.CombinedOutput()

		if options.OnAfterRun != nil {
			if cbErr := options.OnAfterRun(cmdArgs, string(out), err); cbErr != nil {
				lastErr = cbErr
				if mode == ShellModeAnd {
					return "", cbErr
				}
				continue
			}
		}

		if err != nil {
			lastErr = err
			if mode == ShellModeAnd {
				return "", err
			}
			continue
		}

		lastOut = string(out)
		// or 模式下，遇到第一个成功就返回
		if mode == ShellModeOr {
			return lastOut, nil
		}
	}
	// and 模式下，全部成功才返回最后结果，否则返回最后错误
	if mode == ShellModeAnd {
		if lastErr != nil {
			return "", lastErr
		}
		return lastOut, nil
	}
	// or 模式下，全部失败
	return "", lastErr
}

func IsCommandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
