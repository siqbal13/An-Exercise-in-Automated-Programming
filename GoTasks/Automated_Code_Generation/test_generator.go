package main

import (
	"github.com/dave/jennifer/jen"
)

func main() {
	f := jen.NewFile("main")

	// Import statements
	f.ImportName("testing", "testing")

	// Generate test function for linreg
	f.Func().Id("TestLinreg").Params(jen.Id("t").Op("*").Qual("testing", "T")).Block(
		jen.Var().Id("x").Op("=").Index().Float64().Values(jen.Lit(1), jen.Lit(2), jen.Lit(3)),
		jen.Var().Id("y").Op("=").Index().Float64().Values(jen.Lit(1), jen.Lit(2), jen.Lit(3)),
		jen.Var().Id("expectedSlope").Op("=").Lit(1.0),
		jen.Var().Id("expectedIntercept").Op("=").Lit(0.0),
		jen.List(jen.Id("m"), jen.Id("c")).Op(":=").Id("linreg").Call(jen.Id("x"), jen.Id("y")), // Call linreg function from the same package
		jen.If(jen.Id("m").Op("!=").Id("expectedSlope").Op("||").Id("c").Op("!=").Id("expectedIntercept")).Block(
			jen.Id("t").Dot("Errorf").Call(
				jen.Lit("linreg(%v, %v) = (%f, %f), want (%f, %f)"),
				jen.Id("x"), jen.Id("y"), jen.Id("m"), jen.Id("c"), jen.Id("expectedSlope"), jen.Id("expectedIntercept"),
			),
		),
	)

	// Save generated code to a file
	err := f.Save("main_test.go")
	if err != nil {
		panic(err)
	}
}
