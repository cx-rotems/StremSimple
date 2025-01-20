package processors

import (
	"github.com/cx-rotems/StremSimple/consts"
	"time"
	//	"fmt"
	"github.com/cx-rotems/StremSimple/types"
)

type MinioExtractor struct {
}

func NewMinioExtractor() *MinioExtractor {
	return &MinioExtractor{}
}

func (me *MinioExtractor) Process(job types.Job) <-chan types.Result {
	// Simulate extraction
	out := make(chan types.Result)
	go func() {
		defer close(out)
		files := getFiles(job)
		resultID := 1
		for range files {
			for i := 1; i <= consts.NumberOfResultPerFile; i++ {
				time.Sleep(consts.MinioExtractFileTime * time.Millisecond) // simulate parse each result
				result := types.Result{ResultID: resultID, JobID: job.ID}
				job.Results = append(job.Results, result)
				resultID++
				out <- result
				//	fmt.Printf("MinioExtractor: Downlaod result ID %d and job ID  %d\n",  job.Results[i].ResultID,  job.Results[i].JobID)
			}
		}
	}()

	return out
}

func getFiles(job types.Job) int {
	time.Sleep(consts.MinioExtractTime * time.Millisecond) // simulate download from Minio
	return consts.NumberOfFiles
}
