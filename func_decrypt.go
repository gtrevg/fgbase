package fgbase

import (
	"golang.org/x/crypto/nacl/box"
)

type dec struct {
	LocalPrivateKey, RemotePublicKey *[32]byte
	Nonce                            [24]byte
}

func decryptFire(n *Node) error {
	a := n.Srcs[0]
	x := n.Dsts[0]
	s := n.Aux.(*dec)
	privateKey := s.LocalPrivateKey
	publicKey := s.RemotePublicKey
	nonce := s.Nonce

	av := a.SrcGet()
	before := av.([]byte)
	after, ok := box.Open(nil, before, &nonce, publicKey, privateKey)
	_ = ok
	x.DstPut(string(after))
	return nil
}

// FuncDecrypt decrypts a buffer of byte data
func FuncDecrypt(a Edge, x Edge, localPrivateKey, remotePublicKey *[32]byte, nonce [24]byte) Node {

	node := MakeNode("decrypt", []*Edge{&a}, []*Edge{&x}, nil, decryptFire)
	node.Aux = &dec{LocalPrivateKey: localPrivateKey, RemotePublicKey: remotePublicKey, Nonce: nonce}
	return node

}
