package pkg

func CreateMessage(comfortLevel,co2Level int)string{
	message := ""
	switch {
	case comfortLevel==1 && co2Level==1:
		message += "部屋が寒すぎます！暖かくしましょう！"
	case comfortLevel==1 && co2Level==2:
		message += "部屋は寒く、CO2が高い最悪の環境です！\n部屋を温めて換気も行いましょう！"
	case comfortLevel==2 && co2Level==2:
		message += "CO2濃度が高い状態が続いています。\n一度換気してみてはいかがでしょう？"
	case comfortLevel==3 && co2Level==1:
		message += "汗をかきやすい環境です！\n熱中症に気をつけましょう！"
	case comfortLevel==3 && co2Level==2:
		message +="CO2濃度が高い状態が続いています。\n一度換気してみてはいかがでしょう？"
	case comfortLevel==4 && co2Level==1:
		message += "汗をかきやすい環境です！\n熱中症に気をつけましょう！"
	case comfortLevel==4 && co2Level==2:
		message += "大変まずい環境です。クーラや扇風機で少しでも涼しくしつつ、換気も行いましょう！"
	}
	return message
}
