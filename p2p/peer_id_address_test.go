package p2p

import "testing"

func TestAddressContainsPeerID(t *testing.T) {
	peerID := "16Uiu2HAmExactPeer"

	tests := []struct {
		name     string
		address  string
		expected bool
	}{
		{
			name:     "bare peer id",
			address:  peerID,
			expected: true,
		},
		{
			name:     "p2p segment",
			address:  "/ip4/127.0.0.1/tcp/10000/p2p/" + peerID,
			expected: true,
		},
		{
			name:     "legacy ipfs segment",
			address:  "/ip4/127.0.0.1/tcp/10000/ipfs/" + peerID,
			expected: true,
		},
		{
			name:     "substring prefix is not accepted",
			address:  "/ip4/127.0.0.1/tcp/10000/p2p/prefix-" + peerID,
			expected: false,
		},
		{
			name:     "substring suffix is not accepted",
			address:  "/ip4/127.0.0.1/tcp/10000/p2p/" + peerID + "-suffix",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddressContainsPeerID(tt.address, peerID)
			if result != tt.expected {
				t.Fatalf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
