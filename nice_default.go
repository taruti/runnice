
// +build !linux

package runnice

import "os/exec"

func (c *Cmd) CombinedOutput() ([]byte, error) {
	return unwrap(c).CombinedOutput()
}
