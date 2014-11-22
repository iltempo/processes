// Package that provides a list of running processes in the system.
package processes

import (
	"bytes"
	"os/exec"
	"regexp"
	"strconv"
)

// Process represents a running process including it's meta data.
type Process struct {
	Pid  int64
	Tty  string
	Time string
	Cmd  string
}

// ByTime allows sorting of processes by CPU time.
type ByTime []*Process

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time < a[j].Time }

// runCommand captures the standard output when executing a given command.
func runCommand(command string) string {
	cmd := exec.Command(command)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	return out.String()
}

// OfCurrentUser provides a list of all processes of the current user.
func OfCurrentUser(predefined_output ...string) []*Process {
	var output string
	if len(predefined_output) >= 1 {
		output = predefined_output[0]
	} else {
		output = runCommand("ps")
	}

	re := regexp.MustCompile("(?m)^\\s*(?P<PID>[0-9]+)\\s+(?P<TTY>[a-z0-9]+)\\s+(?P<TIME>\\S+)\\s+(?P<CMD>.+)$")
	processes := []*Process{}
	matches := re.FindAllStringSubmatch(output, -1)
	for _, proc := range matches {
		pid, _ := strconv.ParseInt(proc[1], 10, 64)
		process := Process{
			Pid:  pid,
			Tty:  proc[2],
			Time: proc[3],
			Cmd:  proc[4],
		}
		processes = append(processes, &process)
	}
	return processes
}
