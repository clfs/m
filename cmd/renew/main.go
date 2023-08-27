package main

import (
	"debug/buildinfo"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"text/tabwriter"
)

func main() {
	log.SetFlags(0)

	var (
		listFlag   = flag.Bool("list", false, "list binaries")
		updateFlag = flag.String("update", "", "update a binary to @latest")
	)
	flag.Parse()

	var err error

	switch {
	case *listFlag:
		err = list()
	case *updateFlag != "":
		err = update(*updateFlag)
	default:
		flag.Usage()
	}

	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func goInstallDir() (string, error) {
	// See "go help install" for Go's installation strategy.
	s, err := getGoEnv("GOBIN")
	if s != "" && err == nil {
		return s, nil
	}
	s, err = getGoEnv("GOPATH")
	if s != "" && err == nil {
		return filepath.Join(s, "bin"), nil
	}
	s, err = os.UserHomeDir()
	if s != "" && err == nil {
		return filepath.Join(s, "go", "bin"), nil
	}
	return "", fmt.Errorf("failed to determine go install dir: %v", err)
}

type binary struct {
	name       string // e.g. "foo"
	localPath  string // e.g. "/Users/calvin/go/bin/foo"
	importPath string // e.g. "github.com/bar/foo"
}

func binaries(dir string) ([]binary, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var bins []binary

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		localPath := filepath.Join(dir, e.Name())

		buildInfo, err := buildinfo.ReadFile(localPath)
		if err != nil {
			return nil, err
		}

		bins = append(bins, binary{
			name:       e.Name(),
			localPath:  localPath,
			importPath: buildInfo.Path,
		})
	}

	return bins, nil
}

// goEnv holds the results of "go env" run once.
var goEnv struct {
	sync.Once
	m   map[string]string
	err error
}

// getGoEnv returns the value of a key named in "go env".
func getGoEnv(key string) (string, error) {
	goEnv.Once.Do(func() {
		var out []byte
		out, goEnv.err = exec.Command("go", "env", "-json").Output()
		if goEnv.err != nil {
			return
		}
		goEnv.m = make(map[string]string)
		goEnv.err = json.Unmarshal(out, &goEnv.m)
	})

	if goEnv.err != nil {
		return "", goEnv.err
	}

	v, ok := goEnv.m[key]
	if !ok {
		return "", fmt.Errorf("key %s not in go env", key)
	}
	return v, nil
}

func list() error {
	dir, err := goInstallDir()
	if err != nil {
		return err
	}

	bins, err := binaries(dir)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for _, b := range bins {
		fmt.Fprintf(w, "%s\t%s\n", b.name, b.importPath)
	}
	return w.Flush()
}

func update(name string) error {
	dir, err := goInstallDir()
	if err != nil {
		return err
	}

	localPath := filepath.Join(dir, name)

	buildInfo, err := buildinfo.ReadFile(localPath)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "install", buildInfo.Path+"@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
