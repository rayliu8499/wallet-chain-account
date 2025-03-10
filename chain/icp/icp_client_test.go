package icp

import (
	"encoding/base64"
	"encoding/json"
	"github.com/dapplink-labs/wallet-chain-account/chain/evmbase"
	"testing"

	"github.com/dapplink-labs/wallet-chain-account/rpc/account"
	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/wallet-chain-account/chain"
	"github.com/dapplink-labs/wallet-chain-account/config"
)

func setup() (adaptor chain.IChainAdaptor, err error) {
	conf, err := config.New("../../config.yml")
	if err != nil {
		log.Error("load config failed, error:", err)
		return nil, err
	}
	adaptor, err = NewChainAdaptor(conf)
	if err != nil {
		log.Error("create chain adaptor failed, error:", err)
		return nil, err
	}
	return adaptor, nil
}

func TestChainAdaptor_GetSupportChains(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetSupportChains(&account.SupportChainsRequest{
		Chain: ChainName,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Info("isSupported:", rsp.GetSupport(), ", msg: ", rsp.GetMsg())
}

func TestChainAdaptor_ConvertAddress(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.ConvertAddress(&account.ConvertAddressRequest{
		Chain:     ChainName,
		PublicKey: "ab51a3b2dbc7c123ac8e93873611358fff297ea67ca9472125ba54af79c025e4",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Info("address:", rsp.GetAddress())
}

func TestChainAdaptor_ValidAddress(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.ValidAddress(&account.ValidAddressRequest{
		Chain:   ChainName,
		Address: "9ca25295843f8a11bcfb2cff56a2ed236d91aafa96c1584cf9befdf6a3e96c69",
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Info("isValid:", rsp.GetValid())
}

func TestChainAdaptor_GetBlockByNumber(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetBlockByNumber(&account.BlockNumberRequest{
		Chain:  ChainName,
		Height: 20675780,
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("get block by number:", rsp)
}

func TestChainAdaptor_GetBlockByHash(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetBlockByHash(&account.BlockHashRequest{
		Chain: ChainName,
		Hash:  "d105abebd1bf7325bca6917c04a73fed2b120a431c00c6996f1fa46d558d1b3f",
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("get block by hash:", rsp)
}

func TestChainAdaptor_GetAccount(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetAccount(&account.AccountRequest{
		Chain:   ChainName,
		Address: "e9ed7def415a8c323953f578c9ce0decf0031cb5b2a1f88cf6f1d89af80ed43a",
		Coin:    ChainName,
		//ContractAddress: "0x1Bdd8878252DaddD3Af2ba30628813271294eDc0",
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("get account:", rsp)
}

func TestChainAdaptor_GetFee(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetFee(&account.FeeRequest{
		Chain: ChainName,
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("fast fee:", rsp.GetFastFee(), "normal fee:", rsp.GetNormalFee(), "slow fee:", rsp.GetSlowFee())
}

func TestChainAdaptor_GetTxByAddress(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetTxByAddress(&account.TxAddressRequest{
		Chain:   ChainName,
		Address: "9ca25295843f8a11bcfb2cff56a2ed236d91aafa96c1584cf9befdf6a3e96c69",
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("transaction by address:", rsp)
}

func TestChainAdaptor_GetTxByHash(t *testing.T) {
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.GetTxByHash(&account.TxHashRequest{
		Chain: ChainName,
		Hash:  "f840978ecb382b415d6fd4ee9260a8c6ce09cdf10948eb595faea38fa6e41f0d",
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("transaction by hash:", rsp)
}

func TestChainAdaptor_BuildUnSignTx(t *testing.T) {

	base64Tx := createTestBase64Tx("", "e9ed7def415a8c323953f578c9ce0decf0031cb5b2a1f88cf6f1d89af80ed43a", "e9ed7def415a8c323953f578c9ce0decf0031cb5b2a1f88cf6f1d89af80ed43a", "100000")
	adaptor, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	rsp, err := adaptor.BuildUnSignTransaction(&account.UnSignTransactionRequest{
		Chain:    ChainName,
		Base64Tx: base64Tx,
	})
	if err != nil {
		t.Error(err)
	}
	log.Info("build unSign tx:", rsp)
}

func createTestBase64Tx(signature string, from string, to string, amount string) string {

	testTx := evmbase.Eip1559DynamicFeeTx{
		Nonce:       1,
		FromAddress: from,
		ToAddress:   to,
		Amount:      amount,
		Signature:   signature,
	}

	jsonBytes, err := json.Marshal(testTx)
	if err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(jsonBytes)
	return base64Str
}
