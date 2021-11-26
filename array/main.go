package main

import (
	"fmt"
)

func main() {
	testname("array")
	{
		coast := [5]string{"zushi", "yuigahama", "inamura", "shichiri"}
		coast2 := [...]string{"zushi", "yuigahama", "inamura", "shichiri"}
		fmt.Println("Coast :", coast)
		fmt.Println("size  :", len(coast))
		fmt.Println("Coast2:", coast2)
		fmt.Println("size2 :", len(coast2))
		numbers := [...]int{99: -1} // specify -> numbets[99] = -1
		fmt.Println("Numbers :", numbers)
	}
	testname("multiple array")
	{
		var ThreeD [2][3][2]int
		fmt.Println("init: ", ThreeD)
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				for k := 0; k < 2; k++ {
					ThreeD[i][j][k] = (i + 1) * (j + 1) * (k + 1)
				}
			}
		}
		fmt.Println("after loop: ", ThreeD)
	}
	testname("slice")
	{
		months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
		quarter2 := months[3:6]
		quarter2Extended := quarter2[:5]
		quarter2ToEnd := months[3:]
		fmt.Println(months, len(months), cap(months))
		fmt.Println(quarter2, len(quarter2), cap(quarter2))
		fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended))
		fmt.Println(quarter2ToEnd, len(quarter2ToEnd), cap(quarter2ToEnd))
		months = append(months, "messenger")
		fmt.Println(months, len(months), cap(months))

		remove := 2
		copy_months := make([]string, 13)
		copy(copy_months, months)
		if remove < len(copy_months) {
			fmt.Println("copy months:Before", copy_months, "Remove", copy_months[remove])
			copy_months = append(copy_months[:remove], copy_months[remove+1:]...)
			fmt.Println("copy months:After ", copy_months)
		}
	}
}

func testname(s string) {
	fmt.Println("---------------[", s, "]")
}