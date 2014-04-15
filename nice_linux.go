package runnice

import "bytes"
import "syscall"

func (c *Cmd) CombinedOutput() ([]byte, error) {
	var b bytes.Buffer
	c.Stdout = &b
	c.Stderr = &b
	e1 := unwrap(c).Start()
	syscall.Setpriority(syscall.PRIO_PROCESS, c.Process.Pid, 10)
	e2 := unwrap(c).Wait()
	return b.Bytes(), ecomb(e1, e2)
}

func ecomb(e1, e2 error) error {
	if e1 != nil {
		return e1
	}
	return e2
}
