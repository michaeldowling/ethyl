package ethyl

import (
    "net/rpc"
    "log"
)

type EthylClient struct {
    Host      string;
    Port      int;

    Net       NetAPI;

    rpcClient *rpc.Client
}

func CreateClient(host string, port int) (EthylClient, error) {

    client := EthylClient{Host:host, Port:port};
    client.Net = CreateNetAPI(&client);

    return client, nil;
}

func (client *EthylClient) Dial() (error) {

    address := client.Host + ":8545";
    rpcDialClient, err := rpc.Dial("tcp", address);
    if (err != nil) {
        log.Fatal("Error Dialing into " + address + " | " + err.Error());
        return err;
    }

    client.rpcClient = rpcDialClient;

    return nil;

}

func (client *EthylClient) Call(methodName string, args interface{}, replyValue interface{}) (error) {

    dialError := client.Dial();
    if (dialError != nil) {
        return dialError;
    }

    callError := client.rpcClient.Call(methodName, args, replyValue);
    log.Println("Reply Value:  ");
    if (callError != nil) {
        return callError;
    }

    return nil;
}