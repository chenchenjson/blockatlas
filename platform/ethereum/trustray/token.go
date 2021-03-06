package trustray

import (
	"github.com/trustwallet/golibs/types"
)

func (c *Client) GetTokenList(address string, coinIndex uint) ([]string, error) {
	account, err := c.GetTokens(address)
	if err != nil {
		return nil, err
	}
	return NormalizeTokens(account.Docs, coinIndex), nil
}

// NormalizeToken converts a Ethereum token into the generic model
func NormalizeToken(srcToken *Contract, coinIndex uint) types.Token {
	tokenType := types.GetEthereumTokenTypeByIndex(coinIndex)

	return types.Token{
		Name:     srcToken.Name,
		Symbol:   srcToken.Symbol,
		TokenID:  srcToken.Address,
		Coin:     coinIndex,
		Decimals: srcToken.Decimals,
		Type:     tokenType,
	}
}

// NormalizeTxs converts multiple Ethereum tokens
func NormalizeTokens(srcTokens []Contract, coinIndex uint) []string {
	assetIds := make([]string, 0)
	for _, srcToken := range srcTokens {
		token := NormalizeToken(&srcToken, coinIndex)
		assetIds = append(assetIds, token.AssetId())
	}
	return assetIds
}
