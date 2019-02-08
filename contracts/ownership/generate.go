//go:generate abigen --abi ../../build/Administrable.abi --bin ../../build/Administrable.bin --pkg ownership --type Administrable --out administrable.go
//go:generate abigen --abi ../../build/AdminLockable.abi --bin ../../build/AdminLockable.bin --pkg ownership --type AdminLockable --out adminLockable.go
//go:generate abigen --abi ../../build/Ownable.abi --bin ../../build/Ownable.bin --pkg ownership --type Ownable --out ownable.go
package ownership
