//go:generate abigen --abi ../../build/Compliance.abi --bin ../../build/Compliance.bin --pkg compliance --type Compliance --out compliance.go
//go:generate abigen --abi ../../build/ComplianceStorage.abi --bin ../../build/ComplianceStorage.bin --pkg compliance --type Storage --out complianceStorage.go
package compliance
