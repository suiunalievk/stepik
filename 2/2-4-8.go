package main

type testStruct struct {
	On          bool
	Ammo, Power int
}

func (x *testStruct) Shoot() bool {
	if !x.On || x.Ammo == 0 {
		// x.structOnOff(0)
		return false
	}

	x.Ammo--
	return true
}

func (x *testStruct) RideBike() bool {
	if !x.On || x.Power == 0 {
		// x.structOnOff(0)
		return false
	}

	x.Power--
	return true
}

func (x *testStruct) structOnOff(mode int) {
	switch mode {
	case 1:
		x.On = true
	case 0:
		x.On = false
	}
}

func main() {
	testStruct := new(testStruct)

	// test := testStruct{true, 5, 10}
	// fmt.Printf("Struct init values are On=%t, Ammo=%d, Power=%d\n", testStruct.On, testStruct.Ammo, testStruct.Power)
	// for i := 0; i <= 10; i++ {
	// 	testStruct.Shoot()
	// 	fmt.Printf("Ammo = %d\n", testStruct.Ammo)

	// 	testStruct.RideBike()
	// 	fmt.Printf("Power = %d\n", testStruct.Power)

	// 	if !testStruct.On {
	// 		fmt.Println("Struct is off!")
	// 		break
	// 	}
	// }
	// fmt.Printf("Struct after loop values are On=%t, Ammo=%d, Power=%d\n", testStruct.On, testStruct.Ammo, testStruct.Power)
}
