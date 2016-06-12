package contracts

type Contract struct {
    ABI     string
    Address string;
}

func DefineContract(abi string) (Contract, error) {

    return Contract{ABI:abi}, nil;

}

func At(abi string, address string) (Contract, error) {

    c, err := DefineContract(abi);
    c.Address = address;
    return c, err;

}

func (c *Contract) New() (error) {

    return nil;

}


