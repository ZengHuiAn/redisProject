package pack

import (
	"testing"
)

func TestPack_bool_data(t *testing.T) {
	var arr = make([]interface{}, 9)
	arr[0] = 1;
	arr[1] = nil;
	arr[2] = true;
	arr[3] = false;
	arr[4] = 3.5;
	arr[5] = "lskjdfksff";
	arr[6] = []byte{1,2,3,4,5};
	arr[7] = []interface{}{3, "xx"};
	arr[8] = []int{3, 4, 5};

	var test = Pack_common(arr)
	t.Log(test)
	t.Log( UNPack_common(test))
}
