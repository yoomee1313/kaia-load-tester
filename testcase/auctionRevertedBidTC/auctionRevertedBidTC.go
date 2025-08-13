package auctionRevertedBidTC

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "auctionRevertedBidTC"

var (
	endPoint string
	accGrp   *account.AccountSet
	cliPool  clipool.ClientPool

	AuctionEntryPointAccount     *account.Account
	CounterForTestAuctionAccount *account.Account

	TargetTxTypeList []string
)

func Init(accs []*account.Account, contractsParam []*account.Account, endpoint string, gp *big.Int) {
	endPoint = endpoint
	AuctionEntryPointAccount = contractsParam[account.ContractAuctionEntryPoint]
	CounterForTestAuctionAccount = contractsParam[account.ContractCounterForTestAuction]

	cliCreate := func() interface{} {
		c, err := client.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	cliPool.Init(20, 300, cliCreate)

	accGrp = account.NewAccountSet(accs)
}

func Run() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	from := accGrp.GetAccountRoundRobin()
	testRecordName := "AuctionRevertedBid" + " to " + endPoint

	// Select a targetTxType randomly from the list.
	targetTxTypeKey := TargetTxTypeList[rand.Int()%len(TargetTxTypeList)]

	start := boomer.Now()

	_, _, _, err := from.AuctionRevertedBid(cli, endPoint, AuctionEntryPointAccount, CounterForTestAuctionAccount, targetTxTypeKey)

	elapsed := boomer.Now() - start

	if err != nil {
		boomer.RecordFailure("http", testRecordName, elapsed, err.Error())
	} else {
		boomer.RecordSuccess("http", testRecordName, elapsed, int64(10))
	}
}
