package main

import (
	"fmt"
	"github.com/cx-rotems/StremSimple/consts"
	"github.com/cx-rotems/StremSimple/processors"
	"github.com/cx-rotems/StremSimple/types"
	"sync"
	"time"
)

var wg sync.WaitGroup
var start time.Time

func processJob(job types.Job) {
	wg.Add(1)
	go func(job types.Job) {
		defer wg.Done()
		results := processors.NewMinioExtractor().Process(job)
		restructureResults := processors.NewEngineResultsRestructure().Process(results)
		enrichedResults := processors.NewResultEnrichment().Process(restructureResults)
		doneResults, _ := processors.NewResultLoader().Process(enrichedResults)
		if job.ID == consts.NumberOfThreads {
			elapsed := time.Since(start)
			fmt.Printf("Total time took %s\n", elapsed)
		}
		fmt.Printf("Job %d completed with %d results\n", job.ID, doneResults)
	}(job)
}

func main() {
	start = time.Now()
	for i := 1; i <= consts.NumberOfThreads; i++ {
		processJob(types.Job{ID: i})
	}
	for {
	}
}
