package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ExecuteCommand runs a command and returns its output as a string
func ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String() + stderr.String(), err
	}

	return stdout.String(), nil
}

// ClearScreen clears the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

// JoinArgs joins command arguments for display
func JoinArgs(args []string) string {
	var quotedArgs []string

	for _, arg := range args {
		if strings.Contains(arg, " ") {
			quotedArgs = append(quotedArgs, fmt.Sprintf("%q", arg))
		} else {
			quotedArgs = append(quotedArgs, arg)
		}
	}

	return strings.Join(quotedArgs, " ")
}

// ExtractPackageName extracts the package name from source code
func ExtractPackageName(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "package ") {
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

// ExtractTestFunctionNames extracts test function names from code
func ExtractTestFunctionNames(content string) []string {
	var names []string

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "func Test") {
			parts := strings.SplitN(line, "(", 2)
			if len(parts) == 2 {
				name := strings.TrimPrefix(parts[0], "func ")
				names = append(names, name)
			}
		}
	}

	return names
}
