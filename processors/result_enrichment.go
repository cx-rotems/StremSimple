package processors

import(
	//"fmt"
	"time"
	"github.com/cx-rotems/StremSimple/types"
) 

type ResultEnrichment struct {

}

func NewResultEnrichment() *ResultEnrichment {
	return &ResultEnrichment{}
}

func (e *ResultEnrichment) Process(job types.Job) (types.Job, error) {
	for i := 0; i < len(job.Results); i++ {
		job.Results[i].CvssScores = job.Results[i].CvssScores + " enrichment"
		time.Sleep(60 * time.Millisecond) // simulate result enrichment
		//fmt.Printf("ResultEnrichment: Enriching result for result ID %d and job ID  %d\n",  job.Results[i].ResultID,  job.Results[i].JobID) 	// simulate result enrichment
	}

    return job, nil
}