package main

import (
	"net/http"
	"strings"

	"github.com/micro/cli"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
	"github.com/rs/cors"
)

type allowedCors struct {
	allowedHeaders []string
	allowedOrigins []string
	allowedMethods []string
}

func (ac *allowedCors) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "cors-allowed-headers",
			Usage:  "Comma-seperated list of allowed headers",
			EnvVar: "CORS_ALLOWED_HEADERS",
		},
		cli.StringFlag{
			Name:   "cors-allowed-origins",
			Usage:  "Comma-seperated list of allowed origins",
			EnvVar: "CORS_ALLOWED_ORIGINS",
		},
		cli.StringFlag{
			Name:   "cors-allowed-methods",
			Usage:  "Comma-seperated list of allowed methods",
			EnvVar: "CORS_ALLOWED_METHODS",
		},
	}
}

func (ac *allowedCors) Commands() []cli.Command {
	return nil
}

func (ac *allowedCors) Handler() plugin.Handler {
	return func(ha http.Handler) http.Handler {
		hf := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ha.ServeHTTP(w, r)
		})

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			corsOptions := cors.Options{
				Debug: true,
			}

			if len(ac.allowedHeaders) != 0 {
				corsOptions.AllowedHeaders = ac.allowedHeaders
			}

			if len(ac.allowedMethods) != 0 {
				corsOptions.AllowedMethods = ac.allowedMethods
			}

			if len(ac.allowedOrigins) != 0 {
				corsOptions.AllowedOrigins = ac.allowedOrigins
			}

			cors.New(corsOptions).ServeHTTP(w, r, hf)
		})
	}
}

func (ac *allowedCors) Init(ctx *cli.Context) error {
	ac.allowedHeaders = ac.parseAllowed(ctx, "cors-allowed-headers", true)
	ac.allowedMethods = ac.parseAllowed(ctx, "cors-allowed-methods", false)
	ac.allowedOrigins = ac.parseAllowed(ctx, "cors-allowed-origins", false)

	return nil
}

func (ac *allowedCors) parseAllowed(ctx *cli.Context, flagName string, addLowerVariants bool) []string {
	fv := ctx.String(flagName)

	// no op
	if len(fv) == 0 {
		return nil
	}

	list := strings.Split(fv, ",")
	fixedList := []string{}

	for _, val := range list {
		trimmedVal := strings.TrimSpace(val)
		fixedList = append(fixedList, trimmedVal)

		if addLowerVariants {
			lval := strings.ToLower(trimmedVal)

			if trimmedVal != lval {
				fixedList = append(fixedList, lval)
			}
		}
	}

	return fixedList
}

func (ac *allowedCors) String() string {
	return "cors-allowed-(headers|origins|methods)"
}

// NewPlugin Creates the CORS Plugin
func NewPlugin() plugin.Plugin {
	return &allowedCors{}
}

func init() {
	api.Register(NewPlugin())
}
