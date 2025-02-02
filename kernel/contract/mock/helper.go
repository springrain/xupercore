package mock

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/xuperchain/xupercore/kernel/contract/bridge"

	"github.com/golang/protobuf/proto"
	"github.com/xuperchain/xupercore/kernel/contract"
	"github.com/xuperchain/xupercore/kernel/contract/sandbox"
	"github.com/xuperchain/xupercore/kernel/ledger"
	"github.com/xuperchain/xupercore/kernel/permission/acl/utils"
	"github.com/xuperchain/xupercore/protos"
)

const (
	ContractAccount      = "XC1111111111111111@xuper"
	ContractAccount2     = "XC2222222222222222@xuper"
	FeaturesContractName = "features"
)

type TestHelper struct {
	basedir    string
	utxo       *contract.UTXORWSet
	utxoReader sandbox.UtxoReader
	state      *sandbox.MemXModel
	manager    contract.Manager
}

func NewTestHelper(cfg *contract.ContractConfig) *TestHelper {
	basedir, err := ioutil.TempDir("", "contract-test")
	if err != nil {
		panic(err)
	}

	state := sandbox.NewMemXModel()
	core := new(fakeChainCore)
	m, err := contract.CreateManager("default", &contract.ManagerConfig{
		Basedir:  basedir,
		BCName:   "xuper",
		Core:     core,
		XMReader: state,
		Config:   cfg,
	})
	if err != nil {
		panic(err)
	}

	th := &TestHelper{
		basedir: basedir,
		manager: m,
		state:   state,
	}
	th.initAccount()
	return th
}

func (t *TestHelper) Manager() contract.Manager {
	return t.manager
}

func (t *TestHelper) Basedir() string {
	return t.basedir
}

func (t *TestHelper) State() *sandbox.MemXModel {
	return t.state
}
func (t *TestHelper) UTXOState() *contract.UTXORWSet {
	return t.utxo
}

func (t *TestHelper) initAccount() {
	t.state.Put(utils.GetAccountBucket(), []byte(ContractAccount), &ledger.VersionedData{
		RefTxid:  []byte("txid"),
		PureData: nil,
	})

	utxoReader := sandbox.NewUTXOReaderFromInput([]*protos.TxInput{
		{
			RefTxid:      nil,
			RefOffset:    0,
			FromAddr:     []byte(FeaturesContractName),
			Amount:       big.NewInt(9999).Bytes(),
			FrozenHeight: 0,
		},
	})

	t.utxoReader = utxoReader
}

func (t *TestHelper) Deploy(module, lang, contractName string, bin []byte, args map[string][]byte) (*contract.Response, error) {
	m := t.Manager()
	state, err := m.NewStateSandbox(&contract.SandboxConfig{
		XMReader:   t.State(),
		UTXOReader: t.utxoReader,
	})
	if err != nil {
		return nil, err
	}

	ctx, err := m.NewContext(&contract.ContextConfig{
		Module:         "xkernel",
		ContractName:   "$contract",
		State:          state,
		ResourceLimits: contract.MaxLimits,
		Initiator:      ContractAccount,
	})
	if err != nil {
		return nil, err
	}

	desc := &protos.WasmCodeDesc{
		Runtime:      lang,
		ContractType: module,
	}
	descbuf, _ := proto.Marshal(desc)

	argsBuf, _ := json.Marshal(args)

	invokeArgs := map[string][]byte{
		"account_name":  []byte(ContractAccount),
		"contract_name": []byte(contractName),
		"contract_code": bin,
		"contract_desc": descbuf,
		"init_args":     argsBuf,
	}
	if bridge.ContractType(module) == bridge.TypeEvm {
		invokeArgs["contract_abi"] = args["contract_abi"]
	}
	resp, err := ctx.Invoke("deployContract", invokeArgs)

	if err != nil {
		return nil, err
	}

	ctx.Release()
	t.Commit(state)
	return resp, nil
}

func (t *TestHelper) Upgrade(contractName string, bin []byte) error {
	m := t.Manager()
	state, err := m.NewStateSandbox(&contract.SandboxConfig{
		XMReader:   t.State(),
		UTXOReader: t.utxoReader,
	})
	if err != nil {
		return err
	}

	ctx, err := m.NewContext(&contract.ContextConfig{
		Module:         "xkernel",
		ContractName:   "$contract",
		State:          state,
		ResourceLimits: contract.MaxLimits,
	})
	if err != nil {
		return err
	}

	_, err = ctx.Invoke("upgradeContract", map[string][]byte{
		"contract_name": []byte(contractName),
		"contract_code": bin,
	})
	ctx.Release()
	t.Commit(state)
	return err
}

func (t *TestHelper) Invoke(module, contractName, method string, args map[string][]byte) (*contract.Response, error) {
	m := t.Manager()
	state, err := m.NewStateSandbox(&contract.SandboxConfig{
		XMReader:   t.State(),
		UTXOReader: t.utxoReader,
	})
	if err != nil {
		return nil, err
	}

	ctx, err := m.NewContext(&contract.ContextConfig{
		Module:         module,
		ContractName:   contractName,
		State:          state,
		ResourceLimits: contract.MaxLimits,
		Initiator:      ContractAccount,
	})
	if err != nil {
		return nil, err
	}
	defer ctx.Release()

	resp, err := ctx.Invoke(method, args)
	if err != nil {
		return nil, err
	}
	state.Flush()
	t.utxo = state.UTXORWSet()
	t.Commit(state)
	return resp, nil
}

func (t *TestHelper) Commit(state contract.StateSandbox) {
	rwset := state.RWSet()
	txbuf := make([]byte, 32)
	rand.Read(txbuf)
	for i, w := range rwset.WSet {
		t.state.Put(w.Bucket, w.Key, &ledger.VersionedData{
			RefTxid:   txbuf,
			RefOffset: int32(i),
			PureData: &ledger.PureData{
				Bucket: w.Bucket,
				Key:    w.Key,
				Value:  w.Value,
			},
		})
	}
}

func (t *TestHelper) Close() {
	os.RemoveAll(t.basedir)
}
