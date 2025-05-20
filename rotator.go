package aasdk

import (
	"crypto/ecdsa"
	"sync"
	"sync/atomic"
)

type Rotator[T any] interface {
	// Next returns the next available signer.
	Next() T

	// Add adds a new signer to the rotation.
	Add(signer T) error

	// Count returns the number of signers available.
	Count() int
}

type RoundRobinSignerProvider struct {
	signers []*ecdsa.PrivateKey
	index   atomic.Uint32
	mu      sync.RWMutex
}

var _ Rotator[*ecdsa.PrivateKey] = (*RoundRobinSignerProvider)(nil)

func NewRoundRobinSignerProvider(signers []*ecdsa.PrivateKey) Rotator[*ecdsa.PrivateKey] {
	return &RoundRobinSignerProvider{
		signers: signers,
		index:   atomic.Uint32{},
	}
}

func (p *RoundRobinSignerProvider) Next() *ecdsa.PrivateKey {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.signers) == 0 {
		return nil
	}
	current := p.index.Load()
	p.index.Store((current + 1) % uint32(len(p.signers)))
	return p.signers[current]
}

func (r *RoundRobinSignerProvider) Add(signer *ecdsa.PrivateKey) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.signers = append(r.signers, signer)
	return nil
}

func (r *RoundRobinSignerProvider) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return len(r.signers)
}
