package cmd

import (
	"github.com/jedib0t/go-pretty/table"
	"github.com/leicmi/cloud-computing/util"
	"github.com/sirupsen/logrus"
)

func list(url string) {
	jobs, err := util.ListJobs(url)
	if err != nil {
		logrus.WithError(err).Errorf("unable to list pending jobs")
	}

	util.PrintTable(jobs, table.Row{"jobname", "status"}, func(job util.Job) table.Row {
		return table.Row{job.ID, job.Status}
	})
}
