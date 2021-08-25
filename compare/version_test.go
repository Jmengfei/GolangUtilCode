package compare

import (
	"fmt"
	"testing"
)

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/24 3:39 下午
 * @Desc:
 */
func TestVersionOrdinal(t *testing.T) {
	versions := []struct{ a, b string }{
		//{"1.05.00.0156", "1.0.221.9289"},
		//// Go versions
		//{"1", "1.0.1"},
		//{"1.0.1", "1.0.2"},
		//{"1.0.2", "1.0.3"},
		//{"1.0.3", "1.1"},
		//{"1.1", "1.1.1"},
		//{"1.1.1", "1.1.2"},
		//{"1.1.2", "1.2"},
		//{"1.8.8", "1.8.13"},
		//{"1.8.8a", "1.8.8a"},
		{"1.8.8a", "1.8.8b"},
	}
	for _, version := range versions {
		a, b := NewVersion(version.a), NewVersion(version.b)
		t.Logf("a=[%+v], b=[%+v]", a, b)

		if a.LessThan(b) {
			fmt.Println(version.a, "<", version.b)
		}else if a.GreaterThan(b) {
			fmt.Println(version.a, ">", version.b)
		}else {
			fmt.Println(version.a, "=", version.b)
		}

		a1, b1 := versionOrdinal(version.a), versionOrdinal(version.b)
		switch {
		case a1 > b1:
			fmt.Println(version.a, ">", version.b)
		case a1 < b1:
			fmt.Println(version.a, "<", version.b)
		case a1 == b1:
			fmt.Println(version.a, "=", version.b)
		}

		fmt.Println("==========================")
	}
}