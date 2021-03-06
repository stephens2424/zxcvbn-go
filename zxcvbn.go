package zxcvbn

import (
	"github.com/nbutton23/zxcvbn-go/matching"
	"github.com/nbutton23/zxcvbn-go/scoring"
	"time"
	"github.com/nbutton23/zxcvbn-go/utils/math"
)

//func main() {
//	password :="Testaaatyhg890l33t"
//	fmt.Println(PasswordStrength(password, nil))
//}

func PasswordStrength(password string, userInputs []string) scoring.MinEntropyMatch {
	start := time.Now()
	matches := matching.Omnimatch(password, userInputs)
	result := scoring.MinimumEntropyMatchSequence(password, matches)
	end := time.Now()

	calcTime := end.Nanosecond() - start.Nanosecond()
	result.CalcTime = zxcvbn_math.Round(float64(calcTime)*time.Nanosecond.Seconds(), .5, 3)
	return result
}