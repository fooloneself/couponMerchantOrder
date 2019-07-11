package main

import (
	"os"
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"
)
type Test struct {
	Name string
	Age int
}
func main(){
	var a Test

}

func (t Test)  {

}

func New(a reflect.Type) interface{} {
	a.Align()
	return new()
}