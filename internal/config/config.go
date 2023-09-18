package config

type ConfigDB struct {
	DbHost     string `envconfig:"db_host" default:"localhost"`
	DbPort     string `envconfig:"db_port" default:"5432"`
	DbUser     string `envconfig:"db_user" default:"postgres"`
	DbPassword string `envconfig:"db_password" default:"postgres"`
	DbName     string `envconfig:"db_name" default:"postgres"`
}

type ConfigNats struct {
	ClusterName string `envconfig:"cluster_name" default:"test-cluster"`
	ClientName  string `envconfig:"client_name" default:"client-123"`
	ServerNats  string `envconfig:"server_nats" default:"localhost"`
}
