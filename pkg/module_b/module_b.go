package module_b

import (
	"github.com/btcsuite/btcutil/base58"
)

func Do() string {
	encoded := base58.Encode([]byte("test"))
	return encoded
}
