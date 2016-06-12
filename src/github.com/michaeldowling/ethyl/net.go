package ethyl

import (
    "log"
    "encoding/hex"
    "encoding/binary"
    "bytes"
)

type NetAPI struct {
    Client *EthylClient
}

func CreateNetAPI(client *EthylClient) (NetAPI) {

    netapi := NetAPI{Client:client};
    return (netapi);

}

func (n *NetAPI) Version() (string, error) {

    var result StringResultEthereumNetworkResponse;

    n.Client.Call("net_version", nil, &result);
    version := result.Result;


    log.Println("Version: " + version);
    return version, nil;

}

func (n *NetAPI) IsListening() (bool, error) {

    var result BooleanResultEthereumNetworkResponse;

    n.Client.Call("net_listening", nil, &result);
    listening := result.Result;

    return listening, nil;

}

func (n *NetAPI) PeerCount() (uint64, error) {

    var result StringResultEthereumNetworkResponse;

    n.Client.Call("net_peerCount", nil, &result);
    peerCountHex := result.Result;

    log.Println("Peer Count Hex:  " + peerCountHex);

    // convert from hex
    peerCountBytes, _ := hex.DecodeString(peerCountHex);

    peerCount, err := binary.ReadUvarint(bytes.NewBuffer(peerCountBytes));
    return peerCount, err;

}