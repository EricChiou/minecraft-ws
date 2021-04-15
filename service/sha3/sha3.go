package sha3

import (
	"minecraft-ws/config"
	"minecraft-ws/pkg/sha3"
)

// Encrypt sha3 encrypt
func Encrypt(pwd string, lenght ...int) string {
	if len(lenght) > 0 {
		return sha3.Encrypt(pwd, lenght[0])
	}
	return sha3.Encrypt(pwd, config.Sha3Len)
}
