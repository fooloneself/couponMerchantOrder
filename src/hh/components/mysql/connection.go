package mysql

type Connection struct {
	Dsn			string `yaml:"dsn"`
	Username 	string `yaml:"username"`
	Password	string `yaml:"password"`
	Db 			string `yaml:"db"`
	Charset		string `yaml:"charset"`
}