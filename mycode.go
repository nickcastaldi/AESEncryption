import (
    "crypto/rand"
    "fmt"
)
	// Function to perform S-Box substitution
	var sbox [256]byte

	if inverse {
		sbox = sbox1
	} else {
		sbox = sbox0
	}

	for i := 0; i < len(block); i++ {
		block[i] = sbox[block[i]]
	}

	//Function for Sub Key generation

	func generateSubkey(masterKey []byte, length int) []byte {
		// Seed the random number generator with the masterKey
		rand.Seed(int64(binary.BigEndian.Uint64(masterKey)))
	
		randomHash := make([]byte, length)
		_, err := rand.Read(randomHash)
		if err != nil {
			panic(err)
		}
	
		subkey := make([]byte, length)
		for i := range subkey {
			subkey[i] = masterKey[i%len(masterKey)] ^ randomHash[i]
		}
	
		return subkey
	}
	   //Function for 10 rounds of subkey generation
	func generateAndPrintSubkeys(masterKey []byte, numRounds, length int) {
		for round := 0; round < numRounds; round++ {
			subkey := generateSubkey(masterKey, length)
			fmt.Printf("Round %d: Subkey = %x\n", round+1, subkey)
		}
	}

	
