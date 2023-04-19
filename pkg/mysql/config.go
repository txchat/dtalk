package mysql

import (
	"crypto/tls"
	"time"

	"github.com/go-sql-driver/mysql"
)

// Config is a configuration parsed from a DSN string.
// If a new Config is created instead of being parsed from a DSN string,
// the NewConfig function should be used, which sets default values.
type Config struct {
	User             string            `json:","`              // Username
	Passwd           string            `json:",optional"`      // Password (requires User)
	Net              string            `json:",default=tcp"`   // Network type
	Addr             string            `json:","`              // Network address (requires Net)
	DBName           string            `json:","`              // Database name
	Params           map[string]string `json:",optional"`      // Connection parameters
	Collation        string            `json:",optional"`      // Connection collation
	Loc              string            `json:",default=local"` // Location for time.Time values
	MaxAllowedPacket int               `json:",optional"`      // Max packet size allowed
	ServerPubKey     string            `json:",optional"`      // Server public key name
	TLSConfig        string            `json:",optional"`      // TLS configuration name
	TLS              *tls.Config       `json:",optional"`      // TLS configuration, its priority is higher than TLSConfig
	Timeout          time.Duration     `json:",optional"`      // Dial timeout
	ReadTimeout      time.Duration     `json:",optional"`      // I/O read timeout
	WriteTimeout     time.Duration     `json:",optional"`      // I/O write timeout

	AllowAllFiles            bool `json:",optional"` // Allow all files to be used with LOAD DATA LOCAL INFILE
	AllowCleartextPasswords  bool `json:",optional"` // Allows the cleartext client side plugin
	AllowFallbackToPlaintext bool `json:",optional"` // Allows fallback to unencrypted connection if server does not support TLS
	AllowNativePasswords     bool `json:",optional"` // Allows the native password authentication method
	AllowOldPasswords        bool `json:",optional"` // Allows the old insecure password method
	CheckConnLiveness        bool `json:",optional"` // Check connections for liveness before using them
	ClientFoundRows          bool `json:",optional"` // Return number of matching rows instead of rows changed
	ColumnsWithAlias         bool `json:",optional"` // Prepend table alias to column names
	InterpolateParams        bool `json:",optional"` // Interpolate placeholders into query string
	MultiStatements          bool `json:",optional"` // Allow multiple statements in one query
	ParseTime                bool `json:",optional"` // Parse time values to time.Time
	RejectReadOnly           bool `json:",optional"` // Reject read-only connections
}

func (cfg *Config) SetParam(key, value string) {
	if cfg.Params == nil {
		cfg.Params = make(map[string]string)
	}
	cfg.Params[key] = value
}

func (cfg *Config) GetSQLDriverConfig() *mysql.Config {
	cp := &mysql.Config{
		User:                     cfg.User,
		Passwd:                   cfg.Passwd,
		Net:                      cfg.Net,
		Addr:                     cfg.Addr,
		DBName:                   cfg.DBName,
		Collation:                cfg.Collation,
		MaxAllowedPacket:         cfg.MaxAllowedPacket,
		ServerPubKey:             cfg.ServerPubKey,
		TLSConfig:                cfg.TLSConfig,
		Timeout:                  cfg.Timeout,
		ReadTimeout:              cfg.ReadTimeout,
		WriteTimeout:             cfg.WriteTimeout,
		AllowAllFiles:            cfg.AllowAllFiles,
		AllowCleartextPasswords:  cfg.AllowCleartextPasswords,
		AllowFallbackToPlaintext: cfg.AllowFallbackToPlaintext,
		AllowNativePasswords:     cfg.AllowNativePasswords,
		AllowOldPasswords:        cfg.AllowOldPasswords,
		CheckConnLiveness:        cfg.CheckConnLiveness,
		ClientFoundRows:          cfg.ClientFoundRows,
		ColumnsWithAlias:         cfg.ColumnsWithAlias,
		InterpolateParams:        cfg.InterpolateParams,
		MultiStatements:          cfg.MultiStatements,
		ParseTime:                cfg.ParseTime,
		RejectReadOnly:           cfg.RejectReadOnly,
	}
	switch cfg.Loc {
	case "local":
		cp.Loc = time.Local
	}
	if cp.TLS != nil {
		cp.TLS = cfg.TLS.Clone()
	}
	if len(cp.Params) > 0 {
		cp.Params = make(map[string]string, len(cfg.Params))
		for k, v := range cfg.Params {
			cp.Params[k] = v
		}
	}
	return cp
}
