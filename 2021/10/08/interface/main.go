package main

import (
	"fmt"
)

type SourceCollector interface {
	collectSource()
}

type TargetCompiler interface {
	compileToTarget()
}

type CrossCompiler struct {
	collector SourceCollector
	compiler  TargetCompiler
}

func (v CrossCompiler) crossCompile() {
	v.collector.collectSource()
	v.compiler.compileToTarget()
}

type IPhoneCompiler struct {
}

func (v IPhoneCompiler) collectSource() {
	fmt.Println("IPhoneCompiler.collectSource")
}

func (v IPhoneCompiler) compileToTarget() {
	fmt.Println("IPhoneCompiler.compileToTarget")
}

type AndroidCompiler struct {
}

func (v AndroidCompiler) collectSource() {
	fmt.Println("AndroidCompiler.collectSource")
}

func (v AndroidCompiler) compileToTarget() {
	fmt.Println("AndroidCompiler.compileToTarget")
}

func main() {
	iPhone := IPhoneCompiler{}
	androd := AndroidCompiler{}
	compiler := CrossCompiler{iPhone, androd}
	compiler.crossCompile()
}
