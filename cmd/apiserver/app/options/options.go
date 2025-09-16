package options

import (
	"errors"
	"fmt"
	"time"

	genericoptions "github.com/pachirode/pkg/options"
	stringsutil "github.com/pachirode/pkg/util/strings"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/pachirode/monitor/internal/apiserver"
)

var availableServerModes = sets.New(
	apiserver.GinServerMode,
	apiserver.GRPCServerMode,
	apiserver.GRPCGatewayServerMode,
)

type ServerOptions struct {
	// ServerMode 定义服务器模式：gRPC、Gin HTTP、HTTP Reverse Proxy.
	ServerMode string `json:"server-mode" mapstructure:"server-mode"`
	// JWTKey 定义 JWT 密钥.
	JWTKey string `json:"jwt-key" mapstructure:"jwt-key"`
	// Expiration 定义 JWT Token 的过期时间.
	Expiration time.Duration `json:"expiration" mapstructure:"expiration"`
	// TLSOptions 包含 TLS 配置选项.
	TLSOptions *genericoptions.TLSOptions `json:"tls" mapstructure:"tls"`
	// HTTPOptions 包含 HTTP 配置选项.
	HTTPOptions *genericoptions.HTTPOptions `json:"http" mapstructure:"http"`
	// GRPCOptions 包含 gRPC 配置选项.
	GRPCOptions *genericoptions.GRPCOptions `json:"grpc" mapstructure:"grpc"`
	// MySQLOptions 包含 MySQL 配置选项.
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	opts := &ServerOptions{
		ServerMode:   apiserver.GRPCGatewayServerMode,
		JWTKey:       "Vnk8GRhFEz7uD3mltQOBPZZLXiLJZNVjhT81sKpnzIuMf2",
		Expiration:   2 * time.Hour,
		TLSOptions:   genericoptions.NewTLSOptions(),
		HTTPOptions:  genericoptions.NewHTTPOptions(),
		GRPCOptions:  genericoptions.NewGRPCOptions(),
		MySQLOptions: genericoptions.NewMySQLOptions(),
	}
	opts.HTTPOptions.Addr = ":5555"
	opts.GRPCOptions.Addr = ":6666"
	return opts
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&o.ServerMode, "server-mode", o.ServerMode, fmt.Sprintf("Server mode, available options: %v", availableServerModes.UnsortedList()))
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "JWT signing key. Must be at least 6 characters long.")
	// 绑定 JWT Token 的过期时间选项到命令行标志。
	// 参数名称为 `--expiration`，默认值为 o.Expiration
	fs.DurationVar(&o.Expiration, "expiration", o.Expiration, "The expiration duration of JWT tokens.")
	o.TLSOptions.AddFlags(fs)
	o.HTTPOptions.AddFlags(fs)
	o.GRPCOptions.AddFlags(fs)
	o.MySQLOptions.AddFlags(fs)
}

func (o *ServerOptions) Validate() error {
	errs := []error{}

	// 校验 ServerMode 是否有效
	if !availableServerModes.Has(o.ServerMode) {
		errs = append(errs, fmt.Errorf("invalid server mode: must be one of %v", availableServerModes.UnsortedList()))
	}

	// 校验 JWTKey 长度
	if len(o.JWTKey) < 6 {
		errs = append(errs, errors.New("JWTKey must be at least 6 characters long"))
	}

	// 校验子选项
	errs = append(errs, o.TLSOptions.Validate()...)
	errs = append(errs, o.HTTPOptions.Validate()...)
	errs = append(errs, o.MySQLOptions.Validate()...)

	// 如果是 gRPC 或 gRPC-Gateway 模式，校验 gRPC 配置
	if stringsutil.StringIn(o.ServerMode, []string{apiserver.GRPCServerMode, apiserver.GRPCGatewayServerMode}) {
		errs = append(errs, o.GRPCOptions.Validate()...)
	}

	// 合并所有错误并返回
	return utilerrors.NewAggregate(errs)
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		ServerMode:   o.ServerMode,
		JWTKey:       o.JWTKey,
		Expiration:   o.Expiration,
		TLSOptions:   o.TLSOptions,
		HTTPOptions:  o.HTTPOptions,
		GRPCOptions:  o.GRPCOptions,
		MySQLOptions: o.MySQLOptions,
	}, nil
}
