package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//salary := make(map[string]int)
	//if salary != nil {
	//	fmt.Println("salary is not nil")					//此句被执行，nil并不代码map元素是否为空
	//}
	//
	//var salary2 map[string]int
	//if salary2 == nil {
	//	fmt.Println("[salary2 is nil, please Make it]")		//map没有被创建，为nil map
	//	salary2 = make(map[string]int)
	//}
	//
	//salary3 := map[string]int {
	//	"Iverson" : 2400,
	//	"kobe" : 2280,
	//	"TMAC" : 2300,
	//}
	//fmt.Println(salary3)
	//
	////插入数据及修改数据
	//salary["James"] = 1800	//插入
	//salary["James"] = 1900	//修改
	//salary["Wade"] = 1750
	//fmt.Println(salary)
	//fmt.Println("Salary of James is", salary["James"])			//访问元素
	//fmt.Println("Salary of Iverson is", salary["Iverson"])		//访问一个不存在的元素，返回value类型的零值
	//
	////判断一个key值是否存在
	//value, ok := salary["Curry"]
	//if ok == true {
	//	fmt.Println("Salary of Curry is", value)
	//}else {
	//	fmt.Println("Curry doesn't exist, value is", value)
	//}
	//
	////遍历
	//for k, v := range salary {
	//	fmt.Printf("salary[%s] = %d\n", k, v)
	//}
	//
	////删除元素
	//delete(salary, "Wade")        //删除一个存在的元素
	//delete(salary, "Iverson")     //删除一个不存在的元素
	//fmt.Println(salary)
	//
	////元素个数
	//fmt.Println("Length of salary is", len(salary))
	//
	////map是引用类型，赋值给一个新变量，对其操作会影响原变量
	//salary["Curry"] = 500
	//new_salary := salary
	//new_salary["Curry"] = 540
	//fmt.Println(salary)			//Curry的结果为540，因为对new_salary操作也会影响原mapsalary

	//var mMap map[string]map[string]string
	//mMap = make(map[string]map[string]string)
	//if mMap["EUR"] == nil {
	//	mMap["EUR"] = make(map[string]string)
	//}
	//if mMap["USD"] == nil {
	//	mMap["USD"] = make(map[string]string)
	//}
	//mMap["EUR"]["USD"] = "EURUSD"
	//mMap["USD"]["CNH"] = "USDCNH"
	//for k1, mMap := range mMap {
	//	for k2, v := range mMap {
	//		fmt.Printf("symbol[%s][%s] = %s\n", k1, k2, v)
	//	}
	//}



	ConcurrentMapTest()
}

type Person struct {
	name	string
	age		int
}

type PersonMap struct {
	mPerson	map[string]*Person
	sync.RWMutex
}

func (p *PersonMap) Get(key string) *Person {
	p.RLock()
	if p.mPerson[key] == nil {
		return nil
	}
	pi := *p.mPerson[key]
	p.RUnlock()
	return &pi
}

func (p *PersonMap) Set(key, name string, age int) {
	p.Lock()
	if p.mPerson == nil {
		p.mPerson = make(map[string]*Person)
	}
	if p.mPerson[key] == nil {
		p.mPerson[key] = &Person{}
	}
	p.mPerson[key].name = name
	p.mPerson[key].age = age
	p.Unlock()
}

var (
	pm PersonMap
)
func ConcurrentMapTest() {
	pm.Set("ChosenOne", "James", 36)

	// Read-Thread
	go func() {
		p := pm.Get("ChosenOne2")
		if p == nil {
			fmt.Println("Nil Pointer")
			return
		}
		fmt.Printf("[Key: ChosenOne] [Name: %s] [Age: %d]\n", p.name, p.age)
		time.Sleep(time.Second * 2)
		fmt.Printf("[Key: ChosenOne] [Name: %s] [Age: %d]\n", p.name, p.age)
	}()

	// Write-Thread
	go func() {
		pm.Set("ChosenOne", "NilChow", 27)
	}()

	time.Sleep(time.Second * 5)
}