package gobasic

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

type GoBasic struct{}

func New() *GoBasic {
	return &GoBasic{}
}

func (gb *GoBasic) While(condition func() bool, action func()) {
	for condition() {
		action()
	}
}

func (gb *GoBasic) SetTimeout(duration time.Duration, action func()) {
	time.AfterFunc(duration, action)
}

func (gb *GoBasic) SetInterval(duration time.Duration, action func()) {
	ticker := time.NewTicker(duration)
	go func() {
		for range ticker.C {
			action()
		}
	}()
}

func (gb *GoBasic) Random(tipo string, length int) string {
	var charset string
	if tipo == "int" {
		charset = "0123456789"
	} else if tipo == "float" {
		charset = "0123456789."
	} else {
		charset = "abcdefghijklmnopqrstuvwxyz" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (gb *GoBasic) Handler(s string, replacements map[string]string) string {
	replacer := strings.NewReplacer()
	for k, v := range replacements {
		replacer = strings.NewReplacer(k, v)
		s = replacer.Replace(s)
	}
	return s
}

func (gb *GoBasic) SquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func (gb *GoBasic) Power(x, y float64) float64 {
	return math.Pow(x, y)
}

type Value struct {
	Initial, Final float64
}

func (gb *GoBasic) Define(initial, final float64) Value {
	return Value{Initial: initial, Final: final}
}

func (gb *GoBasic) RT(v1, v2 Value) float64 {
	return (v2.Final * v1.Initial) / v1.Final
}
