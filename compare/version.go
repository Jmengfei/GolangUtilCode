package compare

/**
 * @Author: mf
 * @email: 18539271635@163.com
 * @Date: 2021/8/24 3:26 下午
 * @Desc:
 */

type Version struct {
	original string
}

func NewVersion(v string) *Version {
	return &Version{
		versionOrdinal(v),
	}
}

func (v *Version) LessThan (o *Version) bool {
	return v.original < o.original
}

func (v *Version) LessThanOrEqual(o *Version) bool {
	return v.original <= o.original
}

func (v *Version) GreaterThan (o *Version) bool {
	return v.original > o.original
}

func (v *Version) GreaterThanOrEqual(o *Version) bool {
	return v.original >= o.original
}

func versionOrdinal(version string) string {
	// ISO/IEC 14651:2011
	const maxByte = 1<<8 - 1

	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0x00)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			panic("VersionOrdinal: invalid version")
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}