// restapi project main.go
package main

import (
	"context"
	"flag"

	"org/helm/tiller/restapi/gateway"

	"github.com/golang/glog"
)

var (
	httpAddr   = flag.String("listen", ":9090", "address:port to listen on")
	endpoint   = flag.String("endpoint", "localhost:44134", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
	swaggerDir = flag.String("swagger_dir", "proto/hapi", "path to the directory which contains swagger definitions")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	//	md := metadata.Pairs("x-helm-api-client", version.GetVersion())
	//	ctx := metadata.NewOutgoingContext(context.TODO(), md)
	opts := gateway.Options{
		Addr: *httpAddr,
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
		SwaggerDir: *swaggerDir,
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}
