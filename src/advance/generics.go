package main

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

//func Sum[K comparable, V int64 | float64](m map[K]V) V {
//	var s V
//	for _, v := range m {
//		s += v
//	}
//	return s
//}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	intArray := map[string]int64{
		"1": 11,
		"2": 22,
	}
	println(SumInts(intArray))
	println(SumIntsOrFloats[string, int64](intArray))
	println(SumIntsOrFloats(intArray))
}
