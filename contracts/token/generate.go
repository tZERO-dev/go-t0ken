//go:generate abigen --abi ../../build/T0ken.abi --bin ../../build/T0ken.bin --pkg token --type T0ken --out t0ken.go
//go:generate abigen --abi ../../build/T0kenMigrateable.abi --bin ../../build/T0kenMigrateable.bin --pkg token --type T0kenMigrateable --out t0kenMigrateable.go
//go:generate abigen --abi ../../build/T0kenSplittable.abi --bin ../../build/T0kenSplittable.bin --pkg token --type T0kenSplittable --out t0kenSplittable.go
//go:generate abigen --abi ../../build/T0kenSplittableDynamic.abi --bin ../../build/T0kenSplittableDynamic.bin --pkg token --type T0kenSplittableDynamic --out t0kenSplittableDynamic.go
package token
