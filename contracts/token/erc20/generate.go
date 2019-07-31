//go:generate abigen --abi ../../../build/T0ken.abi --bin ../../../build/T0ken.bin --pkg erc20 --type T0ken --out t0ken.go
//go:generate abigen --abi ../../../build/MigrateableT0ken.abi --bin ../../../build/MigrateableT0ken.bin --pkg erc20 --type MigrateableT0ken --out migrateableT0ken.go
//go:generate abigen --abi ../../../build/SplittableT0ken.abi --bin ../../../build/SplittableT0ken.bin --pkg erc20 --type SplittableT0ken --out splittableT0ken.go
//go:generate abigen --abi ../../../build/DynamicSplittableT0ken.abi --bin ../../../build/DynamicSplittableT0ken.bin --pkg erc20 --type DynamicSplittableT0ken --out dynamicSplittableT0ken.go
package erc20
