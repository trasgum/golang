package app

import (
	"flag"
	"time"

	//"github.com/mesos/mesos-go/api/v1/lib/encoding"
	//"github.com/ondrej-smola/mesos-go-http/lib/codec"
)

type Config struct {
	endpoints           []string
	url                 string
	//codec               codec
	timeout             time.Duration
	verbose             bool
	metrics             metrics
	resourceTypeMetrics bool
	summaryMetrics      bool
	compression         bool
	credentials         credentials
	authMode            string
}

func (cfg *Config) AddFlags(fs *flag.FlagSet) {
	//fs.Var(&cfg.codec, "codec", "Codec to encode/decode scheduler API communications [protobuf, json]")
	fs.StringVar(&cfg.url, "url", cfg.url, "Mesos scheduler API URL")
	fs.DurationVar(&cfg.timeout, "timeout", cfg.timeout, "Mesos scheduler API connection timeout")
	fs.BoolVar(&cfg.verbose, "verbose", cfg.verbose, "Verbose logging")
	fs.IntVar(&cfg.metrics.port, "metrics.port", cfg.metrics.port, "Port of metrics server (listens on server.address)")
	fs.StringVar(&cfg.metrics.path, "metrics.path", cfg.metrics.path, "URI path to metrics endpoint")
	fs.BoolVar(&cfg.resourceTypeMetrics, "resourceTypeMetrics", cfg.resourceTypeMetrics, "Collect scalar resource metrics per-type")
	fs.BoolVar(&cfg.summaryMetrics, "summaryMetrics", cfg.summaryMetrics, "Collect summary metrics for tasks launched per-offer-cycle, offer processing time, etc.")
	fs.BoolVar(&cfg.compression, "compression", cfg.compression, "When true attempt to use compression for HTTP streams.")
	fs.StringVar(&cfg.credentials.username, "credentials.username", cfg.credentials.username, "Username for Mesos authentication")
	fs.StringVar(&cfg.credentials.password, "credentials.passwordFile", cfg.credentials.password, "Path to file that contains the password for Mesos authentication")
	fs.StringVar(&cfg.authMode, "authmode", cfg.authMode, "Method to use for Mesos authentication; specify '"+AuthModeBasic+"' for simple HTTP authentication")
}

const AuthModeBasic = "Basic"

func NewConfig() Config {
	return Config{
		url:              env("MESOS_MASTER_HTTP", "http://127.0.0.1:5050/api/v1"),
		timeout:          envDuration("MESOS_CONNECT_TIMEOUT", "20s"),
		//codec:            codec{Codec: encoding.MediaTypeProtobuf.Codec()},
		metrics: metrics{
			port: envInt("PORT0", "64009"),
			path: env("METRICS_API_PATH", "/metrics"),
		},
		credentials: credentials{
			username: env("AUTH_USER", ""),
			password: env("AUTH_PASSWORD_FILE", ""),
		},
		authMode: env("AUTH_MODE", ""),
	}
}

type server struct {
	address string
	port    int
}

type metrics struct {
	port int
	path string
}

type credentials struct {
	username string
	password string
}
