package ethyl

import (
    "net/http"
    "bytes"
    "io/ioutil"
    "encoding/json"
    "github.com/op/go-logging"
)

var LOGGER = logging.MustGetLogger("ethyl");

type EthylClient struct {
    Url     string;
    Accounts []string;

    Net      NetAPI;
    Eth      EthAPI;
}

func CreateClient(url string) (EthylClient, error) {

    client := EthylClient{Url:url};
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
    LOGGER.Infof("Request Body:  %s \n", requestParamsBytes);

    if (err != nil) {
        LOGGER.Infof("There was an error with mashaling the request params:  %v \n", err);
        return err;
    }

    response, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestParamsBytes));


    if (err != nil) {
        LOGGER.Errorf("[Call] There was an error with posting the request:  %v \n", err);
        return err;
    }
    defer response.Body.Close();

    responseBody, _ := ioutil.ReadAll(response.Body);
    LOGGER.Infof("Response Body:  %s", responseBody);

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
    LOGGER.Infof("Request Body:  %s \n", requestParamsBytes);

    if (err != nil) {
        return err;
    }

    response, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestParamsBytes));
    defer response.Body.Close();


    if (err != nil) {
        return err;
    }

    responseBody, _ := ioutil.ReadAll(response.Body);
    LOGGER.Infof("Response Body:  %s", responseBody);

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
    LOGGER.Infof("Request Body:  %s \n", requestParamsBytes);

    if (err != nil) {
        return err;
    }

    response, err := http.Post(client.Url, "application/json", bytes.NewBuffer(requestParamsBytes));
    defer response.Body.Close();


    if (err != nil) {
        return err;
    }

    responseBody, _ := ioutil.ReadAll(response.Body);
    LOGGER.Infof("Response Body:  %s", responseBody);

    json.Unmarshal(responseBody, replyValue);

    return nil;

}

