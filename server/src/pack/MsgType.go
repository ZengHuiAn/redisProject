package pack

type EPackType byte
const (
	UNDEFINED      EPackType = 0
	NULL EPackType= 100
	BOOL EPackType= 101
	CHAR EPackType= 102
	BYTE EPackType= 103
	INT16 EPackType= 104
	UINT16 EPackType = 105
	INT32 EPackType= 106
	UINT32 EPackType = 107
	INT64 EPackType= 108
	UINT64 EPackType= 109
	SINGLE EPackType= 110
	DOUBLE EPackType= 111
	STRING EPackType= 112
	BYTEARRAY = 113
	ARRAY = 114
)