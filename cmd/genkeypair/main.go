package main

import (
	"fmt"
	"github.com/LoopNetworkRaspa/Raspa-Fork/cmd/kaspawallet/libkaspawallet"
	"github.com/LoopNetworkRaspa/Raspa-Fork/util"
)

func main() {
	cfg, err := parseConfig()
	if err != nil {
		panic(err)
	}

	privateKey, publicKey, err := libkaspawallet.CreateKeyPair(false)
	if err != nil {
		panic(err)
	}

	addr, err := util.NewAddressPublicKey(publicKey, cfg.NetParams().Prefix)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Private key: %x\n", privateKey)
	fmt.Printf("Address: %s\n", addr)
}
