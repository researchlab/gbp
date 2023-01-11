package utils

import (
	"testing"
)

func TestNewDes(t *testing.T) {
	tests := []struct {
		data   []byte
		wanted string
	}{
		{data: []byte("china.com"), wanted: "QeuE/iRPxiV+CZZhAL54dQ=="},
	}
	for _, test := range tests {
		er := DesECBEncrypt(test.data, KEY)
		dr := DesECBDecrypter(er.data, KEY)
		if er.Base64String() != test.wanted {
			t.Errorf("encrypt invalid, got:%s, wanted:%s\n", er.Base64String(), test.wanted)
		}
		if string(dr.Bytes()) != string(test.data) {
			t.Errorf("decrypter invalid, got:%s, wanted:%s\n", dr.Bytes(), test.data)
		}
	}
}
