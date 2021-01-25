package main

import (
	"fmt"
	"strings"
)

type LiarInt int

func (i LiarInt) Get() int {
	res := int(i) + 5
	return res
}

type Band struct {
	member	int
	name	string
	vocal	string
	guitar	string
	bass	string
	drummer	string
}

type PriceInfo struct {
	OpenPrice	int
	ClosePrice	int
}

func (b Band) Member() int {
	return b.member
}

func (b Band) Name() string {
	return b.name
}

func (b Band) Vocal() string {
	return b.vocal
}

func (b Band) Guitar() string {
	i := strings.Index(b.guitar, ";")
	if i == -1 {
		return b.guitar
	}
	guitars := strings.Split(b.guitar, ";")

	var guitar_str string
	for i, _ := range guitars {
		guitar_str += guitars[i]
		if i != len(guitars)-1 {
			guitar_str += " & "
		}
	}
	return guitar_str
}

func (b Band) Bass() string {
	return b.bass
}

func (b Band) Drummer() string {
	return b.drummer
}

func (p *Band) AddMember(num int) {
	p.member += num
}

// 不使用指针，这里是值传递，实际上并没有对它进行修改
func (b Band) ChangeName_Err(name string) {
	b.name = name
}

func (b *Band) ChangeName(name string) {
	b.name = name
}

// Band作为匿名字段被嵌套在Show中
type Show struct {
	Time	string
	Arena	string
	Song	string
	Band
}

func main() {
	var a LiarInt = 5
	fmt.Printf("[a = %d]\n", a.Get())

	p_band := &Band{
		member:  5,
		name:    "SoBand",
		vocal:   "AShin",
		guitar:  "Monster;Stone",
		bass:    "Matthew",
		drummer: "Ming",
	}
	var band Band = *p_band
	fmt.Printf("[Band: %s] [Member: %d]\n", band.Name(), band.Member())
	fmt.Printf("[Vocal: %s] [Guitar: %s] [Bass: %s] [Drummer: %s]\n", band.Vocal(), band.Guitar(), band.Bass(), band.Drummer())
	fmt.Println("=============================================================")

	band.ChangeName_Err("MAYDAY")
	fmt.Printf("[Band: %s] [Member: %d]\n", band.Name(), band.Member())
	fmt.Printf("[Vocal: %s] [Guitar: %s] [Bass: %s] [Drummer: %s]\n", band.Vocal(), band.Guitar(), band.Bass(), band.Drummer())
	fmt.Println("=============================================================")

	band.ChangeName("MAYDAY")
	fmt.Printf("[Band: %s] [Member: %d]\n", band.Name(), band.Member())
	fmt.Printf("[Vocal: %s] [Guitar: %s] [Bass: %s] [Drummer: %s]\n", band.Vocal(), band.Guitar(), band.Bass(), band.Drummer())

	show := &Show{"18:00", "Bird's Nest", "Embrace", band}
	bandName := show.Name()                 // Show的对象直接调用其匿名字段Band的方法
	fmt.Printf("[Shower: %s]\n", bandName)	// 输出结果为: MAYDAY
}

func Test_StringToStruct() {


}