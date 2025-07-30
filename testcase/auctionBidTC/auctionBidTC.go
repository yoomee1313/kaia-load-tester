package auctionBidTC

import (
	"log"
	"math/big"
	"math/rand"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "auctionBidTC"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool

	AuctionEntryPointAccount     *account.Account
	CounterForTestAuctionAccount *account.Account

	TargetTxTypeList []string
)

func Init(accs []*account.Account, endpoint string, gp *big.Int) {
	endPoint = endpoint

	cliCreate := func() interface{} {
		c, err := client.Dial(endPoint)
		if err != nil {
			log.Fatalf("Failed to connect RPC: %v", err)
		}
		return c
	}

	cliPool.Init(20, 300, cliCreate)

	for _, acc := range accs {
		accGrp = append(accGrp, acc)
	}

	nAcc = len(accGrp)
}

func Run() {
	cli := cliPool.Alloc().(*client.Client)

	from := accGrp[rand.Int()%nAcc]

	testRecordName := "AuctionBid" + " to " + endPoint

	// Select a targetTxType randomly from the list.
	targetTxTypeKey := TargetTxTypeList[rand.Int()%len(TargetTxTypeList)]

	start := boomer.Now()

	_, _, _, err := from.AuctionBid(cli, endPoint, AuctionEntryPointAccount, CounterForTestAuctionAccount, targetTxTypeKey)

	elapsed := boomer.Now() - start

	cliPool.Free(cli)

	if err != nil {
		boomer.RecordFailure("http", testRecordName, elapsed, err.Error())
	} else {
		boomer.RecordSuccess("http", testRecordName, elapsed, int64(10))
	}
}
