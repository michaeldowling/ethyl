# ethyl
A Golang Client Library for Communicating with Ethereum RPC Servers

- [Usage](#usage)
- [Net API](#api-net)
- [Ethereum API](#api-ethereum)
- [Contracts](#api-contracts)



## Usage

Ethyl attempts to match as closely as possible the JSON-RPC API that is published on the [Ethereum Wiki](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash)



_Import the library_:

    import (
        "github.com/michaeldowling/ethyl"
    )

_Get a client reference_:
    
    client, clientErr := ethyl.CreateClient("localhost", 8545);


## API: Net
    

_Get the [Network Version](https://github.com/ethereum/wiki/wiki/JSON-RPC#net_version)_:

    version, _ := client.Net.Version();


_Get the [Peer Count](https://github.com/ethereum/wiki/wiki/JSON-RPC#net_peercount)_:

    peerCount, _ := client.Net.PeerCount();
    

_Is the node [listening](https://github.com/ethereum/wiki/wiki/JSON-RPC#net_listening) for network connections?_:

    isListening, _ := client.Net.IsListening();


## API: Ethereum

_Get the current [Ethereum Protocol Version](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_protocolversion)_:

    version, _ := client.Eth.ProtocolVersion();
    
_Get the [list of addresses/accounts](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_accounts) owned by client_:

    accounts, _ := client.Eth.Accounts();
    for idx, account := range accounts {
        log.Printf("Account Address #%i:  %s \n", idx, account);
    }
    
_Get the [current price per gas](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gasprice) in wei_:

    gasPrice, _ := client.Eth.GasPrice();
    
_Get the node's [sync status](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_syncing)_:

    isSyncing, startingBlock, currentBlock, highestBlock , err := client.Eth.Syncing();


_[Send a transaction](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction) onto the network_:

    client, clientErr := CreateClient("localhost", 8545);
    instructions := TransactionInstructions{From:client.Accounts[0], To:"0xabd5d148b31f38a8d2aa9eb041c478d36dd51c35", Gas:200000, Value:10};
    txhash, err := client.Eth.SendTransaction(instructions);
    

_Get the [information about a transaction](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionbyhash), by the Transaction Hash_:

    transactionObject, _ := client.Eth.GetTransactionByHash(txhash);
    
_Get the [receipt for an executed transaction](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_gettransactionreceipt), by the Transaction Hash_:

    transactionReceipt, _ := client.Eth.GetTransactionReceipt(txhash);
    
    
    
    


## API: Contracts

_Import the library_:

    import (
        "github.com/michaeldowling/ethyl"
        "github.com/michaeldowling/ethyl/contracts"
    )


_Define a contract for deployment_:

    contractBytecode := "6060604052604051602080610727833981016040528080519060200190919050505b33600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690830217905550600160016000506000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050600001600050819055508060ff166002600050818154818355818115116100fc578183600052602060002091820191016100fb91906100d9565b808211156100f75760006000820160005060009055506001016100d9565b5090565b5b505050505b50610617806101106000396000f360606040526000357c0100000000000000000000000000000000000000000000000000000000900480635c19a95c1461005a578063609ff1bd146100725780639e7b8d6114610098578063b3f98adc146100b057610058565b005b61007060048080359060200190919050506100c8565b005b61007f60048050506103c4565b604051808260ff16815260200191505060405180910390f35b6100ae600480803590602001909190505061045b565b005b6100c6600480803590602001909190505061053d565b005b60006000600160005060003373ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005091508160010160009054906101000a900460ff1615610119576103bf565b5b600073ffffffffffffffffffffffffffffffffffffffff16600160005060008573ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060010160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415801561022757503373ffffffffffffffffffffffffffffffffffffffff16600160005060008573ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060010160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614155b1561028857600160005060008473ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060010160029054906101000a900473ffffffffffffffffffffffffffffffffffffffff169250825061011a565b3373ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156102c1576103bf565b60018260010160006101000a81548160ff02191690830217905550828260010160026101000a81548173ffffffffffffffffffffffffffffffffffffffff02191690830217905550600160005060008473ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005090508060010160009054906101000a900460ff16156103a257816000016000505460026000508260010160019054906101000a900460ff1660ff16815481101561000257906000526020600020900160005b506000016000828282505401925050819055506103be565b8160000160005054816000016000828282505401925050819055505b5b505050565b60006000600060009150600090505b6002600050805490508160ff161015610455578160026000508260ff16815481101561000257906000526020600020900160005b506000016000505411156104475760026000508160ff16815481101561000257906000526020600020900160005b50600001600050549150815080925082505b5b80806001019150506103d3565b5b505090565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415806104f45750600160005060008273ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060010160009054906101000a900460ff165b156104fe5761053a565b6001600160005060008373ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050600001600050819055505b50565b6000600160005060003373ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005090508060010160009054906101000a900460ff168061059857506002600050805490508260ff1610155b156105a257610613565b60018160010160006101000a81548160ff02191690830217905550818160010160016101000a81548160ff02191690830217905550806000016000505460026000508360ff16815481101561000257906000526020600020900160005b506000016000828282505401925050819055505b505056";
    contract := contracts.DefineContract(client, abiString, contractBytecode);
    
    deployConfig := DeployConfig{Gas:300000};
    deployResults, err := contract.Deploy(client, deployConfig);
    
_If you want, you can actively monitor the contract's state on the blockchain_:
 
    deployChan := contract.MonitorDeploy(client, deployResults);
    transactionReceipt := <-deployChan;
    
    
    
