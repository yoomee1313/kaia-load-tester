package tcutil

import (
	"math/big"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
)

// TestOperation defines a test case operation with Init and Run functions
type TestOperation struct {
	Init func(accs []*account.Account, endpoint string, gp *big.Int) *TcConfig
	Run  func(config *TcConfig)
	Name string
}

// TestOperations holds all available test operations
var TestOperations = map[string]*TestOperation{
	"valueTransfer": {
		Init: InitValueTransfer,
		Run:  RunValueTransfer,
		Name: "ValueTransfer",
	},
	"gasless": {
		Init: InitGasless,
		Run:  RunGasless,
		Name: "Gasless",
	},
	"gaslessRevert": {
		Init: InitGaslessRevert,
		Run:  RunGaslessRevert,
		Name: "GaslessRevert",
	},
	// Add more test operations here as needed
}

// GetTestOperation returns a test operation by name
func GetTestOperation(name string) (*TestOperation, bool) {
	op, exists := TestOperations[name]
	return op, exists
}

// ListTestOperations returns all available test operation names
func ListTestOperations() []string {
	names := make([]string, 0, len(TestOperations))
	for name := range TestOperations {
		names = append(names, name)
	}
	return names
}
