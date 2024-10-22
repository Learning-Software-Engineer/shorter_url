package setting

type Configuration struct {
	Cassandra CassandraSettings `mapstructure:"cassandra"`
	Server    ServerSetting     `mapstructure:"server"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type CassandraSettings struct {
	Url      string `mapstructure:"url"`
	Keyspace string `mapstructure:"keyspace"`
	Port     int    `mapstructure:"port"`
}
