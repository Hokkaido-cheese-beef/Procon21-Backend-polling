package pkg

import "log"

//1:快適　2:二酸化炭素濃度が高い
func CheckCo2Level(co2 int)int{

	var level int
	if co2>800{
		level = 1
	}else{
		level=2
	}
	return level
}

//不快指数を参考
//変更　1:寒い 2:普通 3:暑い　4:逃げろ
func CheckComfortLevel(temp, hum float64)int{

	comfortCalc := 0.81*temp+0.01*hum*(0.99*temp-14.3)+46.3

	log.Println(comfortCalc)

	comfort := int(comfortCalc)

	var level int
	switch {
	case comfort<60:
		level=1
	case comfort>=60 && comfort <= 74:
		level =2
	case comfort>=75 && comfort<=79:
		level=3
	case comfort>=80:
		level=4
	}

	return level
}

