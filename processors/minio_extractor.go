package processors

import (
	"time"
//	"fmt"
	"github.com/cx-rotems/StremSimple/types"
	"github.com/cx-rotems/StremSimple/consts"
)

type MinioExtractor struct {
}

func NewMinioExtractor() *MinioExtractor {
	return &MinioExtractor{
	}
}

func (me *MinioExtractor) Process(job types.Job) (types.Job, error) {
	// Simulate extraction
	for i := 1; i <= consts.NumberOfResult; i++ {
		time.Sleep(consts.MinioExtractTime * time.Millisecond) // simulate download from Minio
		job.Results = append(job.Results, types.Result{ResultID: i, JobID: job.ID})
	//	fmt.Printf("MinioExtractor: Downlaod result ID %d and job ID  %d\n",  job.Results[i].ResultID,  job.Results[i].JobID) 
	}

	return  job, nil
}
