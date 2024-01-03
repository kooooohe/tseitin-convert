package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 入力文字列
	//DNF
	input := `( x1 ∧  x2 ∧  ¬x3 ∧  x4 ∧  ¬x5 ) ∨
( ¬x1 ∧  x2 ∧  ¬x3 ) ∨
( x1 ∧  x2 ∧  ¬x3 ∧  x4 ∧  x5 )`
	// 改行で分割
	rows := strings.Split(input, "∨")

	// 2次元配列を作成
	var dnfs [][]string
	for _, row := range rows {
		elements := strings.Split(row, " ")
		t := []string{}
		for _,v := range elements {
			v := strings.TrimSpace(v)
			if len(v) > 0 && v != "\n" && v != " " && v != "(" && v!=")" && v != "∧"{
				t = append(t, v)

			}
		}
		dnfs = append(dnfs, t)
	}


	// 結果の出力
	ys := "( "
	cnfs := ""
	for i, v := range dnfs {
		y := "y" + strconv.Itoa(i)
		for j, vv := range v {
			cnfs += "( " + "¬" + y + " ∨  " + vv + " ) "
			if j != len(v)-1 {
				cnfs += " ∧  "
			}
		}

			ys += " " + y 
		if i != len(dnfs)-1 {
			ys += " ∨  "
			cnfs += " ∧  "
		}
	}
	ys += " ) ∧ "
	fmt.Println(ys)
	fmt.Println(cnfs)
}
