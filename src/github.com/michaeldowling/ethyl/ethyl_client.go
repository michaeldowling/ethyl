package ethyl

import (
    "net/http"
    "bytes"
    "io/ioutil"
    "encoding/json"
)

type EthylClient struct {
    Host      string;
    Port      int;

    Net       NetAPI;

}

func CreateClient(host string, port int) (EthylClient, error) {

    client := EthylClient{Host:host, Port:port};
    client.Net = CreateNetAPI(&client);

    return client, nil;
}

func (client *EthylClient) Call(methodName string, args string, replyValue *EthereumNetworkResponse) (error) {

    // Marshall into JSON string
    requestParams := EthereumNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName};
    responseDocument := EthereumNetworkResponse{};
    requestParamsBytes, err := json.Marshal(requestParams);

    if (err != nil) {
        return err;
    }

    response, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(requestParamsBytes));
    defer response.Body.Close();


    if (err != nil) {
        return err;
    }

    responseBody, _ := ioutil.ReadAll(response.Body);
    json.Unmarshal(responseBody, &responseDocument);

    *replyValue = responseDocument;

    return nil;
}