package ethereum

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/platform/bitcoin/blockbook"
	"github.com/trustwallet/blockatlas/platform/ethereum/bounce"
	"github.com/trustwallet/blockatlas/platform/ethereum/opensea"
	"github.com/trustwallet/blockatlas/platform/ethereum/trustray"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	CoinIndex   uint
	RpcURL      string
	client      EthereumClient
	collectible CollectibleClient
}

func Init(coinType uint, api, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		client:    &trustray.Client{Request: internal.InitClient(api)},
	}
}

func InitWithBlockbook(coinType uint, blockbookApi, rpc string) *Platform {
	return &Platform{
		CoinIndex: coinType,
		RpcURL:    rpc,
		client:    &blockbook.Client{Request: internal.InitClient(blockbookApi)},
	}
}

func InitWithOpenSea(coinType uint, rpc, blockbookApi, collectionApi, collectionKey string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi, rpc)
	platform.collectible = opensea.InitClient(collectionApi, collectionKey)
	return platform
}

func InitWithBounce(coinType uint, rpc, blockbookApi, collectionApi string) *Platform {
	platform := InitWithBlockbook(coinType, blockbookApi, rpc)
	platform.collectible = bounce.InitClient(collectionApi)
	return platform
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
