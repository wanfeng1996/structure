package pkg

const (
	A     = "A"     //=     Source18
	AA    = "AA"    //=   Source21
	AAA   = "AAA"   //=  Source25
	AAAA  = "AAAA"  //=  Source28
	AAAAA = "AAAAA" //= Source30
	A1    = "1A"    //= Source18
	A2    = "2A"    //= Source21
	A3    = "3A"    //= Source25
	A4    = "4A"    //= Source28
	A5    = "5A"    //= Source30
)

var ScenicSpotLevel = map[string]int{
	A1:    Source18,
	A:     Source18,
	A2:    Source21,
	AA:    Source21,
	A3:    Source25,
	AAA:   Source25,
	A4:    Source28,
	AAAA:  Source28,
	A5:    Source30,
	AAAAA: Source30,
}

// GradeOfScenicSpotScore 景区等级 （1A:18,2A:21,3A:25,4A:28,5A:30）
func GradeOfScenicSpotScore(Level string) int {
	if value, ok := ScenicSpotLevel[Level]; ok {
		return value
	}
	return 0
}
