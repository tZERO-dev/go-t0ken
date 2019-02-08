//go:generate abigen --abi ../../build/AccountMap.abi --bin ../../build/AccountMap.bin --pkg mappings --type AccountMap --out accountMap.go
//go:generate abigen --abi ../../build/AddressMap.abi --bin ../../build/AddressMap.bin --pkg mappings --type AddressMap --out addressMap.go
package mappings
