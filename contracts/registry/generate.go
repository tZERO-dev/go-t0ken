//go:generate abigen --abi ../../build/Storage.abi --bin ../../build/Storage.bin --pkg registry --type Storage --out storage.go
//go:generate abigen --abi ../../build/BrokerDealerRegistry.abi --bin ../../build/BrokerDealerRegistry.bin --pkg registry --type BrokerDealer --out brokerDealer.go
//go:generate abigen --abi ../../build/CustodianRegistry.abi --bin ../../build/CustodianRegistry.bin --pkg registry --type Custodian --out custodian.go
//go:generate abigen --abi ../../build/ExternalInvestorRegistry.abi --bin ../../build/ExternalInvestorRegistry.bin --pkg registry --type ExternalInvestor --out externalInvestor.go
//go:generate abigen --abi ../../build/InvestorRegistry.abi --bin ../../build/InvestorRegistry.bin --pkg registry --type Investor --out investor.go
package registry
