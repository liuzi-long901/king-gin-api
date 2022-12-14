package config

type Configuration struct {
	App           App         `mapstructure:"app" json:"app" yaml:"app"`
	Log           Log         `mapstructure:"log" json:"log" yaml:"log"`
	Database      Database    `mapstructure:"database" json:"database" yaml:"database"`
	Swagger       Swagger     `mapstructure:"swagger" json:"swagger" yaml:"swagger"`
	Jwt           Jwt         `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis         Redis       `mapstructure:"redis" json:"redis" yaml:"redis"`
	Storage       Storage     `mapstructure:"storage" json:"storage" yaml:"storage"`
	Rabbitmq      Rabbitmq    `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	RabbitmqQueue RabbitQueue `mapstructure:"rabbitQueue" json:"rabbitQueue" yaml:"rabbitQueue"`
}
