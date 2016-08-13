/*
 * (C) Copyright 2016 Michael Dowling (http://github.com/michaeldowling) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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