// parse domain
// construct clone URL (https://[domain][user][project])
// set path to $GITROOT/src/[domain]/[user]/[project] (reverse domain)
// (don't need) mkdir -p to path
// execute git clone into path
// pipe output to stdout
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No package specified")
		return
	}

	gitpath := os.Getenv("GITPATH")

	if gitpath == "" {
		fmt.Println("No GITPATH specified")
		return
	}

	name := os.Args[1]

	remoteurl := "https://" + name

	pieces := strings.Split(name, "/")

	withroot := append([]string{gitpath, "src"}, pieces...)

	localdir := path.Join(withroot...)

	cmd := exec.Command("git", "clone", remoteurl, localdir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func reverse(pieces []string) []string {
	length := len(pieces)
	reversed := make([]string, length)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = pieces[j], pieces[i]
	}
	if length%2 == 1 {
		reversed[length/2] = pieces[length/2]
	}
	return reversed
}
