package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
    "github.com/cx-rotems/StremSimple/types"
	"github.com/cx-rotems/StremSimple/processors"
)

var wg sync.WaitGroup

func processJob(job types.Job) {
    wg.Add(1)
    go func(job types.Job) {
        defer wg.Done()
        jobWithResult1, _ := processors.NewMinioExtractor().Process(job)
        jobWithResult2, _ := processors.NewEngineResultsRestructure().Process(jobWithResult1)
        jobWithResult3, _ := processors.NewResultEnrichment().Process(jobWithResult2)
        jobWithResult4, _ := processors.NewMinioExtractor().Process(jobWithResult3)
        fmt.Printf("Job %d completed with %d result\n", job.ID, len(jobWithResult4.Results))
    }(job)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    jobID := 1

    for {
        processJob(types.Job{ID: jobID})
        jobID++

        // Sleep for a random duration between 1 and 5 seconds
        time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
    }
}