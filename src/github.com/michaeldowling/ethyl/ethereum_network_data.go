package ethyl

type EthereumNetworkRequest struct {
    JsonRpcVersion string `json:"jsonrpc"`
    Method         string `json:"method"`
    Params         []string `json:"params"`
    Id             string `json:"id"`
}

type StringResultEthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         string `json:"result"`
}

type BooleanResultEthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         bool   `json:"result"`
}




// for transactions
type TransactionObject struct {
    Hash             string `json:"hash"`
    Nonce            string `json:"nonce"`
    BlockHash        string `json:"blockHash"`
    BlockNumber      string `json:"blockNumber"`
    TransactionIndex string `json:"transactionIndex"`
    From             string `json:"from"`
    To               string `json:"to"`
    Value            string `json:"value"`
    Gas              string `json:"gas"`
    GasPrice         string `json:"gasPrice"`
    Input            string `json:"input"`
}

type TransactionObjectResultEthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         TransactionObject `json:"result"`
}

