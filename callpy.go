package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python3", "script.py", "arg1", "arg2")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println("output:", string(output))
}
