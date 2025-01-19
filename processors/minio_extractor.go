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
	files := getFiles(job)

	resultID := 1;
	 for range files{
		for i := 1; i <= consts.NumberOfResultPerFile; i++ {
			time.Sleep(consts.MinioExtractFileTime * time.Millisecond) // simulate parse each result
			job.Results = append(job.Results, types.Result{ResultID: resultID, JobID: job.ID})
			resultID++;
		//	fmt.Printf("MinioExtractor: Downlaod result ID %d and job ID  %d\n",  job.Results[i].ResultID,  job.Results[i].JobID) 
		}
	}

	return  job, nil
}

func getFiles(job types.Job) int {
	time.Sleep(consts.MinioExtractTime * time.Millisecond) // simulate download from Minio
	return consts.NumberOfFiles
}