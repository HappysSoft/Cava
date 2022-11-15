package model

type CavaToken int

const (
	tInt CavaToken = iota
	tInt8
	tInt2
	tInt16
	tFloat
	tString
	tId
	tType

	tCommentOneLine

	tUnKnown
)

var cavaToken = []string{
	tInt:    "0|[1-9][\\d]*",
	tInt8:   "0(0|[1-7][0-7]*)",
	tInt2:   "0b(0|)",
	tInt16:  "",
	tFloat:  "[\\d]+[.][\\d]+",
	tString: "\"[^\"]*\"|'[^']*'",
	tId:     "",
	tType:   "int",

	tCommentOneLine: "\\\\.*",
	tUnKnown:        ".",
}

var cavaTokenFunc = []func(text string) error{
	tInt: func(text string) error {

	},
	tInt8: func(text string) error {

	},
}

func (t CavaToken) String() string {
	return cavaToken[t]
}

func (t CavaToken) Run(text string) error {
	return cavaTokenFunc[t](text)
}
