package hash

import (
	"crypto/aes"
	"crypto/cipher"
	rand1 "crypto/rand"
	rand "math/rand"
	"time"
)

type PasswordGenerator struct {
	Password []byte
}

func NewPassword(length int) PasswordGenerator {
	password := gen(length)
	return PasswordGenerator{Password: password}
}

func gen(length int) []byte {
	items := []string{"qwertyuiopasdfghjklzxcvbnm", "QWERTYUIOPASDFGHJKLZXCVBNM", "!'^#$+%&/|?=-_<>(){}[]", "0123456789"}
	size := length / 4
	rand.Seed(time.Now().UnixNano())
	var count [4]int = [4]int{0, 0, 0, 0}
	var pass []byte
	last := -1
	for i := 0; i < length; {
		rand1 := rand.Intn(4)
		if rand1 == last {
			continue
		}
		if count[rand1] >= size {
			if count[0] == size && count[1] == size && count[2] == size && count[3] == size {
				break
			}
			last = rand1
			continue
		}
		count[rand1] += 1
		rand2 := rand.Intn(len(items[rand1]) - 1)
		pass = append(pass, items[rand1][rand2])
		i += 1
		last = rand1
	}
	return pass
}

func (p PasswordGenerator) Encrypt(key []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand1.Read(nonce); err != nil {
		return nil, err
	}
	encrypted := gcm.Seal(nonce, nonce, p.Password, nil)
	return encrypted, nil
}

func Decrypt(key []byte, encryptedPassword []byte) ([]byte, error) {
	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}
	nonce, encrypted := encryptedPassword[:gcm.NonceSize()], encryptedPassword[gcm.NonceSize():]
	decrypted, err := gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		return nil, err
	}
	return decrypted, nil
}
