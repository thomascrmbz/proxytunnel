package auth

import (
	"github.com/gliderlabs/ssh"
	crypto_ssh "golang.org/x/crypto/ssh"
)

func AuthHandler(f func(crypto_ssh.PublicKey) bool, s ssh.Session) bool {
	return isAuth(f, (s.Context().Value("sshPublicKey")).([]byte))
}

func isAuth(f func(crypto_ssh.PublicKey) bool, byteKey []byte) bool {
	key, err1 := ssh.ParsePublicKey(byteKey)
	publicKey, err2 := crypto_ssh.ParsePublicKey(key.Marshal())

	if err1 == nil && err2 == nil {
		return f(publicKey)
	}

	return false
}
