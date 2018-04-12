// parse domain
// construct clone URL (https://[domain][user][project])
// set path to $GITROOT/src/[domain]/[user]/[project] (reverse domain)
// (don't need) mkdir -p to path
// execute git clone into path
// pipe output to stdout
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"syscall"
	"time"
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

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Minute,
	)
	defer cancel()

	cmd := exec.CommandContext(ctx, "git", "clone", remoteurl, localdir)

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start: %v", err)
	}

	// Exit status capturing reference:
	// https://stackoverflow.com/a/10385867/2684355
	if err := cmd.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {

			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				switch status.ExitStatus() {
				case 128:
					fmt.Println("Error: destination path already exists")
				default:
					log.Fatalf("Failed for unaccounted reason")
				}
			}
		} else {
			log.Fatalf("Failed for unaccounted reason")
		}
	} else {
		fmt.Print(localdir)
	}
}
