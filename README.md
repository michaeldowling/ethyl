# ethyl
A Golang Client Library for Communicating to Ethereum RPC Servers

## Usage

Ethyl attempts to match as closely as possible the JSON-RPC API that is published on the [[Ethereum Wiki (https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash)]]

### net_version

**Usage:**

    client, clientErr := CreateClient("localhost", 8545);
    assert.Nil(t, clientErr);
    
    version, _ := client.Net.Version();
    assert.True(t, version > 0, "Protocol Version should be greater than 0");


