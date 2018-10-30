package main

import (
	"fmt"
	"reflect"
)

func main() {
	// pointer()
	// struct_function()
	// reflect_value()
	// reflect_nonexport()
	// usebind()
	// callFunction()
	// channel()
	selectctrl()
}

func basic() {
	type A = [16]int16

	var c <-chan map[A][]byte

	tc := reflect.TypeOf(c)

	fmt.Println("tc kind", tc.Kind())    // chan
	fmt.Println("tc chan dir", tc.ChanDir()) // <-chan

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	tm := tc.Elem() // 这里elem指的是chan的类型，也就是 map[A][]byte
	ta, tb := tm.Key(), tm.Elem()
	fmt.Printf("the elem of tc %v, the key of elem %v (A [16]int16), the value type of map %v ([]byte) \n", tm.Kind(), ta.Kind(), tb.Kind())

	tx, ty := ta.Elem(), tb.Elem() // tb.Elem() 指的是 []byte 的slice类型，也就是byte uint8

	fmt.Printf("tx (int16): %v, ty: %v \n", tx, ty)

	fmt.Println(tx.Bits(), ty.Bits())

	fmt.Println(tx.ConvertibleTo(ty))

	fmt.Println(ty.ConvertibleTo(tx))
}

type T []interface{
	m()
}
type K struct{}

func (T) m() {}
func (K) m() {}

func pointer() {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})

	fmt.Printf("kind of tp is %v, kind of tt is %v\n", tp.Kind(), tt.Kind())

	// tim is the elem of slice []interface{} which is interface{}
	// ti is the elem of new(interface{}) which is interface{}
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind())

	// tt is type of T{}
	fmt.Println(tt.Implements(tim)) // true
	fmt.Println(reflect.TypeOf(K{}).Implements(tim)) //true
}

type F func(string, int) bool
func (f F) Validate(s string) bool {
	return f(s, 32)
}

func struct_function() {
	var x struct {
		n int
		f F
	}

	typeOfx := reflect.TypeOf(x)

	fmt.Println(typeOfx.Kind()) // struct
	fmt.Println(typeOfx.NumField()) // 2

	fieldf := typeOfx.Field(1)
	typeOfFieldf := fieldf.Type

	fmt.Println(typeOfFieldf.Kind()) // func
	fmt.Println(typeOfFieldf.IsVariadic()) // false
	fmt.Println(typeOfFieldf.NumIn(), typeOfFieldf.NumOut()) // 2 1
	fmt.Println(typeOfFieldf.NumMethod()) // 1
	ts, ti, tb := typeOfFieldf.In(0), typeOfFieldf.In(1), typeOfFieldf.Out(0)
	fmt.Println(ts.Kind(), ti.Kind(), tb.Kind())
}

func create() {

	ta := reflect.ArrayOf(10, reflect.TypeOf(0))
	fmt.Println(ta)

	tc := reflect.ChanOf(reflect.SendDir, ta)
	fmt.Println(tc)

	tp := reflect.PtrTo(ta) // *[10]int
	fmt.Println(tp)

	ts := reflect.SliceOf(reflect.TypeOf(0))
	fmt.Println(ts)

	tm := reflect.MapOf(ta, tc)
	fmt.Println(tm)

	tf := reflect.FuncOf([]reflect.Type{ta, tc}, []reflect.Type{tm}, false)
	fmt.Println(tf)

	tt := reflect.StructOf([]reflect.StructField{
		{
			Name: "Age",
			Type: reflect.TypeOf(""),
		},
	})

	fmt.Println(tt)
	fmt.Println(tt.NumField())
}

func reflect_value(){
	n := 123
	p := &n
	vp := reflect.ValueOf(p)

	fmt.Println(vp.CanSet(), vp.CanAddr())

	vn := vp.Elem() // get the value referenced by vp
	fmt.Println(vn)
	fmt.Println(vn.CanSet(), vn.CanAddr())
	vn.Set(reflect.ValueOf(456))
	fmt.Println(n)
}

func reflect_nonexport() {

	var s struct {
		X interface{} // an exported field
		y interface{} // a non-exported field
	}
	vp := reflect.ValueOf(&s)
	fmt.Println(vp.Type())
	fmt.Println(vp.Elem())
	vs := reflect.Indirect(vp)
	// vx and vy both represent interface values.
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vx.CanAddr()) // true true
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true

	vb := reflect.ValueOf(123)

	vx.Set(vb)

	fmt.Println(s)

}

func invertSlice(args []reflect.Value) []reflect.Value {
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)

	for i := n -1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}
}

func bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	invert := reflect.ValueOf(p).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

func usebind() {
	var invertInts func([]int) []int
	bind(&invertInts, invertSlice)
	fmt.Println(invertInts([]int{2,3,4}))
}

type V struct{
	A, b int
}

func (v V) AddSubThenScale(n int) (int, int) {
	return n * (v.A + v.b), n * (v.A - v.b)
}
func (v *V) Add2(n int) (int, int) {
	return n * (v.A + v.b), n * (v.A - v.b)
}


func callFunction() {
	v := V{5,2}
	vt := reflect.ValueOf(v)

	fmt.Println(vt)
	vm := vt.MethodByName("AddSubThenScale")

	fmt.Println(vm)

	res := vm.Call([]reflect.Value{reflect.ValueOf(3)})
	fmt.Println(res[0].Int())
	fmt.Println(res[1].Int())

	fmt.Println(int8(1) << 6)
}

func channel() {
	c := make(chan string, 2)
	vc := reflect.ValueOf(c)

	fmt.Println(vc.Len(), vc.Cap())
	// this will block
	vc.Send(reflect.ValueOf("C"))
	// will not bock, and return if send success
	successed := vc.TrySend(reflect.ValueOf("Hello"))
	fmt.Println(successed)
	successed = vc.TrySend(reflect.ValueOf("world"))
	fmt.Println(successed)
	fmt.Println(vc.Len(), vc.Cap())

	vs, succeeded := vc.TryRecv()
	fmt.Println(vs.String(), succeeded)

	vs, sendBeforeClosed := vc.Recv()
	fmt.Println(vs.String(), sendBeforeClosed)

	close(c)

	vs, sendBeforeClosed = vc.Recv()
	fmt.Println(vs.String(), sendBeforeClosed)
}

func selectctrl() {
	c := make(chan int, 1)
	vc := reflect.ValueOf(c)
	succeed := vc.TrySend(reflect.ValueOf(123))
	fmt.Println(succeed, vc.Len(), vc.Cap())

	vSend, vZero := reflect.ValueOf(789), reflect.Value{}

	branch := []reflect.SelectCase{
		{
			Dir: reflect.SelectDefault,
			Chan: vZero,
			Send: vZero,
		},
		{
			Dir: reflect.SelectRecv,
			Chan: vc,
		},
		{
			Dir: reflect.SelectSend,
			Chan: vc,
			Send: vSend,
		},
	}

	selIndex, vRec, sendBeforeClosed := reflect.Select(branch)

	fmt.Println(selIndex)
	fmt.Println(sendBeforeClosed) // true
	fmt.Println(vRec.Int())      // 123

	vc.Close()

	selIndex, _ , sendBeforeClosed = reflect.Select(branch[:2])
	fmt.Println(selIndex, sendBeforeClosed)
}