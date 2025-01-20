package processors

import (
	"fmt"
	"github.com/cx-rotems/StremSimple/consts"
	"github.com/cx-rotems/StremSimple/types"
	"time"
)

// EngineResultsRestructure is a struct that represents the processor.
type EngineResultsRestructure struct{}

// NewEngineResultsRestructure creates a new instance of EngineResultsRestructure.
func NewEngineResultsRestructure() *EngineResultsRestructure {
	return &EngineResultsRestructure{}
}

// Process processes the job and returns the modified job.
func (e *EngineResultsRestructure) Process(scanResults <-chan types.Result) <-chan types.Result {

	out := make(chan types.Result)
	go func() {
		defer close(out)
		var i = 1
		for result := range scanResults {
			result.CvssScores = fmt.Sprintf("%d", i*10)
			i++
			time.Sleep(consts.EngineRestructureTime * time.Millisecond) // simulate restructure
			//fmt.Printf("EngineResultsRestructure: Restructuring result for result ID %d and job ID  %d\n", job.Results[i].ResultID, job.Results[i].JobID)
			out <- result
		}
	}()
	return out
}
