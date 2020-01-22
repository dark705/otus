package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func ReadDir(dir string) (map[string]string, error) {
	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	env := make(map[string]string, len(filesInfo))
	for _, fileInfo := range filesInfo {
		value, err := ioutil.ReadFile(dir + fileInfo.Name())
		if err != nil {
			return env, err
		}
		env[fileInfo.Name()] = string(value)
	}
	return env, nil
}

func RunCmd(cmd []string, env map[string]string) int {
	if len(cmd) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Not set command for run\n")
		return 2
	}

	c := exec.Command(cmd[0])
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if len(cmd) >= 1 {
		c.Args = append(c.Args, cmd[1:]...)
	}

	if len(env) >= 1 {
		for e, v := range env {
			c.Env = append(c.Env, e+"="+v)
		}
	}

	err := c.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}

	return 0
}
