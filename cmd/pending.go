package cmd

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/leicmi/cloud-computing/util"
	"github.com/sirupsen/logrus"
)

func pending(url string) {
	jobs, err := util.QueryJobs(url, util.JOB_STATUS_PENDING)
	if err != nil {
		logrus.WithError(err).Errorf("unable to list pending jobs")
	}

	util.PrintTable(jobs, table.Row{"jobname", "status"}, func(job util.Job) table.Row {
		return table.Row{job.ID, job.Status}
	})
}
