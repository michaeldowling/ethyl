package ethyl

type EthereumNetworkError struct {
    Code    int64 `json:"code"`
    Message string `json:"message"`
}

type EthereumNetworkRequest struct {
    JsonRpcVersion string `json:"jsonrpc"`
    Method         string `json:"method"`
    Params         []interface{} `json:"params"`
    Id             string `json:"id"`
}

type EthereumFilterNetworkRequest struct {
    JsonRpcVersion string `json:"jsonrpc"`
    Method         string `json:"method"`
    Params         []FilterOptions `json:"params"`
    Id             string `json:"id"`
}

type StringResultEthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         string `json:"result"`
    Error          EthereumNetworkError `json:"error"`
}

type GenericSliceResultEthereumNetworkResponse struct {
    Id             string `json:"is"`
    JsonRpcVersion string `json:"jsonepc"`
    Result         []interface{} `json:"result"`
    Error          EthereumNetworkError `json:"error"`
}

type BooleanResultEthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         bool   `json:"result"`
    Error          EthereumNetworkError `json:"error"`
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

type TransactionReceiptObject struct {
    TransactionHash   string `json:"transactionHash"`
    TransactionIndex  string `json:"transactionIndex"`
    BlockHash         string `json:"blockHash"`
    BlockNumber       string `json:"blockNumber"`
    CumulativeGasUsed string `json:"cumulativeGasUsed"`
    GasUsed           string `json:"gasUsed"`
    ContractAddress   string `json:"contractAddress"`
    // Logs []string - TODO
}

type TransactionReceiptObjectResultEthereumNetworkResponse struct {
    Id             string `json:"id"`
    JsonRpcVersion string `json:"jsonrpc"`
    Result         TransactionReceiptObject `json:"result"`
}

type FilterOptions struct {
    FromBlock string `json:"fromBlock,omitempty"`
    ToBlock   string `json:"toBlock,omitempty"`
    Address   string `json:"address,omitempty"`
    Topics    []string `json:"topics,omitempty"`
}
