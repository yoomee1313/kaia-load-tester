package account

import (
	"math/big"
	"sync"

	"github.com/kaiachain/kaia/common"
)

// ERC721TokenLedger is a thread-safe token ledger for ERC721 tokens
type ERC721TokenLedger struct {
	// Account-based token queues using sync.Map for thread-safety
	tokenQueues sync.Map // map[common.Address][]*big.Int

	// Account-specific mutexes for fine-grained locking
	accountMutexes sync.Map // map[common.Address]*sync.Mutex
}

// NewERC721TokenLedger creates a new thread-safe token ledger
func NewERC721TokenLedger() *ERC721TokenLedger {
	return &ERC721TokenLedger{}
}

// getAccountMutex returns or creates a mutex for a specific account
func (l *ERC721TokenLedger) getAccountMutex(addr common.Address) *sync.Mutex {
	// Try to get existing mutex
	if value, exists := l.accountMutexes.Load(addr); exists {
		return value.(*sync.Mutex)
	}

	// Create new mutex if it doesn't exist
	mutex := &sync.Mutex{}

	// Store the new mutex (LoadOrStore ensures thread-safety)
	if actual, loaded := l.accountMutexes.LoadOrStore(addr, mutex); loaded {
		// Another goroutine created the mutex first
		return actual.(*sync.Mutex)
	}

	return mutex
}

// PutToken safely adds a token to the ledger
func (l *ERC721TokenLedger) PutToken(addr common.Address, tokenId *big.Int) {
	mutex := l.getAccountMutex(addr)
	mutex.Lock()
	defer mutex.Unlock()

	value, exists := l.tokenQueues.Load(addr)

	var tokens []*big.Int
	if exists {
		tokens = value.([]*big.Int)
	} else {
		tokens = make([]*big.Int, 0)
	}

	// Add token to queue
	tokens = append(tokens, tokenId)

	// Store atomically
	l.tokenQueues.Store(addr, tokens)
}

// RemoveToken safely removes a token from the ledger (FIFO)
func (l *ERC721TokenLedger) RemoveToken(addr common.Address) *big.Int {
	mutex := l.getAccountMutex(addr)
	mutex.Lock()
	defer mutex.Unlock()

	value, exists := l.tokenQueues.Load(addr)
	if !exists {
		return nil
	}

	tokens := value.([]*big.Int)
	if len(tokens) == 0 {
		return nil
	}

	// Get the first token (FIFO queue)
	token := tokens[0]
	newTokens := tokens[1:] // Remove the first token

	// Update the queue atomically
	l.tokenQueues.Store(addr, newTokens)

	return token
}

// InitializeAccount initializes an account with empty token queue
func (l *ERC721TokenLedger) InitializeAccount(addr common.Address) {
	mutex := l.getAccountMutex(addr)
	mutex.Lock()
	defer mutex.Unlock()

	// Check if already initialized
	if _, exists := l.tokenQueues.Load(addr); !exists {
		l.tokenQueues.Store(addr, make([]*big.Int, 0))
	}
}

// Global instance
var ERC721Ledger = NewERC721TokenLedger()
