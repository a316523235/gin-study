package lib

import "testing"

func TestMd5(t *testing.T) {
	infos := []struct {
		str string
		want string
	}{
		{
			str:"123456",
			want:"e10adc3949ba59abbe56e057f20f883e",
		},
		{
			str:"",
			want:"d41d8cd98f00b204e9800998ecf8427e",
		},
	}

	for _, item := range infos {
		md5Value := Md5(item.str)
		if md5Value != item.want {
			t.Fatal(md5Value)
		}
	}
}
