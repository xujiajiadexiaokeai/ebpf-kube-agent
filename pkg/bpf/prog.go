package bpf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type ProgInfo struct {
	Name string
	Type string
}

func comm_for_pid(pid int) string {
	var comm string
	err := filepath.Walk(fmt.Sprintf("/proc/%d/comm", pid), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			f, e := os.Open(path)
			if e != nil {
				return e
			}
			defer f.Close()
			b, e := ioutil.ReadAll(f)
			if e != nil {
				return e
			}
			comm = string(b)
		}
		return nil
	})
	if err != nil {
		return "[unknown]"
	}
	return comm
}

type Proc struct {
	Pid  int
	Type string
}

func findBpfFds(pid int, counts map[Proc]int) error {
	root := fmt.Sprintf("/proc/%d/fd", pid)
	files, _ := ioutil.ReadDir(root)
	for _, fd := range files {
		link, _ := os.Readlink(filepath.Join(root, fd.Name()))
		re := regexp.MustCompile(`anon_inode:bpf-([\w-]+)`)
		match := re.FindStringSubmatch(link)
		if len(match) > 0 {
			proc := Proc{
				Pid:pid, 
				Type:match[1],
			}
			if _, ok := counts[proc]; ok {
				counts[proc] = counts[proc] + 1
			} else {
				counts[proc] = 1
			}
		}
	}
	return nil
}

// list all eBPF programs currently running on the system
func ListBPFProgs() {
	f, _ := os.Open("/proc")
	defer f.Close()
	counts := make(map[Proc]int)
	pdirs, _ := f.Readdirnames(0)
	for _, pdir := range pdirs {
		if matched, _ := regexp.MatchString("\\d+", pdir); matched {
			pid, err := strconv.Atoi(pdir)
			if err != nil {
				continue
			} else {
				findBpfFds(pid, counts)
			}
		}
	}

	fmt.Printf("%-6s %-16s %-6s %s\n","PID", "COMM", "TYPE", "COUNT")
	for proc, count := range counts {
		comm := comm_for_pid(proc.Pid)
		fmt.Printf("%-6d %-16s %-6s %d\n", proc.Pid, strings.TrimSpace(comm), proc.Type, count)
	}
}
