package ll

import (
	"fmt"
	"unsafe"
)

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

// PrintMemoryAndData prints the the given memory address
// and given the value of the passes systemAllocatedBytes
// it could be printing more than one memory.
// The data stored stored at these locations are also printed
// out.
func PrintMemoryAndData(firstByteAddress unsafe.Pointer, systemAllocatedBytes uintptr) {
	fmt.Println("Memory Address | Hex Value | Binary Value")
	fmt.Println("---------------|-----------|-------------")
	for i := uintptr(0); i < systemAllocatedBytes; i++ {
		address := (*byte)(unsafe.Pointer(uintptr(firstByteAddress) + i))
		byteValue := *address
		fmt.Printf("%p  |   0x%02x    |  %08b\n", address, byteValue, byteValue)
	}
}
