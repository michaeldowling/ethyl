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


