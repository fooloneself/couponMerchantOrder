package main

import "hh/components/mysql"

type Config struct {
	Components struct{
		Mysql mysql.Connection `yaml:"db"`
	} `yaml:"components"`
}