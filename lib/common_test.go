package lib

import "testing"

func TestMd5(t *testing.T) {
	infos := []struct {
		str       string
		want      string
		encodePwd string
	}{
		{
			str:       "123456",
			want:      "e10adc3949ba59abbe56e057f20f883e",
			encodePwd: "7426a9906192b16ea57a89063c95935c",
		},
	}

	for _, item := range infos {
		md5Value := Md5(item.str)
		encodePwd := EncodePwd(item.str)
		if md5Value != item.want {
			t.Fatal(md5Value)
		}
		if encodePwd != item.encodePwd {
			t.Fatal("encode pwd error :", encodePwd)
		}
	}
}
