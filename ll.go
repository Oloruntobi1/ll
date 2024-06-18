package ll

import "unsafe"

// DerefUnsafePointer is useful for situations where you have
// determined the address of a variable using the
// address := unsafe.Pointer(&x) //method (so this line be copied)
// and you want to know value stored at said address.
// It can also be called GetUnsafePointerAddressValue.
// NOTE: Every address points to specifically one byte
// and one byte of data alone.
func DerefUnsafePointer(address unsafe.Pointer) byte {
	return *(*byte)(address)
}
