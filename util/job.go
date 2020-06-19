package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	JOB_STATUS_PENDING = "PENDING"
	JOB_STATUS_RUNNING = "RUNNING"
	JOB_STATUS_DONE    = "DONE"
)

type Job struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type JobData struct {
	Name string `json:"name"`
	Data []byte `json:"data"`
}

func ListJobs(url string) ([]Job, error) {
	resp, err := http.Get(fmt.Sprintf("%s/list", url))
	if err != nil {
		return nil, errors.Wrap(err, "unable to perform request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("responseBody", string(body)).Info("response was")
		return nil, fmt.Errorf("statusCode was not OK, was: %d", resp.StatusCode)
	}

	jobs := []Job{}
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response")
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("responseBody", string(body)).WithField("status", resp.StatusCode).Error("response was not OK")
		return nil, fmt.Errorf("statuscode was not %d, was %d", http.StatusOK, resp.StatusCode)
	}

	return jobs, nil
}

func QueryJobs(url string, status string) ([]Job, error) {
	job := &List{
		Status: status,
	}

	queryBody, err := json.Marshal(job)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal query data")
	}

	resp, err := http.Post(fmt.Sprintf("%s/pending", url), "application/json", bytes.NewReader(queryBody))
	if err != nil {
		return nil, errors.Wrap(err, "unable to execute query")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("responseBody", string(body)).WithField("status", resp.StatusCode).Error("response was not OK")
		return nil, fmt.Errorf("statusCode was not OK, was: %d", resp.StatusCode)
	}

	jobs := []Job{}
	err = json.Unmarshal(body, &jobs)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response")
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("responseBody", string(body)).WithField("requestBody", string(queryBody)).Info("response was")
		return nil, fmt.Errorf("statuscode was not %d, was %d", http.StatusOK, resp.StatusCode)
	}

	return jobs, nil
}
