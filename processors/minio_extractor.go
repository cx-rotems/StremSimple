package processors

import (
	"time"

	"github.com/cx-rotems/StremSimple/types"
)

type MinioExtractor struct {
}

func NewMinioExtractor() *MinioExtractor {
	return &MinioExtractor{
	}
}

func (me *MinioExtractor) Process(job types.Job) (types.Job, error) {
	// Simulate extraction
	for i := 0; i < 50; i++ {
		time.Sleep(100 * time.Millisecond) // simulate download from Minio
		job.Results = append(job.Results, types.Result{ResultID: i, JobID: job.ID})
	}

	return  job, nil
}
