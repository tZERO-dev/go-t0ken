//go:generate abigen --abi ../../build/Compliance.abi --bin ../../build/Compliance.bin --pkg compliance --type Compliance --out compliance.go
//go:generate abigen --abi ../../build/ComplianceStorage.abi --bin ../../build/ComplianceStorage.bin --pkg compliance --type ComplianceStorage --out complianceStorage.go
//go:generate abigen --abi ../../build/T0kenCompliance.abi --bin ../../build/T0kenCompliance.bin --pkg compliance --type T0kenCompliance --out t0kenCompliance.go
package compliance
