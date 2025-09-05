package evaluator

import (
	"compiler/object"
)

var builtins = map[string]*object.Builtin {
	"len": object.GetBuiltInByName("len"),
	"puts": object.GetBuiltInByName("puts"),
	"first": object.GetBuiltInByName("first"),
	"last": object.GetBuiltInByName("last"),
	"rest": object.GetBuiltInByName("rest"),
	"push": object.GetBuiltInByName("push"),
}