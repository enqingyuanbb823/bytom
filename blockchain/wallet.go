package blockchain

import (
	"context"

	"github.com/bytom/crypto/ed25519/chainkd"
)

func (bcr *BlockchainReactor) walletExportKey(ctx context.Context, in struct {
	Password string       `json:"password"`
	XPub     chainkd.XPub `json:"xpub"`
}) Response {
	type privateKey struct {
		PrivateKey string `json:"private_key"`
	}

	key, err := bcr.wallet.ExportAccountPrivKey(bcr.hsm, in.XPub, in.Password)
	if err != nil {
		return resWrapper(nil, err)
	}
	return resWrapper(&privateKey{PrivateKey: *key})
}

func (bcr *BlockchainReactor) walletImportKey(ctx context.Context, in struct {
	Alias    string       `json:"alias"`
	Password string       `json:"password"`
	XPrv     chainkd.XPrv `json:"xprv"`
	Index    uint64       `json:"index"`
}) Response {
	xpub, err := bcr.wallet.ImportAccountPrivKey(bcr.hsm, in.XPrv, in.Alias, in.Password, in.Index)
	if err != nil {
		return resWrapper(nil, err)
	}
	return resWrapper(xpub)
}
