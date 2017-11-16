package main

import (
	"fmt"
	"testing"
)

func TestCompile(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{
			in: "2 + 3",
			out: `[
    {
        "IsStatement": false,
        "IsExpression": true,
        "Abs": {
            "Name": "+",
            "Arguments": [
                [
                    2,
                    3
                ]
            ],
            "Val": 5
        }
    }
]`,
		},
		{
			in: "print \"My name is Jack White\"",
			out: `[
    {
        "IsStatement": true,
        "IsExpression": false,
        "Abs": {
            "Name": "print",
            "Arguments": [
                "\"My name is Jack White\""
            ]
        }
    }
]`,
		},
	}
	for _, c := range cases {
		r, err := Compile(c.in)
		if err != nil {
			fmt.Println(err)
		}
		if string(r) != c.out {
			fmt.Printf("got %s\nvs  %s\n", string(r), c.out)
			t.Fail()
		}
	}
}
