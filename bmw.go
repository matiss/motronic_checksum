package main

import (
	"fmt"
)

// BMW402Validate validates ECU with code ending 402
func BMW402C098Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B15
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW402SW599Validate validates ECU with code ending 402 and chip number ending with 599
func BMW402C599Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B31
		// reg1Store = 0x7B32
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x44 || buf[reg1End-1] != 0xFF || buf[reg1End-2] != 0xFF || buf[reg1End-3] != 0xFF {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW403Validate validates ECU with code ending 403
func BMW403Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x69FE
		reg1End   = 0x797F
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW403C950Validate validates ECU with code ending 403 and chip code ending 950
func BMW403C950Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x8000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x69FE
		reg1End   = 0x799B
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW405Validate validates ECU with code ending 403
func BMW405Validate(buf []byte) error {
	// Validate length
	if len(buf) != 0x10000 {
		return fmt.Errorf("Invalid file length")
	}

	const (
		reg1Start = 0x69FE
		reg1End   = 0x797F
	)

	// Validate Data Region 1
	if buf[reg1Start] != 0x00 || buf[reg1Start+1] != 0x00 || buf[reg1Start+2] != 0x00 {
		return fmt.Errorf("Invalid start of Data Region 1")
	}

	if buf[reg1End] != 0x4F || buf[reg1End-1] != 0x13 || buf[reg1End-2] != 0x22 || buf[reg1End-3] != 0x40 {
		return fmt.Errorf("Invalid end of Data Region 1")
	}

	return nil
}

// BMW402C098Checksum calculates checksum for ECU with code ending 402 and chip ending with 098
func BMW402C098Checksum(buf []byte) uint16 {
	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B15
		reg1Store = 0x7B16
	)

	// Big Endian
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := SimpleSum16bit(0, reg1Start, reg1End, buf)

	fmt.Printf("Old: %X New: %X\n", checksumStored, sum)

	// Patch buffer
	PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)

	return sum
}

// BMW402C599Checksum calculates checksum for ECU with code ending 402 and chip ending with 599
func BMW402C599Checksum(buf []byte) uint16 {
	const (
		reg1Start = 0x6AFE
		reg1End   = 0x7B31
		reg1Store = 0x7B32
	)

	// Big Endian
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := SimpleSum16bit(0, reg1Start, reg1End, buf)

	fmt.Printf("Old: %X New: %X\n", checksumStored, sum)

	// Patch buffer
	PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)

	return sum
}

// BMW403Checksum calculates checksum for ECU with code ending 403
func BMW403Checksum(buf []byte) uint16 {
	const (
		reg1Start = 0x69FE
		reg1End   = 0x797F
		reg1Store = 0x7980
	)

	// Big Endian
	checksumStored := uint16((uint16(buf[reg1Store]) << 8) | uint16(buf[reg1Store+1]))

	// Calculate new checksum
	sum := SimpleSum16bit(0, reg1Start, reg1End, buf)

	fmt.Printf("Old: %X New: %X\n", checksumStored, sum)

	// Patch buffer
	PatchBuffer(reg1Store, []byte{byte(sum >> 8), byte(sum)}, buf)

	return sum
}

// BMW405Checksum calculates checksum for ECU with code ending 405
func BMW405Checksum(buf []byte) uint16 {
	// WIP
	return 0
}
