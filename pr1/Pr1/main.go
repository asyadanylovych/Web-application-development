package main

import (
	"fmt"
)

func main() {
	fmt.Println("------- Завдання 1 --------")

	// Вихідні дані
	hp := 1.4
	cp := 71.7
	sp := 1.8
	np := 0.8
	op := 1.4
	wp := 6.0
	ap := 16.9

	// 1. Коефіцієнти переходу (Таблиця 1.1)
	kpc := 100 / (100 - wp)      // до сухої маси
	kpg := 100 / (100 - wp - ap) // до горючої маси
	fmt.Printf("KPC (до сухої маси) = %.4f\n", kpc)
	fmt.Printf("KPG (до горючої маси) = %.4f\n\n", kpg)

	// 2. Склад сухої маси
	hc, cc, sc, nc, oc, ac := hp*kpc, cp*kpc, sp*kpc, np*kpc, op*kpc, ap*kpc
	sumC := hc + cc + sc + nc + oc + ac

	fmt.Println("СКЛАД СУХОЇ МАСИ:")
	fmt.Printf("H = %.3f%%, C = %.3f%%, S = %.3f%%, N = %.3f%%, O = %.3f%%, A = %.3f%%\n", hc, cc, sc, nc, oc, ac)
	fmt.Printf("Сума = %.2f%%\n\n", sumC)

	// 3. Склад горючої маси
	hg, cg, sg, ng, og := hp*kpg, cp*kpg, sp*kpg, np*kpg, op*kpg
	sumG := hg + cg + sg + ng + og

	fmt.Println("СКЛАД ГОРЮЧОЇ МАСИ:")
	fmt.Printf("H = %.3f%%, C = %.3f%%, S = %.3f%%, N = %.3f%%, O = %.3f%%\n", hg, cg, sg, ng, og)
	fmt.Printf("Сума = %.2f%%\n\n", sumG)

	// 4. Нижча теплота згоряння (Формула Менделєєва 1.2)
	qhp := 339*cp + 1030*hp - 108.8*(op-sp) - 25*wp // кДж/кг
	qhpMJ := qhp / 1000                             // МДж/кг

	// 5. Теплота для різних мас (Таблиця 1.2)
	qhd := (qhpMJ + 0.025*wp) * (100 / (100 - wp))
	qhdaf := (qhpMJ + 0.025*wp) * (100 / (100 - wp - ap))

	fmt.Println("ТЕПЛОТА ЗГОРЯННЯ:")
	fmt.Printf("Робоча маса: %.4f МДж/кг\n", qhpMJ)
	fmt.Printf("Суха маса:   %.4f МДж/кг\n", qhd)
	fmt.Printf("Горюча маса: %.4f МДж/кг\n\n", qhdaf)

	fmt.Println("------- Завдання 2 --------")

	// Дані для мазуту (з контрольного прикладу завдання 2)
	cdaf, hdaf, odaf, sdaf := 85.5, 11.2, 0.8, 2.5
	qdaf, wr, ad, vdaf := 40.4, 2.0, 0.15, 333.3

	// Розрахунки (згідно з методикою п. 1.3.1)
	ar := ad * (100 - wr) / 100
	kgr := (100 - wr - ar) / 100

	cr, hr, or, sr := cdaf*kgr, hdaf*kgr, odaf*kgr, sdaf*kgr
	vr := vdaf * (100 - wr) / 100
	sumR := cr + hr + or + sr + ar + wr

	fmt.Println("СКЛАД РОБОЧОЇ МАСИ МАЗУТУ:")
	fmt.Printf("C = %.2f%%, H = %.2f%%, O = %.2f%%, S = %.2f%%, A = %.2f%%, W = %.2f%%\n", cr, hr, or, sr, ar, wr)
	fmt.Printf("V = %.2f мг/кг\n", vr)
	fmt.Printf("Сума компонентів = %.2f%%\n\n", sumR)

	qr := qdaf*kgr - 0.025*wr
	fmt.Println("ТЕПЛОТА ЗГОРЯННЯ МАЗУТУ:")
	fmt.Printf("Горюча маса: %.2f МДж/кг\n", qdaf)
	fmt.Printf("Робоча маса: %.2f МДж/кг\n", qr)
}
