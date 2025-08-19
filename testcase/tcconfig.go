package testcase

import (
	"log"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/kaiachain/kaia/networks/rpc"
)

// TCConfig holds common configuration for test cases
type TCConfig struct {
	Name                    string
	EndPoint                string
	NAcc                    int
	AccGrp                  []*account.Account
	CliPool                 clipool.ClientPool
	RpcCliPool              clipool.ClientPool                        // For RPC client (used by specific Read API test cases)
	SmartContractAccounts   map[account.TestContract]*account.Account // For multiple contracts
	TestContracts           []account.TestContract
	AuctionTargetTxTypeList []string // For auction test cases
}

// Init initializes common configuration for test cases
func Init(accGrp *account.AccGroup, endpoint string, testContracts []account.TestContract, tcName string, auctionTargetTxTypeList []string) *TCConfig {
	config := &TCConfig{
		Name:                  tcName,
		EndPoint:              endpoint,
		SmartContractAccounts: make(map[account.TestContract]*account.Account),
		TestContracts:         testContracts,
	}

	cliCreate := func() interface{} {
		c, err := client.Dial(config.EndPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	config.CliPool.Init(20, 300, cliCreate)

	// Read API test cases use rpc.Client
	rpcCliCreate := func() interface{} {
		c, err := rpc.Dial(config.EndPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}
	config.RpcCliPool.Init(20, 300, rpcCliCreate)

	// Get accounts from accGrp
	accs := accGrp.GetAccListByName(account.AccListForSignedTx)
	if tcName == "transferUnsignedTx" {
		accs = accGrp.GetAccListByName(account.AccListForUnsignedTx)
	} else if tcName == "gaslessRevertTransactionTC" {
		accs = accGrp.GetAccListByName(account.AccListForGaslessRevertTx)
	} else if tcName == "gaslessOnlyApproveTC" {
		accs = accGrp.GetAccListByName(account.AccListForGaslessApproveTx)
	}
	for _, acc := range accs {
		config.AccGrp = append(config.AccGrp, acc)
	}

	config.NAcc = len(config.AccGrp)

	// Set SmartContractAccounts if a specific contract is required
	contractsParam := accGrp.GetTestContractList()
	for _, contractType := range testContracts {
		config.SmartContractAccounts[contractType] = contractsParam[contractType]
	}

	// Set TargetTxTypeList for auction test cases
	if tcName == "auctionBidTC" || tcName == "auctionRevertedBidTC" {
		config.AuctionTargetTxTypeList = auctionTargetTxTypeList
	}

	// Initialize ethereum specific variables for ethereum test cases
	if tcName == "ethereumTxLegacyTC" || tcName == "ethereumTxAccessListTC" || tcName == "ethereumTxDynamicFeeTC" ||
		tcName == "newEthereumAccessListTC" || tcName == "newEthereumDynamicFeeTC" {
		initEthereum()
	}

	return config
}
