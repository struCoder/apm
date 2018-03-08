package process

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/struCoder/pidusage"
	"github.com/danielckv/pmgo/lib/utils"
)

// ProcStatus is a wrapper with the process current status.
type ProcStatus struct {
	Status    string
	Restarts  int
	StartTime int64
	Uptime    string
	Sys       *pidusage.SysInfo
}

// SetStatus will set the process string status.
func (proc_status *ProcStatus) SetStatus(status string) {
	proc_status.Status = status
}

// AddRestart will add one restart to the process status.
func (proc_status *ProcStatus) AddRestart() {
	proc_status.Restarts++
}

// InitUptime will record proc start time
func (proc_status *ProcStatus) InitUptime() {
	proc_status.StartTime = time.Now().Unix()
}

// SetUptime will figure out process uptime
func (proc_status *ProcStatus) SetUptime() {
	proc_status.Uptime = utils.FormatUptime(proc_status.StartTime, time.Now().Unix())
}

// ResetUptime will Reset uptime
func (proc_status *ProcStatus) ResetUptime() {
	proc_status.Uptime = "0s"
}

// SetSysInfo will get current proc cpu and memory usage
func (proc_status *ProcStatus) SetSysInfo(pid int) {
	var err error
	proc_status.Sys, err = pidusage.GetStat(pid)
	if err != nil {
		log.Error(err)
	}
}
