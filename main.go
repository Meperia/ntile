package ntile

func calcBucketCount(lenXS int, bucket int) int {
	c := lenXS / bucket
	if lenXS%bucket != 0 || c == 0 {
		c += 1
	}
	return c
}

func Ntile(xs []interface{}, bucket int) [][]interface{} {
	ret := make([][]interface{}, bucket)
	bucketIndex := 0
	i := 0
	lenXS := len(xs)
	bucketCount := calcBucketCount(lenXS, bucket)
	for _, x := range xs {
		if i > 0 && i%bucketCount == 0 {
			bucketIndex += 1
			bucket -= 1
			bucketCount = calcBucketCount(lenXS-i, bucket)
			i = 0
		}
		ret[bucketIndex] = append(ret[bucketIndex], x)
		i += 1
	}

	return ret
}

func NtileAt(xs []interface{}, bucket int, bucketIndex int) []interface{} {
	var ret []interface{}
	bi := 0
	i := 0
	lenXS := len(xs)
	bucketCount := calcBucketCount(lenXS, bucket)
	for _, x := range xs {
		if bi > bucketIndex {
			return ret
		}
		if i > 0 && i%bucketCount == 0 {
			bi += 1
			bucket -= 1
			bucketCount = calcBucketCount(lenXS-i, bucket)
			i = 0
		}
		if bi == bucketIndex {
			ret = append(ret, x)
		}
		i += 1
	}

	return ret
}
