package processors

import(
	"time"
	"fmt"
	"github.com/cx-rotems/StremSimple/types"
) 

// EngineResultsRestructure is a struct that represents the processor.
type EngineResultsRestructure struct{}

// NewEngineResultsRestructure creates a new instance of EngineResultsRestructure.
func NewEngineResultsRestructure() *EngineResultsRestructure {
    return &EngineResultsRestructure{}
}

// Process processes the job and returns the modified job.
func (e *EngineResultsRestructure) Process(job types.Job) (types.Job, error) {
	for i := 0; i < len(job.Results); i++ {

		job.Results[i].CvssScores = fmt.Sprintf("%d", i*10)
		time.Sleep(70 * time.Millisecond) // simulate restructure
			//fmt.Printf("EngineResultsRestructure: Restructuring result for result ID %d and job ID  %d\n", job.Results[i].ResultID, job.Results[i].JobID)
	}

    return job, nil
}