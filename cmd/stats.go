package cmd

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	log "github.com/sirupsen/logrus"
)

func stats(accesskeyid string, secretkey string, logGroupName string) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		log.WithField("err", err).Fatal("unable to auth against aws")
	}
	svc := cloudwatchlogs.New(sess)

	sqi := &cloudwatchlogs.StartQueryInput{
		StartTime:    aws.Int64(time.Now().Add(time.Hour * -1).Unix()),
		EndTime:      aws.Int64(time.Now().Unix()),
		LogGroupName: aws.String(logGroupName),
		QueryString:  aws.String("fields @timestamp, @message, @billedDuration, @duration, @maxMemoryUsed, @memorySize | filter @type='REPORT'"),
	}
	sqo, err := svc.StartQuery(sqi)
	if err != nil {
		log.WithField("err", err).Fatal("unable to start insights query")
	}

	time.Sleep(time.Second * 5)

	gqri := &cloudwatchlogs.GetQueryResultsInput{QueryId: sqo.QueryId}
	req, resp := svc.GetQueryResultsRequest(gqri)
	for {
		if err := req.Send(); err == nil {
			break
		} else {
			log.Warn("query not completed, retying")
		}
	}
	if err != nil {
		log.WithField("err", err).Fatal("error fetching insights query result")
	}
	//log.Info(resp.String())
	fmt.Println("duration, billedDuration, memorySize, maxMemoryUsed")
	for _, r := range resp.Results {
		var m map[string]string
		m = make(map[string]string)
		for _, tuple := range r {
			m[*tuple.Field] = *tuple.Value
		}
		fmt.Printf("%s, %s, %s, %s\n", m["@duration"], m["@billedDuration"], m["@memorySize"], m["@maxMemoryUsed"])
	}
}
