// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	"enode://fcaa4976e879df791db6553afca73a2951b3e04f03efb71b2694b3e4e6554ee73ff0f98c3b7f85ac032bdba4ce2988611032fe619bb3cd263f217d3c67acf1b4@207.148.124.173:33177",
	"enode://c9df13815253d839a6cfe9bdb78c35947f2bb7bbd4cf2f50aee6215e3a8b9bbb307d4651381b36009fe382c9c177ee771f709ac515b24c7a9550c7d15907f69d@45.32.114.207:33177",
	"enode://52cc75cd9490203cb60a376c3ac51153076e580759f09333f58d628948850beadcd45730647fe63811f7166a6259f332cdd70e5696ce8182039e7a9bb65500e0@139.180.212.77:33177",
	"enode://d85bc7cb4eacadf35c6d70c4e5e91a401879222f295722928223d3bb0a0c888f27b86c9fcebf5243b3958f724ee9e7ac5c816777a73ad9088639d96996f73cac@45.77.34.176:33177",
	"enode://c25dcf682f99d9418d3a282945c287a0f19d14bde24f7ee00e9a9780a6d53946f86f9ea653e802424b6ac13f9e66565b591644183b55faef2a6ef91eed565843@139.180.208.204:33177",
	"enode://5cd06b81f27b299e943b059dd5ebeee3ed7a5e30613eb47b17ef1ce56afd210d0a24539e4b7130c9af36751bbcae37dd7839566ee4bfdb3b7f83f363a07e0a62@149.28.143.238:33177",
	"enode://ad26f9b184b981f4f5fae9b8561f072eb5538d669b92c277e842f0615fc81acfba98cfd46a3db2aacaef321e7f9806d2952b381497cad80eb3693082250be417@45.32.103.123:33177",
	"enode://750cd34195bf9ca12cb5c39b8ee8ddd5d5537ead275f49dad7889d2b719d9de6f5d2094ddda1b5fb6dd12a2a41a9b6b17f8d5d1b6879bdab4c9757098506c344@139.180.222.55:33177",
	"enode://8da8c1e5ae5a16464e00eedffce06e4e68610f32fbe65e9a4be3c914d4ed677cbd95c57267d37b11a593f5976ae79ea3c6ad75e81e53620c0b48ce677e645a73@149.28.142.123:33177",
	"enode://09a39ff52c79ab12caf93e2f443e2fe4316dd07d59b88f7053ea0f5f5c21ae1d01092a0f8b215043fe07623691b89ed8a674352e373fe8adb0293d9d32e6311a@139.180.220.172:33177",
}
var TestnetBootnodes = []string{}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
