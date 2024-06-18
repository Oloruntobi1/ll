package ll

import "unsafe"

// DerefUnsafePointer is useful for situations where you have
// determined the address of a variable using the
// address := unsafe.Pointer(&x) //method (so this line be copied)
// and you want to know value of that byte.
// Basically it helps with dereference an unsafe.Pointer type.
func DerefUnsafePointer(address unsafe.Pointer) byte {
	return *(*byte)(address)
}
