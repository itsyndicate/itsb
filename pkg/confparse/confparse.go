package confparse

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Specs struct {
	Path             string `yaml:"path"`
	RdsArn           string `yaml:"rds_arn"`
	Extension        string `yaml:"extension"`
	ConnectionString string `yaml:"connection_string"`
	// rdsArn string
	// connectionString string
	// ...
}

type BackupJob struct {
	Id       string `yaml:"id"`
	Provider string `yaml:"provider"`
	Spec     Specs  `yaml:"spec"`
}

type Jobs struct {
	Jobs []BackupJob `yaml:"jobs"`
}

func StructTest() {
	f, err := os.ReadFile("examples/config.yaml")
	if err != nil {
		log.Fatal(err)
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
	fmt.Println("------------")
	fmt.Printf("jobs[1].id: %s\n", jobs.Jobs[1].Id)
	fmt.Printf("jobs[1].provdier: %s\n", jobs.Jobs[1].Provider)
	fmt.Printf("jobs[1].spec.rds_arn: %s\n", jobs.Jobs[1].Spec.RdsArn)
	fmt.Printf("jobs[1].spec.connection_string: %s\n", jobs.Jobs[1].Spec.ConnectionString)
	fmt.Printf("jobs[1].spec.path: %s\n", jobs.Jobs[1].Spec.Path)
	fmt.Printf("jobs[1].spec.extension: %s\n", jobs.Jobs[1].Spec.Extension)
}