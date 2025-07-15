package internal

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type ConfigHTTPServer struct {
	Host string `env:"HOST" env-default:"0.0.0.0"`
	Port string `env:"PORT" env-default:"8080"`
}

type ConfigArchive struct {
	ObjectCount  int `env:"OBJECTCOUNT" env-default:"3"`
	ArchiveCount int `env:"ARCHIVECOUNT" env-default:"3"`
	Extensions   []string
}

type Config struct {
	Http *ConfigHTTPServer
	Arch *ConfigArchive
}

func MustLoad() *Config {
	godotenv.Load("./configs/.env")

	var cfghttp ConfigHTTPServer
	var cfgarch ConfigArchive

	if err := cleanenv.ReadEnv(&cfghttp); err != nil {
		log.Fatalln("Ошибка парсинга настроек сервера")
	}
	if err := cleanenv.ReadEnv(&cfgarch); err != nil {
		log.Fatalln("Ошибка парсинга настроек архива")
	}

	ext := os.Getenv("EXTENSIONS")
	if ext == "" {
		ext = "pdf,jpeg,jpg"
	}

	re := regexp.MustCompile(`^\.?[a-zA-Z0-9]+(?:,\.?[a-zA-Z0-9]+)*$`)

	if !re.MatchString(ext) {
		log.Fatalln("Ошибка парсинга допустимых расширений")
	}

	arrext := strings.Split(ext, ",")

	cfgarch.Extensions = arrext

	return &Config{
		Http: &cfghttp,
		Arch: &cfgarch,
	}
}
