package uuidfactory

import (
	"crypto/rand"
	"fmt"
	"io"
)

const VERSION byte = 0x40
const VARIANT byte = 0x80

func New() string {
	u := make([]byte, 16)

	if _, err := io.ReadFull(rand.Reader, u); err != nil {
		return ""
	}

	u[6] = (u[6] & 0x0f) | VERSION

	u[8] = (u[8] & 0x3f) | VARIANT

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		u[0:4],
		u[4:6],
		u[6:8],
		u[8:10],
		u[10:16],
	)
}
