package util

import "time"

func DropUnwantedField() {

}

//TimrToEpoch n in second
func TimeToEpoch(n int64) int64 {
	now := time.Now()
	now = now.Add(time.Duration(n) * (time.Second))
	return now.Unix()
}
