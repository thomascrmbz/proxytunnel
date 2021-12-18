package auth

import (
	"github.com/gliderlabs/ssh"
	crypto_ssh "golang.org/x/crypto/ssh"
	"thomascrmbz.com/proxytunnel"
)

func AuthHandler(f func(crypto_ssh.PublicKey) bool, s ssh.Session) {
	if !isAuth(f, (s.Context().Value("sshPublicKey")).([]byte)) {
		s.Exit(int(proxytunnel.NOT_ALLOWED))
		return
	}
}

func isAuth(f func(crypto_ssh.PublicKey) bool, byteKey []byte) bool {
	key, err1 := ssh.ParsePublicKey(byteKey)
	publicKey, err2 := crypto_ssh.ParsePublicKey(key.Marshal())

	if err1 == nil && err2 == nil {
		return f(publicKey)
	}

	return false
}
