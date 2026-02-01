package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Conf Config

// Config 总配置结构体
type Config struct {
	MySQL         MySQLConfig    `mapstructure:"mysql"`
	Redis         RedisConfig    `mapstructure:"redis"`
	Cos           CosConfig      `mapstructure:"cos"`
	Etcd          EtcdConfig     `mapstructure:"etcd"`
	Kafka         KafkaConfig    `mapstructure:"kafka"`
	Elasticsearch elasticsearch  `mapstructure:"elasticsearch"`
	Server        ServerConfig   `mapstructure:"server"`
	Otel          OtelConfig     `mapstructure:"otel"`
	RabbitMQ      RabbitMQConfig `mapstructure:"rabbitmq"`
	Secret        SecretConfig   `mapstructure:"secret"`
	JWT           JWTConfig      `mapstructure:"jwt"`
}

// MySQLConfig MySQL配置结构体
type MySQLConfig struct {
	Addr      string `mapstructure:"addr"`
	User      string `mapstructure:"user"`
	UserDB    string `mapstructure:"userdb"`
	ProductDB string `mapstructure:"productdb"`
	CouponDB  string `mapstructure:"coupondb"`
	CartDB    string `mapstructure:"cartdb"`
	StockDB   string `mapstructure:"stockdb"`
	OrderDB   string `mapstructure:"orderdb"`
	PaymentDB string `mapstructure:"paymentdb"`
	Password  string `mapstructure:"password"`
	Charset   string `mapstructure:"charset"`
}

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
type CosConfig struct {
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	SecretId  string `mapstructure:"secretId"`
	SecretKey string `mapstructure:"secretKey"`
}

type EtcdConfig struct {
	Endpoints []string `mapstructure:"endpoints"`
}

type KafkaConfig struct {
	Addr     string `mapstructure:"addr"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type RabbitMQConfig struct {
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Address      string `mapstructure:"address"`
	Exchange     string `mapstructure:"exchange"`
	DelayQueue   string `mapstructure:"delayqueue"`
	ProcessQueue string `mapstructure:"processqueue"`
	RoutingKey   string `mapstructure:"routingkey"`
}

type elasticsearch struct {
	Address      string `mapstructure:"address"`
	ProductIndex string `mapstructure:"productIndex"`
}

type ServerConfig struct {
	Gateway string `mapstructure:"gateway"`
	User    string `mapstructure:"user"`
	Product string `mapstructure:"product"`
	Stock   string `mapstructure:"stock"`
	Cart    string `mapstructure:"cart"`
	Coupon  string `mapstructure:"coupon"`
	Order   string `mapstructure:"order"`
	Payment string `mapstructure:"payment"`
}

type OtelConfig struct {
	Address string `mapstructure:"address"`
}

type JWTConfig struct {
	AccessExpire  string `mapstructure:"accessExpire"`
	RefreshExpire string `mapstructure:"refreshExpire"`
	PrivateKey    string `mapstructure:"privateKey"`
	PublicKey     string `mapstructure:"publicKey"`
	Issuer        string `mapstructure:"issuer"`
}

type SecretConfig struct {
	TopSecret     string `mapstructure:"topSecret"`
	PaymentSecret string `mapstructure:"paymentSecret"`
}

// Init 初始化配置
func Init() {
	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %s", err)
	}

	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置查找路径
	viper.AddConfigPath(filepath.Join(workDir, "config"))
	viper.AddConfigPath(workDir)
	viper.AddConfigPath(filepath.Join(workDir, "../../config"))

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 映射到结构体
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Unable to decode into struct: %s", err)
	}

	log.Println("Config loaded successfully!")
}
