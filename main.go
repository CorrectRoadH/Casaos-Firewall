//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate types,server,spec -package codegen api/local_storage/openapi.yaml > codegen/local_storage_api.go"
//go:generate bash -c "mkdir -p codegen/message_bus && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4 -generate types,client -package message_bus https://raw.githubusercontent.com/IceWhaleTech/CasaOS-MessageBus/main/api/message_bus/openapi.yaml > codegen/message_bus/api.go"

package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/CorrectRoadH/CasaOS-Firewall/common"
	"github.com/CorrectRoadH/CasaOS-Firewall/pkg/config"
	"github.com/CorrectRoadH/CasaOS-Firewall/pkg/sqlite"
	"github.com/CorrectRoadH/CasaOS-Firewall/route"
	"github.com/CorrectRoadH/CasaOS-Firewall/service"
	"github.com/IceWhaleTech/CasaOS-Common/model"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/robfig/cron/v3"
)

const localhost = "127.0.0.1"

var (
	commit = "private build"
	date   = "private build"
)

func init() {
	configFlag := flag.String("c", "", "config address")
	dbFlag := flag.String("db", "", "db path")

	versionFlag := flag.Bool("v", false, "version")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("v%s\n", common.Version)
		os.Exit(0)
	}

	println("git commit:", commit)
	println("build date:", date)

	config.InitSetup(*configFlag)

	logger.LogInit(config.AppInfo.LogPath, config.AppInfo.LogSaveName, config.AppInfo.LogFileExt)

	if len(*dbFlag) == 0 {
		*dbFlag = config.AppInfo.DBPath
	}

	sqliteDB := sqlite.GetGlobalDB(*dbFlag)

}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// sendStorageStats()

	crontab := cron.New(cron.WithSeconds())
	// if _, err := crontab.AddFunc("@every 5s", sendStorageStats); err != nil {
	// 	logger.Error("crontab add func error", zap.Error(err))
	// }

	crontab.Start()
	defer crontab.Stop()

	listener, err := net.Listen("tcp", net.JoinHostPort(localhost, "0"))
	if err != nil {
		panic(err)
	}

	// register at gateway
	apiPaths := []string{
		route.V2APIPath,
		route.V2DocPath,
	}

	for _, apiPath := range apiPaths {
		err = service.MyService.Gateway().CreateRoute(&model.Route{
			Path:   apiPath,
			Target: "http://" + listener.Addr().String(),
		})

		if err != nil {
			panic(err)
		}
	}

}
