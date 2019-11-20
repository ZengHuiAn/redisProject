package pack

func Pack_bool_data(value bool) []byte {
	if value == true {
		return []byte{1}
	}

	return []byte{0}
}

// pack byte to do


func init() {
	Pack_bool_data(true)
}
