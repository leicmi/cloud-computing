package util

var (
	JOB_STATUS_PENDING = "PENDING"
	JOB_STATUS_RUNNING = "RUNNING"
	JOB_STATUS_DONE    = "DONE"
)

type Job struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}
