package main

import (
	"fmt"
	"jobapp/job"
)

func main() {
	var job = job.Job{"Software Engineer", "San Fran", true}
	fmt.Println(job)
	job.SetJobName("Software Tester")
	fmt.Println(job)
}

