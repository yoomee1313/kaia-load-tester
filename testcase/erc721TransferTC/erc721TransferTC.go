package erc721TransferTC

import (
	"log"
	"math/big"
	"math/rand"
	"time"

	"github.com/kaiachain/kaia-load-tester/klayslave/account"
	"github.com/kaiachain/kaia-load-tester/klayslave/clipool"
	"github.com/kaiachain/kaia/client"
	"github.com/myzhan/boomer"
)

const Name = "erc721TransferTC"

var (
	endPoint string
	nAcc     int
	accGrp   []*account.Account
	cliPool  clipool.ClientPool
	gasPrice *big.Int

	// multinode tester
	transferedValue *big.Int
	expectedFee     *big.Int

	fromAccount     *account.Account
	prevBalanceFrom *big.Int

	toAccount     *account.Account
	prevBalanceTo *big.Int

	SmartContractAccount *account.Account
)

func Init(accs []*account.Account, endpoint string, gp *big.Int) {
	gasPrice = gp

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

	rand.Seed(time.Now().UnixNano())
}

func Run() {
	cli := cliPool.Alloc().(*client.Client)
	defer cliPool.Free(cli)

	toAcc := accGrp[rand.Intn(nAcc)]

	// Find an account with available tokens
	var fromAcc *account.Account
	var tokenId *big.Int

	// Try multiple accounts to find one with tokens
	candidateIdx := rand.Intn(len(accGrp))

	// limit the number of attempts to find a token
	for i := 0; i < 200; i++ {
		candidateAcc := accGrp[candidateIdx]
		tokenId = account.ERC721Ledger.RemoveToken(candidateAcc.GetAddress())
		if tokenId != nil {
			fromAcc = candidateAcc
			break
		}
		candidateIdx = (candidateIdx + 1) % len(accGrp)
	}

	if tokenId == nil {
		// No tokens available in any account
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, int64(0), "No tokens available")
		return
	}

	start := boomer.Now()
	_, _, err := fromAcc.TransferERC721(false, cli, SmartContractAccount.GetAddress(), toAcc, tokenId)
	elapsed := boomer.Now() - start

	if err == nil {
		boomer.Events.Publish("request_success", "http", Name+" to "+endPoint, elapsed, int64(10))
		// Transfer successful, add token to destination account
		account.ERC721Ledger.PutToken(toAcc.GetAddress(), tokenId)
	} else {
		boomer.Events.Publish("request_failure", "http", Name+" to "+endPoint, elapsed, err.Error())
		// Transfer failed, put token back to original owner
		account.ERC721Ledger.PutToken(fromAcc.GetAddress(), tokenId)
	}
}
