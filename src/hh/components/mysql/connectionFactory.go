package mysql


type ConnectionFactory struct {
	Dsn			string `yaml:"dsn"`
	Username 	string `yaml:"username"`
	Password	string `yaml:"password"`
	Db 			string `yaml:"db"`
	Charset		string `yaml:"charset"`
}

func (factory ConnectionFactory) New()(interface{},error) {
	return nil,nil
}