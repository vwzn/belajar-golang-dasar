package main

import "fmt"

func main() {
	names := [...]string{"ahmad", "fauzan", "abdurrohman", "shyamil", "arraniya", "diandra"}
	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println("1", daySlice1) // [sabtu , minggu]

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println("2", daySlice1) // [sabtu baru, minggu baru]
	fmt.Println("3", days)      // [senin, selasa, rabu, kamis, jumat, sabtu baru, minggu baru]

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "sabtu lama"
	fmt.Println("4", daySlice1) // [sabtu baru, minggu baru]
	fmt.Println("5", daySlice2) // [sabtu lama, minggu baru, libur baru]
	fmt.Println("6", days)      // [senin, selasa, rabu, kamis, jumat, sabtu baru, minggu baru]

	// make slice
	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Ahmad"
	newSlice[1] = "Ahmad"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Fauzan")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
