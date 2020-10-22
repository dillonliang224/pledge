package debug

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"git.dillonliang.cn/micro-svc/pledge/library/config"
)

type Service interface {
	Ping(ctx context.Context) error
}

func Start(c config.Common, svc Service) {
	debugMux := http.NewServeMux()

	debugMux.Handle("/metrics", promhttp.Handler())
	debugMux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := map[string]string{
			"status": "UP",
		}

		statusCode := http.StatusOK
		if err := svc.Ping(r.Context()); err != nil {
			statusCode = http.StatusInternalServerError
			res["status"] = "DOWN"
			res["err"] = err.Error()
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		_ = json.NewEncoder(w).Encode(res)
	}))

	debugMux.HandleFunc("/debug/pprof/", pprof.Index)
	debugMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	debugMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	debugMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	debugMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	go func() {
		if err := http.ListenAndServe(c.Port.DEBUG, debugMux); err != nil {
			panic(err)
		}
	}()
}
