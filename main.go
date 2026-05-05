package main

import (
	"fmt"
	"myproj/api"
	"myproj/info"
)

func main() {
	httpHandlers := api.NewHTTPHandlers(info.MyCompany);
	httpServer := api.NewHTTPServer(httpHandlers);

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err);
	}
}