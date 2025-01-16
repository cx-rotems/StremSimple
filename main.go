package main

import (
    "fmt"
    "sync"
    "github.com/cx-rotems/StremSimple/types"
    "github.com/cx-rotems/StremSimple/processors"
    "time"
)

var wg sync.WaitGroup
var start time.Time

func processJob(job types.Job) {
    wg.Add(1)
    go func(job types.Job) {
        defer wg.Done()
        jobWithResult, _ := processors.NewMinioExtractor().Process(job)
        jobWithResult, _ = processors.NewEngineResultsRestructure().Process(jobWithResult)
        jobWithResult, _ = processors.NewResultEnrichment().Process(jobWithResult)
        jobWithResult, _ = processors.NewResultLoader().Process(jobWithResult)
        if (jobWithResult.ID == 3) {
            elapsed := time.Since(start)
			fmt.Printf("Total time took %s\n", elapsed)
        }
        fmt.Printf("Job %d completed with %d results\n", job.ID, len(jobWithResult.Results))
    }(job)
}

func main() {
    start = time.Now()
    for i := 1; i <= 3; i++ {
        processJob(types.Job{ID: i})
    }
    for { }
}
