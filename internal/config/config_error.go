package config

const (
	ErrNotReadFile  = ConfigErr("can not read config file")
	ErrNotUnmarshal = ConfigErr("can not unmarshal config data")
)

type ConfigErr string

func (c ConfigErr) Error() string {
	return string(c)
}
