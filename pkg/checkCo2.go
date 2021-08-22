package pkg

import "log"

//1:快適　2:やや高め　3:いますぐ換気
func CheckCo2Level(co2 int)int{

	var level int
	if co2<800{
		level = 1
		/*
	}else if 800<co2 && co2>1000{
		level=2

		 */
	}else{
		level=2
	}
	return level
}

//不快指数を参考
//1:寒い　2:普通　3:快適　4:普通　5:暑い　6：逃げろ
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

	/*　レベルの数を減らしました
		case comfort>=60 && comfort<=64:
			level=2
		case comfort>=65 && comfort<=69:
			level=3
		case comfort>=70 && comfort<=74:
			level=4
	*/

	return level
}

