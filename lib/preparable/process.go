package preparable

import (
	"os/exec"
	"strings"

	"github.com/danielckv/pmgo/lib/process"
)

// ProcPreparable is a preparable with all the necessary informations to run
// a process. To actually run a process, call the Start() method.
type ProcPreparable interface {
	PrepareBin() ([]byte, error)
	Start() (process.ProcContainer, error)
	getPath() string
	Identifier() string
	getBinPath() string
	getPidPath() string
	getOutPath() string
	getErrPath() string
}

type Preparable struct {
	Name       string
	SourcePath string
	Cmd        string
	SysFolder  string
	Language   string
	KeepAlive  bool
	Args       []string
}

// PrepareBin will compile the Golang project from SourcePath and populate Cmd with the proper
// command for the process to be executed.
// Returns the compile command output.
func (preparable *Preparable) PrepareBin() ([]byte, error) {
	// Remove the last character '/' if present
	if preparable.SourcePath[len(preparable.SourcePath)-1] == '/' {
		preparable.SourcePath = strings.TrimSuffix(preparable.SourcePath, "/")
	}
	cmd := ""
	cmdArgs := []string{}
	binPath := preparable.getBinPath()
	if preparable.Language == "go" {
		cmd = "go"
		cmdArgs = []string{"build", "-o", binPath, preparable.SourcePath + "/."}
	}

	preparable.Cmd = preparable.getBinPath()
	return exec.Command(cmd, cmdArgs...).Output()
}

// Start will execute the process based on the information presented on the preparable.
// This function should be called from inside the master to make sure
// all the watchers and process handling are done correctly.
// Returns a tuple with the process and an error in case there's any.
func (preparable *Preparable) Start() (process.ProcContainer, error) {
	proc := &process.Proc{
		Name:      preparable.Name,
		Cmd:       preparable.Cmd,
		Args:      preparable.Args,
		Path:      preparable.getPath(),
		Pidfile:   preparable.getPidPath(),
		Outfile:   preparable.getOutPath(),
		Errfile:   preparable.getErrPath(),
		KeepAlive: preparable.KeepAlive,
		Status:    &process.ProcStatus{},
	}

	err := proc.Start()
	return proc, err
}

// Identifier is a function that get proc name
func (preparable *Preparable) Identifier() string {
	return preparable.Name
}

func (preparable *Preparable) getPath() string {
	if preparable.SysFolder[len(preparable.SysFolder)-1] == '/' {
		preparable.SysFolder = strings.TrimSuffix(preparable.SysFolder, "/")
	}
	return preparable.SysFolder + "/" + preparable.Name
}

func (preparable *Preparable) getBinPath() string {
	return preparable.getPath() + "/" + preparable.Name
}

func (preparable *Preparable) getPidPath() string {
	return preparable.getBinPath() + ".pid"
}

func (preparable *Preparable) getOutPath() string {
	return preparable.getBinPath() + ".out"
}

func (preparable *Preparable) getErrPath() string {
	return preparable.getBinPath() + ".err"
}
