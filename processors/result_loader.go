package processors

import (
	//"fmt"
	"time"
	"math"
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
	resultLoaderTransactionTime := int(math.Ceil(float64(len(transaction)) / 4.0))
	
	transactionCounter++
	//fmt.Printf("\nResultLoader: Saving transaction #%d (%d results)\n", transactionCounter, len(transaction))
	//fmt.Println("Results in this transaction:")
	for _, result := range transaction {
		result = result

	}
	time.Sleep(time.Duration(resultLoaderTransactionTime) * time.Millisecond)
}