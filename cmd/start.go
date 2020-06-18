package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/leicmi/cloud-computing/util"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

func start(url string, filename string) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.WithError(err).Errorf("failed to read file %q", filename)
		return
	}

	result, err := invokeUpload(url, filename, f)
	if err != nil {
		logrus.WithError(err).Errorf("unable to upload file")
		return
	}

	fmt.Printf("Queued job %q\n", result)
}

// func invokeUpload(sess *session.Session, filename string, data []byte) (*lambda.InvokeOutput, error) {
func invokeUpload(url string, filename string, data []byte) (string, error) {
	job := &util.JobData{
		Name: filename,
		Data: data,
	}

	jobBody, err := json.Marshal(job)
	if err != nil {
		return "", errors.Wrap(err, "unable to marshal job data")
	}

	resp, err := http.Post(fmt.Sprintf("%s/default/upload", url), "application/json", bytes.NewReader(jobBody))
	if err != nil {
		return "", errors.Wrap(err, "unable to upload job")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "unable to read response body")
	}

	if resp.StatusCode != http.StatusOK {
		logrus.WithField("responseBody", string(body)).WithField("requestBody", string(jobBody)).WithField("status", resp.StatusCode).Error("response was not OK")
		return "", fmt.Errorf("statuscode was not %d, was %d", http.StatusOK, resp.StatusCode)
	}

	return string(body), nil
}
