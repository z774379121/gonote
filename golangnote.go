//package main
//
//import (
//	"time"
//	"context"
//	"fmt"
//)

//import (
//	"sync"
//	"math/rand"
//	"time"
//	"fmt"
//)
//
//var wg sync.WaitGroup
//
//func main() {
//	rand.Seed(time.Now().UnixNano())
//	wg.Add(1)
//	score := make(chan int)
//	go run(score)
//	score <- 1
//	wg.Wait()
//
//}
//
//func run(score chan int) {
//	var newruner int
//	newruner = <- score
//	fmt.Println(newruner, "i caush it")
//
//	fmt.Println("begin run----", newruner)
//	//创建下一位
//	if newruner != 4 {
//		newruner ++
//		fmt.Println( newruner, " ready")
//		go run(score)
//	}
//	time.Sleep(time.Second * 2)
//	if newruner == 4 {
//		wg.Done()
//		fmt.Println("match over")
//		return
//	}
//	score <- newruner
//}

// context包
//package main
//
//import (
//"fmt"
//"time"
//	"context"
//)
//
//// 模拟一个最小执行时间的阻塞函数
//func inc(a int) int {
//	res := a + 1                // 虽然我只做了一次简单的 +1 的运算,
//	time.Sleep(1 * time.Second) // 但是由于我的机器指令集中没有这条指令,
//	// 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
//	return res
//}
//
//// 向外部提供的阻塞接口
//// 计算 a + b, 注意 a, b 均不能为负
//// 如果计算被中断, 则返回 -1
//func Add(ctx context.Context, a, b int) int {
//	res := 0
//	for i := 0; i < a; i++ {
//		res = inc(res)
//		select {
//		case <-ctx.Done():
//			return -1
//		default:
//		}
//	}
//	for i := 0; i < b; i++ {
//		res = inc(res)
//		select {
//		case <-ctx.Done():
//			return -1
//		default:
//		}
//	}
//	return res
//}
//
//func main() {
//	{
//		// 使用开放的 API 计算 a+b
//		a := 1
//		b := 2
//		timeout := 2 * time.Second
//		ctx, _ := context.WithTimeout(context.Background(), timeout)
//		res := Add(ctx, 1, 2)
//		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
//	}
//	{
//		// 手动取消
//		a := 1
//		b := 2
//		ctx, cancel := context.WithCancel(context.Background())
//		go func() {
//			time.Sleep(2 * time.Second)
//			cancel() // 在调用处主动取消
//		}()
//		res := Add(ctx, 1, 2)
//		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
//	}
//}
package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"unsafe"
)

type Handler interface {
	Do(k, v interface{})
}

type HandlerFunc func(k, v interface{})

func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}
func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26
	EachFunc(persons, selfInfo)
	c := Coupons{}
	c.Name = "!"
	c.Belongto = bson.NewObjectId()
	c.Descrition = "!!"
	c.Price = 10
	fmt.Printf("%+v", c)
	match := DoWork(&c, TranseMatch)
	order := DoWork(&c, TranseOrder)
	fmt.Println(match, order)
	k := Additon{"@", bson.NewObjectId()}
	testi(&k)
	mi := Mi(20)
	mi.Show()
	ei := Easytype{2, 3}
	fmt.Printf("%p", &ei)
	getaddress(ei)
	fmt.Println(ei)
	kk := []Additon{
		{"2", bson.NewObjectId()},
		{"3", bson.NewObjectId()},
		{"3", bson.NewObjectId()},
		{"8", bson.NewObjectId()},
		{"9", bson.NewObjectId()},
		{"7", bson.NewObjectId()},
	}
	d := Slice(kk, reflect.TypeOf([]Additoni{}))
	fmt.Println(kk)
	fmt.Println(d.([]Additoni))

}

type Coupons struct {
	Id_   bson.ObjectId `bson:"id_"`
	Name  string        `bson:"name" json:"name"`
	Price int           `bson:"price" json:"price"`
	Additon
}

type Testpoint interface {
	Do()
	DoAgain()
}

func testi(testpoint Testpoint) {
	fmt.Println("1")
}

//方法集为指针接受者，实现对象必须是指针对象
// a:= Additon{} testi(a)报错
// a:= &Additon{} testi(a)通过
func (this *Additon) Do() {

}

func (this *Additon) DoAgain() {

}

//方法集的接受者为值时，实现对象可以是值对象也可以是指针对象

//也就是说该对象为指针对象时可以实现 以指针和值为接受者的方法集接口
//而值对象只能实现以值为接受者的方法集接口

// a:= Additon{} testi(a)通过
// a:= &Additon{} testi(a)通过

//func (this Additon) Do()  {
//
//}
//
//func (this Additon) DoAgain()  {
//
//}

// 强制切片类型转换
type Additon struct {
	Descrition string
	Belongto   bson.ObjectId
}

type Additoni struct {
	Descrition string
	Belongto   bson.ObjectId
}

type Mi int

func (this *Mi) Show() {
	fmt.Println("show")
}

type Easytype struct {
	I1 int
	I2 int
}

func getaddress(easytype Easytype) {
	easytype.I1 = 20
	easytype.I2 = 30
	fmt.Println(easytype)
}

// 面向接口编程
type TranseSomething interface {
	Transefiled(input interface{}) bson.M
}

type MyTranse func(input interface{}) bson.M

func (this MyTranse) Transefiled(input interface{}) bson.M {
	return this(input)
}

func TranseMatch(input interface{}) bson.M {
	t := reflect.TypeOf(input).Elem()
	v := reflect.ValueOf(input).Elem()
	results := make(bson.M, 10)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("bson") != "" {
			results[t.Field(i).Tag.Get("bson")] = v.Field(i).Interface()
		}

	}
	return results
}

func TranseOrder(input interface{}) bson.M {
	t := reflect.TypeOf(input).Elem()
	v := reflect.ValueOf(input).Elem()
	results := make(bson.M, 10)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") != "" {
			results[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
		}
	}
	return results
}

func RealDoWork(input interface{}, transe TranseSomething) bson.M {
	return transe.Transefiled(input)
}

func DoWork(input interface{}, method func(input interface{}) bson.M) bson.M {
	return RealDoWork(input, MyTranse(method))
}

// 	----------------------------
// 强制切片转换
func Slice(slice interface{}, newSliceType reflect.Type) interface{} {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice value of type %T", slice))
	}
	if newSliceType.Kind() != reflect.Slice {
		panic(fmt.Sprintf("Slice called with non-slice type of type %T", newSliceType))
	}
	newSlice := reflect.New(newSliceType)
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(newSlice.Pointer()))
	hdr.Cap = sv.Cap() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Len = sv.Len() * int(sv.Type().Elem().Size()) / int(newSliceType.Elem().Size())
	hdr.Data = uintptr(sv.Pointer())
	return newSlice.Elem().Interface()
}

//其他类型数组转换为[]byte切片
func ByteSlice(slice interface{}) (data []byte) {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("ByteSlice called with non-slice value of type %T", slice))
	}
	h := (*reflect.SliceHeader)((unsafe.Pointer(&data)))
	h.Cap = sv.Cap() * int(sv.Type().Elem().Size())
	h.Len = sv.Len() * int(sv.Type().Elem().Size())
	h.Data = sv.Pointer()
	return
}
