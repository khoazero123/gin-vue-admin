package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                 // Environmental value
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                              // Port value
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`                      // Database type:mysql(default)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`                   // OSS type
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"` // Multiple login interception
}
