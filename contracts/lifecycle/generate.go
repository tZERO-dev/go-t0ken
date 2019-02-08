//go:generate abigen --abi ../../build/AdminLockableDestroyable.abi --bin ../../build/AdminLockableDestroyable.bin --pkg lifecycle --type AdminLockableDestroyable --out adminLockableDestroyable.go
//go:generate abigen --abi ../../build/Destroyable.abi --bin ../../build/Destroyable.bin --pkg lifecycle --type Destroyable --out destroyable.go
//go:generate abigen --abi ../../build/Lockable.abi --bin ../../build/Lockable.bin --pkg lifecycle --type Lockable --out lockable.go
//go:generate abigen --abi ../../build/LockableDestroyable.abi --bin ../../build/LockableDestroyable.bin --pkg lifecycle --type LockableDestroyable --out lockableDestroyable.go
package lifecycle
