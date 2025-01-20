package processors

import (
	"github.com/cx-rotems/StremSimple/consts"
	"github.com/cx-rotems/StremSimple/types"
	//"fmt"
	"time"
)

type ResultEnrichment struct {
}

func NewResultEnrichment() *ResultEnrichment {
	return &ResultEnrichment{}
}

func (e *ResultEnrichment) Process(scanResults <-chan types.Result) <-chan types.Result {
	out := make(chan types.Result)
	go func() {
		defer close(out)
		for result := range scanResults {
			result.CvssScores = result.CvssScores + " enrichment"
			time.Sleep(consts.EngineRestructureTime * time.Millisecond) // simulate result enrichment
			//fmt.Printf("ResultEnrichment: Enriching result for result ID %d and job ID  %d\n",  job.Results[i].ResultID,  job.Results[i].JobID) 	// simulate result enrichment
			out <- result
		}
	}()
	return out
}
