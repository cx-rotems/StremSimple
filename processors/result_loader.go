package processors

import (
	"fmt"
	"github.com/cx-rotems/StremSimple/types"
	"time"
)

const transactionSize = 4

type ResultLoader struct {
}

func NewResultLoader() *ResultLoader {
	return &ResultLoader{}
}

func (rl *ResultLoader) Process(job types.Job) (types.Job, error)  {

		//fmt.Printf("ResultLoader: Processing job ID %d\n", job.ID)

		// Process results in transactions
		transaction := make([]types.Result, 0, transactionSize)
		for _, result := range job.Results {
			transaction = append(transaction, result)

			if len(transaction) == transactionSize {
				processTransaction(transaction)
				transaction = transaction[:0]
			}
		}

		// Process remaining results if any
		if len(transaction) > 0 {
			processTransaction(transaction)
		}

		// Notify job completion
		return job, nil
	
}

var transactionCounter int

func processTransaction(transaction []types.Result) {
	transactionCounter++
	fmt.Printf("\nResultLoader: Saving transaction #%d (%d results)\n", transactionCounter, len(transaction))
	fmt.Println("Results in this transaction:")
	for i, result := range transaction {
		fmt.Printf("  [%d] Result ID: %d, Job ID: %d\n",
			i+1, result.ResultID, result.JobID)
	}
	time.Sleep(30 * time.Millisecond)
}