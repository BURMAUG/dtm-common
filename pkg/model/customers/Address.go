package customers

//Everything that has to do with address should be here

type AddressInfo interface {
	CreateNewAddress(address, city, state, zip string)
	GetAddress() string
}

type Address struct {
	address string
	city    string
	state   string
	zip     string
}

func (a *Address) CreateNewAddress(address, city, state, zip string) Address {
	return Address{}
}

func (a *Address) GetAddress() string {
	return a.address + " " + a.address + " " + a.city + "  " + a.zip
}
