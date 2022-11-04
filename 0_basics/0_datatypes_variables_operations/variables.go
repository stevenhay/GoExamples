package main

var MyString string
var InfererdString = "Hello!"
var FullString string = "I'm type declared!"

// Is this an int8? uint8? int32? int?
var WhatIsThis = 5
var WhatAboutThis = 4.20

var (
	I         int
	Dont      string
	Duplicate float32
	Var       bool
)

var HexValuesAreOK = 0xDEADC0DE
var AsAreOctals = 0o644
var AndBinary = 0b110100100

var ComparisonsJustWork = (AsAreOctals == AndBinary) // true

var FloatsCan = 9.987
var BeDeclared = .987
var ManyWays = 1.98e2
var EvenInHex = 0x2.p10 // 0xCpN = C * 2^N

var ImaginaryValues = 5 + 5i

var ThisIsARune = 'g'
var AsIsThis = '漢'
var OctalsAsRunes = '\141'
var HexAsRunes = '\x61'

var StringBlocksAreUseful = `I can
Write
Wherever
I want to
`

const DefineConstants = 420
const (
	Grouping  = 5
	AlsoWorks = 6
	ForConsts = 7
)
const A, B, C = 1, 2, 3

const (
	Iota  = iota // 0
	Is           // 1
	Petty        // 2
	Handy        // 3
)

var TypeConversionIsSimple = uint8(A)
var AndBack = int(TypeConversionIsSimple)
