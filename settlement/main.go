// Settlement demonstrates the core property: a loop over external calls where
// each call must happen exactly once, even if the process dies mid-loop.
//
//	alva dev &
//	go run ./settlement
package main

import (
	"errors"
	"fmt"
	"os"
)

type Txn struct {
	ID     string
	Amount int64
}

type Receipt struct {
	TxnID     string
	Reference string
}

func main() {
	if err := run("batch-2025-11-04"); err != nil {
		fmt.Fprintf(os.Stderr, "settlement failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("settled")
}

// run is the workflow body. In a real deployment this is registered with a
// worker; here it is called directly so the example stays one file.
func run(batchID string) error {
	txns, err := loadPending(batchID)
	if err != nil {
		return fmt.Errorf("load pending: %w", err)
	}

	for _, t := range txns {
		// The step id is derived from the transaction, not from a counter.
		// That is what makes it stable if the batch is reordered or the code
		// is edited — see the note in the SDK README about step identity.
		if _, err := post(t); err != nil {
			if errors.Is(err, errPermanent) {
				return fmt.Errorf("post %s: %w", t.ID, err)
			}
			return err
		}
	}

	return closeBatch(batchID)
}

var errPermanent = errors.New("permanent")

func loadPending(batchID string) ([]Txn, error) {
	return []Txn{
		{ID: "tx-1", Amount: 12_50},
		{ID: "tx-2", Amount: 4_00},
		{ID: "tx-3", Amount: 199_99},
	}, nil
}

func post(t Txn) (Receipt, error) {
	return Receipt{TxnID: t.ID, Reference: "ref-" + t.ID}, nil
}

func closeBatch(batchID string) error { return nil }
