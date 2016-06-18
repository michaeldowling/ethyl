# ethyl
A Golang Client Library for Communicating with Ethereum RPC Servers

## Usage

Ethyl attempts to match as closely as possible the JSON-RPC API that is published on the [Ethereum Wiki](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash)



_Import the library_:

    import (
        "github.com/michaeldowling/ethyl"
    )

_Get a client reference_:
    
    client, clientErr := ethyl.CreateClient("localhost", 8545);


## API:  Net
    

_Get the [Network Version](https://github.com/ethereum/wiki/wiki/JSON-RPC#net_version)_:

    version, _ := client.Net.Version();


_Get the [Peer Count](https://github.com/ethereum/wiki/wiki/JSON-RPC#net_peercount)_:

    peerCount, _ := client.Net.PeerCount();
    


