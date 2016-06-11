package ethyl

import "log"

type NetAPI struct {
    Client *EthylClient
}

func CreateNetAPI(client *EthylClient) (NetAPI) {

    netapi := NetAPI{Client:client};
    return (netapi);

}

func (n *NetAPI) Version() (string, error) {

    var result string;

    n.Client.Call("net_version", "", &result);
    log.Println("Result: " + result);
    return result, nil;


}
