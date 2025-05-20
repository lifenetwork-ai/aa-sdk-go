package aasdk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"sync"
	"testing"
)

// Helper function to generate random private keys for testing
func generatePrivateKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	return privateKey
}

func TestNewRoundRobinSignerProvider(t *testing.T) {
	// Test with empty signers
	provider := NewRoundRobinSignerProvider(nil)
	if provider == nil {
		t.Fatal("Expected non-nil provider")
	}
	if provider.Count() != 0 {
		t.Errorf("Expected count 0, got %d", provider.Count())
	}

	// Test with some signers
	key1 := generatePrivateKey(t)
	key2 := generatePrivateKey(t)
	signers := []*ecdsa.PrivateKey{key1, key2}

	provider = NewRoundRobinSignerProvider(signers)
	if provider.Count() != 2 {
		t.Errorf("Expected count 2, got %d", provider.Count())
	}
}

func TestNext(t *testing.T) {
	// Create test keys
	key1 := generatePrivateKey(t)
	key2 := generatePrivateKey(t)
	key3 := generatePrivateKey(t)

	// Test rotation behavior
	provider := NewRoundRobinSignerProvider([]*ecdsa.PrivateKey{key1, key2, key3})

	// Should rotate through all keys in order
	for i := 0; i < 6; i++ {
		expectedKey := []*ecdsa.PrivateKey{key1, key2, key3}[i%3]
		gotKey := provider.Next()
		if gotKey != expectedKey {
			t.Errorf("Rotation cycle %d: Expected key %v, got %v", i, expectedKey, gotKey)
		}
	}
}

func TestNextWithEmptyProvider(t *testing.T) {
	provider := NewRoundRobinSignerProvider(nil)

	// Should return nil when no signers
	if signer := provider.Next(); signer != nil {
		t.Errorf("Expected nil signer for empty provider, got %v", signer)
	}
}

func TestAdd(t *testing.T) {
	provider := NewRoundRobinSignerProvider(nil)

	// Add signers and verify count
	key1 := generatePrivateKey(t)
	err := provider.Add(key1)
	if err != nil {
		t.Errorf("Unexpected error when adding signer: %v", err)
	}
	if provider.Count() != 1 {
		t.Errorf("Expected count 1 after adding a signer, got %d", provider.Count())
	}

	// Verify the added signer is returned by Next
	if signer := provider.Next(); signer != key1 {
		t.Errorf("Expected signer %v, got %v", key1, signer)
	}

	// Add another signer and check rotation
	key2 := generatePrivateKey(t)
	err = provider.Add(key2)
	if err != nil {
		t.Errorf("Unexpected error when adding signer: %v", err)
	}

	// Should now rotate between key1 and key2
	if signer := provider.Next(); signer != key1 {
		t.Errorf("Expected first signer %v, got %v", key1, signer)
	}

	// Should rotate to key2
	if signer := provider.Next(); signer != key2 {
		t.Errorf("Expected second signer %v, got %v", key2, signer)
	}
}

func TestConcurrentAccess(t *testing.T) {
	provider := NewRoundRobinSignerProvider(nil)

	// Add some initial signers
	for i := 0; i < 3; i++ {
		err := provider.Add(generatePrivateKey(t))
		if err != nil {
			t.Fatalf("Error adding initial signer: %v", err)
		}
	}

	// Test concurrent access with multiple goroutines
	const numGoroutines = 10
	const iterationsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2) // Half for Add, half for Next

	// Launch goroutines that call Next
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterationsPerGoroutine; j++ {
				signer := provider.Next()
				if signer == nil {
					t.Errorf("Got nil signer during concurrent access")
				}
			}
		}()
	}

	// Launch goroutines that call Add
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < iterationsPerGoroutine/10; j++ { // Fewer adds than Next calls
				err := provider.Add(generatePrivateKey(t))
				if err != nil {
					t.Errorf("Error adding signer: %v", err)
				}
			}
		}()
	}

	wg.Wait()

	// Final count should be initial 3 plus 10 goroutines * (iterations/10) additions
	expectedCount := 3 + numGoroutines*(iterationsPerGoroutine/10)
	if count := provider.Count(); count != expectedCount {
		t.Errorf("Expected final count %d, got %d", expectedCount, count)
	}
}

func TestIndex(t *testing.T) {
	// Create a provider with several signers
	keys := make([]*ecdsa.PrivateKey, 5)
	for i := range keys {
		keys[i] = generatePrivateKey(t)
	}

	provider := NewRoundRobinSignerProvider(keys)

	// Track which keys are returned and how many times
	counts := make(map[*ecdsa.PrivateKey]int)

	// Call Next many times
	const iterations = 100
	for i := 0; i < iterations; i++ {
		key := provider.Next()
		counts[key]++
	}

	// Each key should be returned approximately the same number of times
	expectedCount := iterations / len(keys)
	for key, count := range counts {
		// Allow for some small variation
		if count < expectedCount-1 || count > expectedCount+1 {
			t.Errorf("Key %v: expected approximately %d calls, got %d", key, expectedCount, count)
		}
	}
}

func BenchmarkNext(b *testing.B) {
	// Create a provider with 10 signers
	keys := make([]*ecdsa.PrivateKey, 10)
	for i := range keys {
		var err error
		keys[i], err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			b.Fatalf("Failed to generate key: %v", err)
		}
	}

	provider := NewRoundRobinSignerProvider(keys)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Next()
	}
}

func BenchmarkConcurrentNext(b *testing.B) {
	// Create a provider with 10 signers
	keys := make([]*ecdsa.PrivateKey, 10)
	for i := range keys {
		var err error
		keys[i], err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			b.Fatalf("Failed to generate key: %v", err)
		}
	}

	provider := NewRoundRobinSignerProvider(keys)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			provider.Next()
		}
	})
}
