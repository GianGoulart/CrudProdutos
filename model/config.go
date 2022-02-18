package model

import (
	"bytes"
	"encoding/base32"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xtgo/uuid"
)

// Config interface de configuração
type Config interface {
	GetInt(key string) int
	GetBool(key string) bool
	GetString(key string) string
	GetFloat64(key string) float64
	GetDuration(key string) time.Duration
	GetStringSlice(key string) []string
	Close()
}

type configImpl struct {
	vmain     *viper.Viper
	vreplacer *viper.Viper

	kill chan bool
}

func (c *configImpl) GetBool(key string) bool {
	return c.vmain.GetBool(key)
}

// GetString realiza a troca das variaveis em tempo de execução e retorna uma string
func (c *configImpl) GetString(key string) string {
	value := c.vmain.GetString(key)

	for k, v := range c.vreplacer.AllSettings() {
		old := "$" + k + "$"
		new, ok := v.(string)
		if !ok {
			continue
		}

		value = strings.ReplaceAll(value, old, new)
	}

	return value
}

func (c *configImpl) GetDuration(key string) time.Duration {
	return c.vmain.GetDuration(key)
}

func (c *configImpl) GetInt(key string) int {
	return c.vmain.GetInt(key)
}

func (c *configImpl) GetFloat64(key string) float64 {
	return c.vmain.GetFloat64(key)
}

func (c *configImpl) GetStringSlice(key string) []string {
	return c.vmain.GetStringSlice(key)
}

// Close encerra a função de watch
func (c *configImpl) Close() {
	c.kill <- true
}

// readRemoteConfig adiciona o provider e le a configuração remota
func readRemoteConfig(v *viper.Viper, provider, endpoint, path, token string) error {
	v.SetConfigType("json")
	v.AddRemoteProvider(provider, endpoint, path)
	if err := v.ReadRemoteConfig(); err != nil {
		return err
	}
	return nil
}

// Watch fica escutando modificações no config remoto
func Watch(fn func(c Config, quit chan bool)) {
	quit, kill := make(chan bool), make(chan bool)
	vmain, vreplacer := viper.New(), viper.New()

	c := &configImpl{
		vmain:     vmain,
		vreplacer: vreplacer,
		kill:      kill,
	}

	//Substitui o _ por .
	vmain.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// realiza o bind das variaveis de ambiente
	vmain.AutomaticEnv()

	// setando por default os ambientes com config remoto
	tcenv := vmain.GetString("tc")
	if tcenv == "prod" || tcenv == "preprod" || tcenv == "hml" {
		provider := vmain.GetString("remote.provider")
		endpoint := vmain.GetString("remote.endpoint")
		token := vmain.GetString("remote.token")
		path := vmain.GetString("remote.path")

		if err := readRemoteConfig(vmain, provider, endpoint, path, token); err != nil {
			logrus.Fatal("não consegui ler o arquivo de configuração remoto. com o path", path, err.Error())
		}

		replaces := vmain.GetStringSlice("remote.replace")
		for _, rpath := range replaces {
			vpath := viper.New()

			if err := readRemoteConfig(vpath, provider, endpoint, rpath, token); err != nil {
				logrus.Fatal("não consegui ler o arquivo de configuração remoto. com o path", rpath, err.Error())
			}

			if err := vreplacer.MergeConfigMap(vpath.AllSettings()); err != nil {
				logrus.Fatal("não consegui realizar o merge do arquivo de configuração remoto. com o path", rpath, err.Error())
			}
		}
	} else {
		// seta o arquivo de configuração local
		vmain.SetConfigFile("./config.json")

		// realiza a leitura das configurações locais
		if err := vmain.ReadInConfig(); err != nil {
			logrus.Fatal("não consegui ler o arquivo de configuração local")
		}
	}

	// inicia o server
	go fn(c, quit)

	<-kill
}

// validatorImpl modelo para a validação do bind dos requests
type validatorImpl struct {
	v *validator.Validate
}

// Validate metodo que implementa a interface do validator para execução da validação
func (cv *validatorImpl) Validate(i interface{}) error {
	return cv.v.Struct(i)
}

// Validator interface do componente de validação
type Validator interface {
	Validate(i interface{}) error
}

// New cria uma nova implementação da interface Validator
func New() Validator {
	return &validatorImpl{
		v: validator.New(),
	}
}

var encoding = base32.NewEncoding("ybndrfg8ejkmcxotp11uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	w := uuid.NewRandom()
	encoder.Write(w.Bytes())
	encoder.Close()
	b.Truncate(26)
	return b.String()
}
