package config

import (
	"errors"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Telegram *Telegram `yaml:"telegram"`
	HTTP     *HTTP     `yaml:"http"`
	Servo    *Servo    `yaml:"servo"`
}

type Telegram struct {
	Token string `yaml:"token"`
	Chat  int64  `yaml:"chat"`
}

type HTTP struct {
	Bind      string `yaml:"bind"`
	SecretURI string `yaml:"secret_uri"`
}

type Servo struct {
	Pin           int           `yaml:"pin"`
	Pushes        int           `yaml:"pushes"`
	PushedAngle   int           `yaml:"pushed_angle"`
	PushedWait    time.Duration `yaml:"pushed_wait"`
	ReleasedAngle int           `yaml:"released_angle"`
	ReleasedWait  time.Duration `yaml:"released_wait"`
	AngleInactive int           `yaml:"angle_inactive"`
}

func NewConfig(path string) (*Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	if yaml.Unmarshal(content, &c) != nil {
		return nil, err
	}
	if validateConfig(&c) != nil {
		return nil, err
	}
	return &c, nil
}

func validateConfig(c *Config) error {
	if c.Telegram.Token == "" {
		return errors.New("missing Telegram token field")
	}
	if c.Telegram.Chat == 0 {
		return errors.New("missing Telegram chat ID field")
	}

	if c.HTTP.Bind == "" {
		return errors.New("missing http bind field")
	}
	if c.HTTP.SecretURI == "" {
		return errors.New("missing http secret URI field")
	}

	if c.Servo.AngleInactive == c.Servo.PushedAngle {
		return errors.New("servo inactive angle cannot be equal to servo pushed angle")
	}
	if c.Servo.PushedAngle == c.Servo.ReleasedAngle {
		return errors.New("servo pushed angle cannot be equal to servo released angle")
	}
	if c.Servo.PushedWait == 0 {
		return errors.New("servo pushed wait duration cannot be equal to zero")
	}
	if c.Servo.ReleasedWait == 0 {
		return errors.New("servo released wait duration cannot be equal to zero")
	}

	return nil
}
