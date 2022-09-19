package httpserver

import (
	"context"
	"fmt"
	mytime "gin-template/pkg/utils/time"
	"gin-template/pkg/utils/waitgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	wg waitgroup.WaitGroupWapper
)

func NewServer(host string, port int, handler http.Handler) *http.Server {
	var addr string
	if host == "" || port == 0 {
		addr = ":8080"
	} else {
		addr = fmt.Sprintf("%s:%d", host, port)
	}

	return &http.Server{
		Addr: addr,
		Handler: handler,
	}
}

func ListenAndServe(server *http.Server) {
	wg.Go(func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		fmt.Println("Server shutdown at", time.Now().Format(mytime.FormatYYYYMMDDhhmmss))
	})
}

type AfterCloseHandler func()

// 监听linux信号，收到信号，停止服务
func WaitForShutdown(server *http.Server, handlers ...AfterCloseHandler) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	select {
	case sig := <-c:
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
			defer cancelFunc()
			if err := server.Shutdown(ctx); err != nil {
				fmt.Printf("An error occurs when Server shut:%v", err)
			}

			for _, handler := range handlers {
				handler()
			}
		}
	}

	wg.Wait()
}
