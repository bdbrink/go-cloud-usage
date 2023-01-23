package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	log "github.com/sirupsen/logrus"
)

func generateReportGET() {

	sess := session.Must(session.NewSession())
	costClient := costexplorer.New(sess)
	now := time.Now()
	formatDate := now.Format("2006-01-01")
	sixMonths := time.Now().AddDate(0, -6, 0)
	sixMonthsFormat := sixMonths.Format("2006-01-01")

	fmt.Println(formatDate)
	fmt.Println(sixMonthsFormat)

	result, err := costClient.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		Filter: &costexplorer.Expression{
			Dimensions: &costexplorer.DimensionValues{
				Key:    aws.String("SERVICE"),
				Values: aws.StringSlice([]string{"Amazon Simple Storage Service"}),
			},
		},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(sixMonthsFormat),
			End:   aws.String(formatDate),
		},
		Granularity: aws.String("MONTHLY"),
		GroupBy: []*costexplorer.GroupDefinition{
			&costexplorer.GroupDefinition{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("OPERATION"),
			},
		},
		Metrics: aws.StringSlice([]string{"USAGE_QUANTITY"}),
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}

func main() {

	generateReportGET()
}
