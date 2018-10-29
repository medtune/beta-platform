package custom

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/medtune/beta-platform/internal/jsonutil"
)

var mutex sync.Mutex

func run(job *jsonutil.Job, jobResult *jsonutil.JobResult) {
	fnIdentifiers := strings.Split(job.Function, "/")

	service := fnIdentifiers[0]
	demo := fnIdentifiers[1]
	fn := fnIdentifiers[2]

	if service != "demo" {
		jobResult.Errors = append(jobResult.Errors, fmt.Sprintf("unkown service: %s", service))
		return
	}
	if demo != "mura" && demo != "inception" && demo != "chexray" {
		jobResult.Errors = append(jobResult.Errors, fmt.Sprintf("unkown demo: %s", demo))
		return
	}
	if fn != "run_cam" && demo != "run_inference" {
		jobResult.Errors = append(jobResult.Errors, fmt.Sprintf("unkown method: %s", fn))
		return
	}

}

// simple job execution
func executeJob(job *jsonutil.Job) *jsonutil.JobResult {
	jobResult := &jsonutil.JobResult{}
	start := time.Now()
	defer func() {
		jobResult.RoundTrip = time.Since(start)
	}()

	run(job, jobResult)
	return jobResult
}

// execute sequentially a list of jobs
func sequentialExecution(jobs []*jsonutil.Job) []*jsonutil.JobResult {
	return nil
}

// Execute concurrently list of jobs
func concurrentExecution(jobs []*jsonutil.Job) []*jsonutil.JobResult {
	var wg sync.WaitGroup
	var jobResults = make([]*jsonutil.JobResult, 0, len(jobs))
	for _, job := range jobs {
		wg.Add(1)
		go func(j *jsonutil.Job) {
			jobRes := executeJob(j)
			jobResults = append(jobResults, jobRes)
			wg.Done()
		}(job)
	}
	wg.Wait()
	return jobResults
}

// Execute will run the configuration job
// Used mainly to run parallel inferences/cams
func Execute(cx *jsonutil.CustomExecutionRequest) ([]*jsonutil.JobResult, error) {

	var jobResults []*jsonutil.JobResult
	if cx.Concurrency {
		jobResults = concurrentExecution(cx.Jobs)
	} else {
		jobResults = sequentialExecution(cx.Jobs)
	}

	return jobResults, nil
}
