package chouxiang

import "fmt"

func Chouxiang() {
	fbFactory, _ := GetSportsFactory("fb")
	xhsFactory, _ := GetSportsFactory("xhs")

	fb := fbFactory.Push()
	xhs := xhsFactory.Push()

	printDetails(fb)
	printDetails(xhs)

}

func printDetails(p Push) {
	fmt.Println(p.msg)
	fmt.Println(p.error)
}
