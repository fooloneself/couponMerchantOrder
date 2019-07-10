package main

import (
	"os"
	"flag"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	)
func main(){
	var path=flag.String("conf","./config.yml","配置文件路径")
	flag.Parse()
	var file,err=os.Open(*path)
	if err!=nil {
		fmt.Print(err)
		return
	}
	content,err:=ioutil.ReadAll(file)
	if err!=nil{
		fmt.Print(err)
		return
	}
	var a Config
	err=yaml.Unmarshal(content,&a)
	if err!=nil{
		fmt.Print(err)
		return
	}
	fmt.Print(a.Components.Mysql.Charset)
}