package processes

import (
	"regexp"
	"sort"
	"testing"
)

const examplePsOutput = `  PID TTY           TIME CMD
	  931 ttys000    0:00.07 -fish
		 1035 ttys000    0:00.01 tmux
		 10151 ttys002    0:01.80 -fish
		 13771 ttys002    0:00.27 vim processes_test.go
		  8103 ttys003    0:04.18 -fish
			10063 ttys003    0:00.68 vim netps.go`

func TestCallingPs(t *testing.T) {
	output := runCommand("ps")
	match, _ := regexp.MatchString("\\s*PID", output)
	if match == false {
		t.Errorf("ps command output unexpected:\n%v", output)
	}
}

func TestFetchingRunningProcesses(t *testing.T) {
	procs := OfCurrentUser()
	if len(procs) == 0 {
		t.Errorf("Expected more than 0 processes")
	}
}

func TestParsingSpecificPsOutput(t *testing.T) {
	length := 6
	procs := OfCurrentUser(examplePsOutput)
	if len(procs) != length {
		t.Errorf("Expected %d processes", length)
	}
}

func TestProcessAttributes(t *testing.T) {
	procs := OfCurrentUser(examplePsOutput)
	proc := procs[0]
	if proc.Pid != 931 {
		t.Errorf("Unexpected pid: %d", proc.Pid)
	}
	if proc.Tty != "ttys000" {
		t.Errorf("Unexpected tty: %v", proc.Tty)
	}
	if proc.Time != "0:00.07" {
		t.Errorf("Unexpected time: %v", proc.Time)
	}
	if proc.Cmd != "-fish" {
		t.Errorf("Unexpected cmd: %v", proc.Cmd)
	}
}

func TestSortingProcesses(t *testing.T) {
	procs := OfCurrentUser(examplePsOutput)
	sort.Sort(ByTime(procs))
	if procs[5].Pid != 8103 {
		t.Errorf("Unexpected sorting order of processes")
	}
}
