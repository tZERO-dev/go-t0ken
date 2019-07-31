//go:generate abigen --abi ../../../build/ComplianceRule.abi --bin ../../../build/ComplianceRule.bin --pkg rules --type Rule --out rule.go
//go:generate abigen --abi ../../../build/Restrict.abi --bin ../../../build/Restrict.bin --pkg rules --type Restrict --out restrict.go
//go:generate abigen --abi ../../../build/RestrictAll.abi --bin ../../../build/RestrictAll.bin --pkg rules --type RestrictAll --out restrictAll.go
//go:generate abigen --abi ../../../build/RestrictDuringLockup.abi --bin ../../../build/RestrictDuringLockup.bin --pkg rules --type RestrictDuringLockup --out restrictDuringLockup.go
//go:generate abigen --abi ../../../build/RestrictFromAdvisor.abi --bin ../../../build/RestrictFromAdvisor.bin --pkg rules --type RestrictFromAdvisor --out restrictFromAdvisor.go
//go:generate abigen --abi ../../../build/RestrictFromAffiliate.abi --bin ../../../build/RestrictFromAffiliate.bin --pkg rules --type RestrictFromAffiliate --out restrictFromAffiliate.go
//go:generate abigen --abi ../../../build/RestrictFromInvestor.abi --bin ../../../build/RestrictFromInvestor.bin --pkg rules --type RestrictFromInvestor --out restrictFromInvestor.go
//go:generate abigen --abi ../../../build/RestrictHolderThreshold.abi --bin ../../../build/RestrictHolderThreshold.bin --pkg rules --type RestrictHolderThreshold --out restrictHolderThreshold.go
//go:generate abigen --abi ../../../build/RestrictToAccreditedInvestor.abi --bin ../../../build/RestrictToAccreditedInvestor.bin --pkg rules --type RestrictToAccreditedInvestor --out restrictToAccreditedInvestor.go
//go:generate abigen --abi ../../../build/RestrictToBrokerOrCustodialAccount.abi --bin ../../../build/RestrictToBrokerOrCustodialAccount.bin --pkg rules --type RestrictToBrokerOrCustodialAccount --out restrictToBrokerOrCustodialAccount.go
//go:generate abigen --abi ../../../build/RestrictToContract.abi --bin ../../../build/RestrictToContract.bin --pkg rules --type RestrictToContract --out restrictToContract.go
//go:generate abigen --abi ../../../build/RestrictToCustodianOrCustodialAccountOrBroker.abi --bin ../../../build/RestrictToCustodianOrCustodialAccountOrBroker.bin --pkg rules --type RestrictToCustodianOrCustodialAccountOrBroker --out restrictToToCustodianOrCustodialAccountOrBroker.go
//go:generate abigen --abi ../../../build/RestrictTransferFrom.abi --bin ../../../build/RestrictTransferFrom.bin --pkg rules --type RestrictTransferFrom --out restrictTransferFrom.go
package rules
