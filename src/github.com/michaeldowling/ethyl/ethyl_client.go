package ethyl

import (
    "net/http"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "log"
)

type EthylClient struct {
    Host     string;
    Port     int;
    Accounts []string;

    Net      NetAPI;
    Eth      EthAPI;
}

func CreateClient(host string, port int) (EthylClient, error) {

    client := EthylClient{Host:host, Port:port};
    client.Net = CreateNetAPI(&client);
    client.Eth = CreateEthAPI(&client);

    // Pre-fetch accounts
    accounts, accountsErr := client.Eth.Accounts();
    if (accountsErr != nil) {
        return client, accountsErr;
    }

    client.Accounts = accounts;

    return client, nil;
}

func (client *EthylClient) Call(methodName string, args []string, replyValue interface{}) (error) {

    // Marshall into JSON string
    var requestParams EthereumNetworkRequest;
    if (args != nil) {

        // convert
        convertedArgs := make([]interface{}, len(args));
        for idx, val := range args {
            convertedArgs[idx] = val;
        }

        requestParams = EthereumNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName, Params:convertedArgs};

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

func (client *EthylClient) CallWithFilterOptions(methodName string, filterOptions FilterOptions, replyValue interface{}) (error) {

    // Marshall into JSON string
    var requestParams EthereumFilterNetworkRequest;
    filterOptionList := make([]FilterOptions, 1);
    filterOptionList[0] = filterOptions;

    requestParams = EthereumFilterNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName, Params:filterOptionList};

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

func (client *EthylClient) CallWithTransaction(methodName string, instructions TransactionInstructions, args []string, replyValue interface{}) (error) {

    // Marshall into JSON string
    var requestParams EthereumNetworkRequest;
    if (args != nil) {

        // convert
        convertedArgs := make([]interface{}, len(args) + 1);
        for idx, val := range args {
            convertedArgs[(idx + 1)] = val;
        }

        convertedArgs[0] = instructions;
        requestParams = EthereumNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName, Params:convertedArgs};

    } else {

        // convert
        convertedArgs := make([]interface{}, 1);
        convertedArgs[0] = instructions;
        requestParams = EthereumNetworkRequest{Id:"67", JsonRpcVersion:"2.0", Method:methodName, Params:convertedArgs};

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

