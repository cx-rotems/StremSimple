package processors

import (
	"github.com/cx-rotems/StremSimple/consts"
	"github.com/cx-rotems/StremSimple/types"
	"math"
	//"fmt"
	"time"
)

type ResultLoader struct {
}

func NewResultLoader() *ResultLoader {
	return &ResultLoader{}
}

func (rl *ResultLoader) Process(scanResults <-chan types.Result) (int, error) {

	//fmt.Printf("ResultLoader: Processing job ID %d\n", job.ID)
	loaded := 0
	transaction := make([]types.Result, 0, consts.TransactionSize)

	for result := range scanResults {
		transaction = append(transaction, result)

		if len(transaction) == consts.TransactionSize {
			processTransaction(transaction)
			transaction = transaction[:0]
		}
		loaded++
	}

	// Process remaining results if any
	if len(transaction) > 0 {
		processTransaction(transaction)
	}

	// Notify job completion
	return loaded, nil

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
