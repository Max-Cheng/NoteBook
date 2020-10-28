// 6kyu	Most Frequent Weekdays
package main

import (
	"fmt"
)

func main() {
	fmt.Println(MostFrequentDays(1990))
}

func MostFrequentDays(year int) []string {
	weekdaymap:=map[int]string{0:,1:,2:,3:,4:,5:,6:}
	centry := (year / 100)
	//first
	y := year%100 - 1
	w := ((y + (y / 4) + (centry / 4) - 2*centry) + (26 * (13 + 1) / 10)) % 7
	//end
	newcentry := ((year + 1) / 100)
	ny := y + 1
	nw := ((ny + (ny / 4) + (newcentry / 4) - 2*newcentry) + (26 * (12 + 1) / 10) + 30) % 7

	// TODO:date&&weekday map
	// if w<nw{
	// 	return []strings{weekdaymap[w],weekdaymap[w+1]}
	// }else{
	// 	if w>nw {
	// 		return []strings{weekdaymap[w-1],weekdaymap[w]}
	// 	}
	// }
	// return []strings{weekdaymap[w-1],weekdaymap[w]}
}
