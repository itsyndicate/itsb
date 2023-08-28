package confparse

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ParseConfig(configPath string) Jobs {
	var jobs Jobs

	f, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(f, &jobs); err != nil {
		log.Fatalln(err)
	}

	return jobs
}

func StructTest() {
	f, err := os.ReadFile("examples/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	var jobs Jobs

	if err := yaml.Unmarshal(f, &jobs); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", jobs)
	fmt.Println("------------")
	fmt.Printf("jobs[0].id: %s\n", jobs.Jobs[0].Id)
	fmt.Printf("jobs[0].provdier: %s\n", jobs.Jobs[0].Provider)
	fmt.Printf("jobs[0].spec.rds_arn: %s\n", jobs.Jobs[0].Spec.RdsArn)
	fmt.Printf("jobs[0].spec.connection_string: %s\n", jobs.Jobs[0].Spec.ConnectionString)
	fmt.Printf("jobs[0].spec.path: %s\n", jobs.Jobs[0].Spec.Path)
	fmt.Printf("jobs[0].spec.extension: %s\n", jobs.Jobs[0].Spec.Extension)
	fmt.Printf("jobs[0].targets[0].type: %s\n", jobs.Jobs[0].Targets[0].Type)
	fmt.Printf("jobs[0].targets[0].target_spec.s3_bucket_name: %s\n", jobs.Jobs[0].Targets[0].TargetSpec.S3BucketName)
	fmt.Printf("jobs[0].targets[0].target_spec.s3_bucket_key: %s\n", jobs.Jobs[0].Targets[0].TargetSpec.S3BucketKey)
	fmt.Println("------------")
	fmt.Printf("jobs[1].id: %s\n", jobs.Jobs[1].Id)
	fmt.Printf("jobs[1].provdier: %s\n", jobs.Jobs[1].Provider)
	fmt.Printf("jobs[1].spec.rds_arn: %s\n", jobs.Jobs[1].Spec.RdsArn)
	fmt.Printf("jobs[1].spec.connection_string: %s\n", jobs.Jobs[1].Spec.ConnectionString)
	fmt.Printf("jobs[1].spec.path: %s\n", jobs.Jobs[1].Spec.Path)
	fmt.Printf("jobs[1].spec.extension: %s\n", jobs.Jobs[1].Spec.Extension)
	fmt.Printf("jobs[1].targets[0].type: %s\n", jobs.Jobs[1].Targets[0].Type)
	fmt.Printf("jobs[1].targets[0].target_spec.s3_bucket_name: %s\n", jobs.Jobs[1].Targets[0].TargetSpec.S3BucketName)
	fmt.Printf("jobs[1].targets[0].target_spec.s3_bucket_key: %s\n", jobs.Jobs[1].Targets[0].TargetSpec.S3BucketKey)
}
