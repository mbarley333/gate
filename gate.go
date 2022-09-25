package gate

func AND(a, b bool) bool {
	result := false

	if a == b && a {
		result = true
	}

	return result
}

func OR(a, b bool) bool {
	result := false

	if a != b {
		result = true
	} else if a && b {
		result = true
	}

	return result
}

func NOT(a bool) bool {
	result := false

	if !a {
		result = true
	}

	return result
}

func NAND(a, b bool) bool {
	result := AND(a, b)

	return NOT(result)
}

func NOR(a, b bool) bool {
	result := OR(a, b)

	return NOT(result)
}

func XOR(a, b bool) bool {
	result := false

	if a != b {
		result = true
	}

	return result
}
