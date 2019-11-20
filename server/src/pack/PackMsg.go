package pack

func Pack_bool_data(value bool) []byte {
	if value == true {
		return []byte{1}
	}

	return []byte{0}
}

func Pack_rune_data(value rune) []byte {

	return []byte{0}
}

func init() {
	Pack_bool_data(true)
}
