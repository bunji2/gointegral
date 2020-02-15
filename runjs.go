package main

import (
	"fmt"
	"math"

	"github.com/robertkrimen/otto"
)

const (
	varInterval  = "interval"
	varNumSample = "n"
	varFunction  = "f"
	varAnswer    = "answer"
)

func runJS(jsFilePath string) (err error) {

	vm, jsF, interval, n, err := initialize(jsFilePath)

	if err != nil {
		return
	}

	f := func(x float64) float64 {

		value, err := jsF.Call(jsF, x)
		if err != nil {
			panic(err)
		}

		v, err := value.ToFloat()
		if err != nil {
			panic(err)
		}
		return v
	}

	fmt.Println(varInterval+" =", interval)
	fmt.Println("n =", n)

	result := simpson(interval[0], interval[1], n, f)
	fmt.Println("result =", result)

	// answer 変数が定義されている場合には
	// 計算値と正解との比較を行う。
	hasAnswer := func() (answer float64, ok bool) {
		// 変数 answer の取得
		v, err := vm.Get(varAnswer)
		if err != nil {
			return
		}
		// 未定義な変数の場合。ちなみに v は NaN になる。
		if !v.IsDefined() {
			return
		}
		// float64 の値を取り出す
		answer, err = v.ToFloat()
		if err != nil {
			return
		}
		ok = true
		return
	}

	answer, ok := hasAnswer()
	if ok {
		// 正解の表示
		fmt.Println(varAnswer+" =", answer)
		// エラー率の表示 error = |answer - result|/answer * 100 [%]
		fmt.Println("error  =", math.Abs(answer-result)/answer*100, "[%]")
	}
	return
}

// initialize は指定された JS ファイルを読み込み初期化を行う関数。
// otto オブジェクトを生成し、関数 f、積分区間interval、標本数 n を返す。
func initialize(jsFilePath string) (vm *otto.Otto, jsF otto.Value, interval [2]float64, n int64, err error) {
	vm = otto.New()

	var script *otto.Script
	script, err = vm.Compile(jsFilePath, nil)
	if err != nil {
		return
	}

	_, err = vm.Run(script)
	if err != nil {
		return
	}

	jsF, err = vm.Get("f")
	if err != nil {
		return
	}
	if !jsF.IsDefined() {
		err = fmt.Errorf(`function "%s" is not found`, varFunction)
		return
	}

	var v otto.Value
	v, err = vm.Get(varInterval)
	if err != nil {
		return
	}
	if !v.IsDefined() { // 未定義の場合はエラー
		err = fmt.Errorf(`var "%s" is not found`, varInterval)
		return
	}

	interval, err = valueToArrayFloat64(v)
	if err != nil {
		return
	}
	//fmt.Println(varInterval+" =", interval)

	v, err = vm.Get(varNumSample)
	if err != nil {
		return
	}
	if !v.IsDefined() {
		err = fmt.Errorf(`var "%s" is not found`, varNumSample)
		return
	}
	n, err = v.ToInteger()

	return
}

// valueToArrayFloat64 は otto から取得した値が配列かどうかをチェックし、
// 数値の配列ならば [2]float64 に変換する関数。
// なぜか otto には配列を操作するメソッドが用意されていない。
func valueToArrayFloat64(v otto.Value) (a [2]float64, err error) {
	vv, _ := v.Export()

	// func (self Value) Export() (interface{}, error)
	// Export will attempt to convert the value to a Go representation and return it via an interface{} kind.
	// Export returns an error, but it will always be nil. It is present for backwards compatibility.
	// If a reasonable conversion is not possible, then the original value is returned.
	// undefined   -> nil (FIXME?: Should be Value{})
	// null        -> nil
	// boolean     -> bool
	// number      -> A number type (int, float32, uint64, ...)
	// string      -> string
	// Array       -> []interface{}
	// Object      -> map[string]interface{}

	// ドキュメントには Array -> []interface{} とあるが、
	// 今回のケースでは、[]float64, []int64, []interface{} の
	// いずれかになる。

	//fmt.Println("vv =", vv)

	switch vvv := vv.(type) {
	case []float64:
		a = [2]float64{vvv[0], vvv[1]}
	case []int64:
		a = [2]float64{float64(vvv[0]), float64(vvv[1])}
	case []interface{}:
		a = [2]float64{}
		for i, vvvv := range vvv {
			switch vvvv2 := vvvv.(type) {
			case int64:
				a[i] = float64(vvvv2)
			case float64:
				a[i] = vvvv2
			default:
				err = fmt.Errorf("value %v is not expected type", vvvv2)
				return
			}
		}
	default:
		err = fmt.Errorf("value %v is not expected type", vvv)
	}
	/*
		// []float64 かどうかをチェック

		vvv, ok := vv.([]float64)
		if ok {
			a = [2]float64{vvv[0], vvv[1]}
			return
		}

		// []int64 かどうかをチェック

		vvv2, ok := vv.([]int64)
		if ok {
			a = [2]float64{float64(vvv2[0]), float64(vvv2[1])}
			return
		}

		// []interface{} かどうかをチェック

		vvv3, ok := vv.([]interface{})
		if !ok {
			err = fmt.Errorf("v %v is not array", v)
			return
		}
		a = [2]float64{}
		//fmt.Println("vvv =", vvv)

		if i, ok := vvv3[0].(int64); ok {
			a[0] = float64(i)
		} else if f, ok := vvv3[0].(float64); ok {
			a[0] = f
		} else {
			err = fmt.Errorf("value %v is not expected type", vvv3[0])
			return
		}

		if i, ok := vvv3[1].(int64); ok {
			a[1] = float64(i)
		} else if f, ok := vvv3[1].(float64); ok {
			a[1] = f
		} else {
			err = fmt.Errorf("value %v is not expected type", vvv3[1])
		}
	*/

	return
}
