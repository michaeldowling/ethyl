package ethyl

import (
    "net/http"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "log"
)

type EthylClient struct {
    Host string;
    Port int;

    Net  NetAPI;
    Eth  EthAPI;

}

func CreateClient(host string, port int) (EthylClient, error) {

    client := EthylClient{Host:host, Port:port};
    client.Net = CreateNetAPI(&client);
    client.Eth = CreateEthAPI(&client);

    return client, nil;
}

func (client *EthylClient) Call(methodName string, args []string, replyValue interface{}) (error) {

    // Marshall into JSON string
    var requestParams EthereumNetworkRequest;
    if (args != nil) {
        requestParams = EthereumNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName, Params:args};
    } else {
        requestParams = EthereumNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName};
    }
    requestParamsBytes, err := json.Marshal(requestParams);
    log.Printf("Request Body:  %s \n", requestParamsBytes);

    if (err != nil) {
        return err;
    }

    response, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer(requestParamsBytes));
    defer response.Body.Close();


    if (err != nil) {
        return err;
    }

    responseBody, _ := ioutil.ReadAll(response.Body);
    log.Printf("Response Body:  %s", responseBody);

    json.Unmarshal(responseBody, replyValue);

    return nil;
}