package iso

import (
	"testing"
)

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/27 2:33 下午
 * @Desc:
 */

func TestGetISOMapFromFile(t *testing.T) {
	GetISOMapFromFile()
}

func TestGetCountryName(t *testing.T) {
	type data struct {
		iso string
		want string
	}

	cases := map[string]data {
		"one":{iso: "BI", want: "Burundi"},
		"two":{iso: "BJ", want: "Benin"},
		"three":{iso: "AA", want: "AA"},
		"four":{iso: "BN", want: "Brunei Darussalam"},
	}
	for name,d := range cases{
		t.Run(name, func(t *testing.T) {
			got := GetCountryName(d.iso)
			if got != d.want {
				t.Errorf("name:%s excepted:%#v, got:%#v", name, d.want, got)
			}
		})
	}

}