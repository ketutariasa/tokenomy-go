// Copyright 2019 Tokenomy Technologies Pte. Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package tokenomy

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//
// Environment contains default and dynamic values that gathered from external
// resources, for example system environment variables.
//
type Environment struct {
	//
	// Debug define level of logging in our library.
	// debug value is set from environment variable "TOKENOMY_DEBUG".
	// TOKENOMY_DEBUG=1 is for logging in configuration.
	// TOKENOMY_DEBUG=2 is for logging input and output.
	//
	Debug int

	// Address of API server, optional. It will default to DefaultAddress
	// constant on each package.
	Address string

	// Token, required, is the public part of API key.
	Token string

	// Secret, required, is the private part of API key.
	Secret string

	// IsInsecure, optional, allow self-signed certificate, should be use
	// for testing only.
	IsInsecure bool

	pairs map[string]struct{}
}

//
// NewEnvironment create and initialize environment.
//
// If token and/or secret is empty it will set from environment variables
// TOKENOMY_TOKEN and TOKENOMY_SECRET.
//
func NewEnvironment(token, secret string) (env *Environment) {
	log.SetFlags(0)

	env = &Environment{
		Address: os.Getenv(EnvNameAddress),
		Token:   os.Getenv(EnvNameToken),
		Secret:  os.Getenv(EnvNameSecret),
	}

	if len(token) > 0 {
		env.Token = token
	}
	if len(secret) > 0 {
		env.Secret = secret
	}

	v := os.Getenv(EnvNameDebug)
	if len(v) > 0 {
		env.Debug, _ = strconv.Atoi(v)
	}

	if env.Debug >= 1 {
		fmt.Printf(">>> Environment: %+v\n", env)
	}

	env.initializePairs()

	return env
}

//
// IsValidPairName will return true if pairName value is valid pair name;
// otherwise it will return false.
//
func (env *Environment) IsValidPairName(pairName string) (ok bool) {
	_, ok = env.pairs[pairName]
	return ok
}

func (env *Environment) initializePairs() {
	env.pairs = make(map[string]struct{}, 38)

	env.pairs[PairBitcoinabcBitcoin] = struct{}{}
	env.pairs[PairBitcoinsvBitcoin] = struct{}{}
	env.pairs[PairBittorrentBitcoin] = struct{}{}
	env.pairs[PairEosBitcoin] = struct{}{}
	env.pairs[PairEthclassicBitcoin] = struct{}{}
	env.pairs[PairEthereumBitcoin] = struct{}{}
	env.pairs[PairHonestBitcoin] = struct{}{}
	env.pairs[PairLitecoinBitcoin] = struct{}{}
	env.pairs[PairLoopringneoBitcoin] = struct{}{}
	env.pairs[PairLyfeBitcoin] = struct{}{}
	env.pairs[PairMoneroBitcoin] = struct{}{}
	env.pairs[PairOntologyBitcoin] = struct{}{}
	env.pairs[PairPlaygameBitcoin] = struct{}{}
	env.pairs[PairPundixBitcoin] = struct{}{}
	env.pairs[PairSixBitcoin] = struct{}{}
	env.pairs[PairStellarBitcoin] = struct{}{}
	env.pairs[PairStoriqaBitcoin] = struct{}{}
	env.pairs[PairTokenomyBitcoin] = struct{}{}
	env.pairs[PairTronBitcoin] = struct{}{}
	env.pairs[PairVexaniumBitcoin] = struct{}{}
	env.pairs[PairZcashBitcoin] = struct{}{}

	env.pairs[PairBitcoinIdk] = struct{}{}
	env.pairs[PairTetherIdk] = struct{}{}

	env.pairs[PairHaraEthereum] = struct{}{}
	env.pairs[PairInmaxEthereum] = struct{}{}
	env.pairs[PairPundixEthereum] = struct{}{}
	env.pairs[PairStoriqaEthereum] = struct{}{}
	env.pairs[PairTokenomyEthereum] = struct{}{}
	env.pairs[PairTronEthereum] = struct{}{}
	env.pairs[PairVexaniumEthereum] = struct{}{}

	env.pairs[PairBitcoinTether] = struct{}{}
	env.pairs[PairDaexTether] = struct{}{}
	env.pairs[PairEthereumTether] = struct{}{}
	env.pairs[PairTokenomyTether] = struct{}{}

	env.pairs[PairSixTokenomy] = struct{}{}
	env.pairs[PairStoriqaTokenomy] = struct{}{}
}
