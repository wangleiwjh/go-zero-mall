package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	fmt.Println("Starting services...")

	services := []Service{
		{"user/rpc", "go run user.go"},
		{"user/api", "go run user.go"},
		{"product/rpc", "go run product.go"},
		{"product/api", "go run product.go"},
		{"order/rpc", "go run order.go"},
		{"order/api", "go run order.go"},
		{"pay/rpc", "go run pay.go"},
		{"pay/api", "go run pay.go"},
	}

	for _, service := range services {
		if err := startService(service); err != nil {
			fmt.Printf("Error starting service %s: %v\n", service.Dir, err)
			return
		}
	}

	fmt.Println("All services started successfully.")
	select {} // 阻止主程序退出，保持服务运行
}

// Service 定义服务的信息
type Service struct {
	Dir string
	Cmd string
}

// startService 启动单个服务
func startService(service Service) error {
	fmt.Printf("Starting service in directory: %s\n", service.Dir)
	cmd := exec.Command("cmd", "/C", service.Cmd)
	cmd.Dir = fmt.Sprintf("service/%s", service.Dir)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start service %s: %w", service.Dir, err)
	}
	fmt.Printf("Service started: %s\n", service.Dir)

	// 等待一段时间，确保前一个服务启动完成后再启动下一个服务
	time.Sleep(2 * time.Second)
	return nil
}
