// Runs processes with less priority.
package runnice

import (
	"os/exec"
)

type Cmd exec.Cmd

func Command(name string, arg ...string) *Cmd {
	return (*Cmd)(exec.Command(name, arg...))
}

func unwrap(c *Cmd) *exec.Cmd {
	return (*exec.Cmd)(c)
}
