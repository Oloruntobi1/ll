package ll

import (
	"fmt"
	"unsafe"
)

// DerefUnsafePointer retrieves the value stored at a given memory address.
// This function is useful when you have determined the address of a variable
// using the unsafe.Pointer(&x) method and you want to know the value stored
// at that address.
//
// Example usage:
//
//	address := unsafe.Pointer(&x)
//	value := DerefUnsafePointer(address)
//
// Parameters:
//   - address: The memory address to dereference, represented as an unsafe.Pointer.
//
// Returns:
//   - byte: The value stored at the specified memory address.
//
// NOTE:
//   - Every address points to exactly one byte and one byte of data alone.
func DerefUnsafePointer(address unsafe.Pointer) byte {
	return *(*byte)(address)
}

// PrintMemoryAndData prints the memory addresses and values of a specified memory range.
// It starts from the given memory address and continues for the number of bytes specified
// by systemAllocatedBytes. For each byte, the function prints the memory address, the
// hexadecimal value, and the binary value of the byte.
//
// Parameters:
//   - firstByteAddress: A pointer to the first byte of the memory range to be printed.
//   - systemAllocatedBytes: The number of bytes to be printed from the starting address.
//
// Output:
// The function prints a table with three columns:
//   - Memory Address: The address of the current byte in memory.
//   - Hex Value: The value of the byte in hexadecimal format.
//   - Binary Value: The value of the byte in binary format.
//
// Example output:
// Memory Address   | Hex Value | Binary Value
// -----------------|-----------|-------------
// 0x12345678       | 0x1f      | 00011111
// 0x12345679       | 0xa0      | 10100000
// ...
func PrintMemoryAndData(firstByteAddress unsafe.Pointer, systemAllocatedBytes uintptr) {
	fmt.Println("Memory Address   | Hex Value | Binary Value")
	fmt.Println("-----------------|-----------|-------------")
	for i := uintptr(0); i < systemAllocatedBytes; i++ {
		address := (*byte)(unsafe.Pointer(uintptr(firstByteAddress) + i))
		byteValue := *address
		fmt.Printf("%p  |   0x%02x    |  %08b\n", address, byteValue, byteValue)
	}
}

// PrintIntMemoryAndData prints the memory addresses and values of the memory
// allocated to store an integer value. It determines the size of the integer
// based on the architecture of the machine (32-bit or 64-bit) and prints the
// memory addresses, hexadecimal values, and binary values of each byte.
//
// Parameters:
//   - value: The integer whose memory representation will be printed.
//
// Output:
// The function prints a table with three columns:
//   - Memory Address: The address of the current byte in memory.
//   - Hex Value: The value of the byte in hexadecimal format.
//   - Binary Value: The value of the byte in binary format.
//
// Example output:
// Memory Address   | Hex Value | Binary Value
// -----------------|-----------|-------------
// 0x12345678       | 0x1f      | 00011111
// 0x12345679       | 0xa0      | 10100000
// ...
func PrintIntMemoryAndData(value int) {
	// Determine the number of bytes allocated for the integer
	// On 64-bit machines this will be 8 bytes
	// On 32-bit machines it will be 4 bytes
	overralBytesAllocated := unsafe.Sizeof(value)

	// Get the address of the first byte of the integer
	addressOfFirstByte := unsafe.Pointer(&value)

	// Print the memory data starting from the first byte address
	PrintMemoryAndData(addressOfFirstByte, overralBytesAllocated)
}
