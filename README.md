# ethyl
A Golang Client Library for Communicating to Ethereum RPC Servers

## Usage

Ethyl attempts to match as closely as possible the JSON-RPC API that is published on the [Ethereum Wiki](https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash)


**Usage:**

_Import the library_:

    import (
        "github.com/michaeldowling/ethyl"
    )

_Get a client reference_:
    
    client, clientErr := ethyl.CreateClient("localhost", 8545);
    

_API:  Get the Client Version_:

    version, _ := client.Net.Version();




