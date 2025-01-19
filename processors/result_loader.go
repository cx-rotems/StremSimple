package processors

import (
	"fmt"
	"time"
	"github.com/cx-rotems/StremSimple/types"
	"github.com/cx-rotems/StremSimple/consts"
)

type ResultLoader struct {
}

func NewResultLoader() *ResultLoader {
	return &ResultLoader{}
}

func (rl *ResultLoader) Process(job types.Job) (types.Job, error)  {

		//fmt.Printf("ResultLoader: Processing job ID %d\n", job.ID)
		transaction := make([]types.Result, 0, consts.TransactionSize)
		for _, result := range job.Results {
			transaction = append(transaction, result)

			if len(transaction) == consts.TransactionSize {
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