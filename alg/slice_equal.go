package alg

// StringSliceEqual 判断两个字符串切片是否相等
func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ShouldSummerBeComing(actual interface{}, expected ...interface{}) string {
	if actual == "summer" && expected[0] == "coming" {
		return ""
	} else {
		return "summer is not coming!"
	}
}

