package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		path, _ := os.Getwd()
		fmt.Print(path)
		return
	}

	fmt.Print(string(out))
}
