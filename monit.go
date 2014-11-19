package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

type Process struct {
	pid  int64
	tty  string
	time string
	cmd  string
}

func PsOutput() string {
	cmd := exec.Command("ps")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func ParseOutput(output string) []*Process {
	re := regexp.MustCompile("(?m)^\\s*(?P<PID>[0-9]+)\\s+(?P<TTY>[a-z0-9]+)\\s+(?P<TIME>\\S+)\\s+(?P<CMD>.+)$")
	processes := []*Process{}
	matches := re.FindAllStringSubmatch(output, -1)
	for _, proc := range matches {
		pid, _ := strconv.ParseInt(proc[1], 10, 64)
		process := Process{
			pid:  pid,
			tty:  proc[2],
			time: proc[3],
			cmd:  proc[4],
		}
		processes = append(processes, &process)
	}
	return processes
}

func PrintProcesses(procs []*Process) string {
	var buffer bytes.Buffer
	for _, proc := range procs {
		buffer.WriteString(fmt.Sprintf("%s: %s\n", proc.time, proc.cmd))
	}
	return buffer.String()
}

func main() {
	processes := ParseOutput(PsOutput())
	fmt.Println(PrintProcesses(processes))
	fmt.Printf("Number processes: %d", len(processes))
}
