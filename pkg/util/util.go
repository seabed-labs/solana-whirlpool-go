package util

// GetInstructionVersion - View README.md for version history
func GetInstructionVersion(timeUnix int) int {
	if timeUnix >= 1674853004 {
		return 4
	} else if timeUnix >= 1671600241 {
		return 3
	} else if timeUnix >= 1671598594 {
		return 2
	} else if timeUnix >= 1670122365 {
		return 1
	}
	return 0
}
