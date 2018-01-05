package turtl

import (
	"fmt"
	"unsafe"
)

const (
	major uint8 = 0
	minor uint8 = 1
	patch uint8 = 0
)

var versionNumStr = fmt.Sprintf("%d.%d.%d", major, minor, patch)
var intSize = unsafe.Sizeof(uint(0)) * 8 // number of bits in an uint

/*
VersionNum returns the string representation of the
version number: Major.Minor.Patch
*/
func VersionNum() string {
	return versionNumStr
}

// Major version number
func Major() uint8 {
	return major
}

// Minor version number
func Minor() uint8 {
	return minor
}

// Patch version number
func Patch() uint8 {
	return patch
}

/*
GitHash returns the git commit hash for the version
of turtl that was compiled
*/
func GitHash() string {
	return gitHash
}

/*
IntSize returns the size of an unsigned int on this
platform
*/
func IntSize() uintptr {
	return intSize
}

/*
InfoGraphic returns a welcome message with some system
info and an ascii turtle.
*/
func InfoGraphic() string {
	return fmt.Sprintf(
		`

    .=.  ____  .=.
    \ .-'    '-. /       turtl  v%s  - %d bit
    /.'\_/\_/'.\.-p.     %s
--=|: -<_><_>- :|   >
    \'./ \/ \.'/'-b'
    / '-.____.-' \
    '='        '='

`, versionNumStr, intSize, gitHash[:13])
}
