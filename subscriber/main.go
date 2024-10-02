package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	std_msgs_msg "paultio/ros-go/msgs/std_msgs/msg"
	std_srvs_srv "paultio/ros-go/msgs/std_srvs/srv"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

var client *std_srvs_srv.EmptyClient

func call_service() {
	resp, msg_info, err := client.Send(context.Background(), &std_srvs_srv.Empty_Request{})
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	fmt.Println("#### RESP BEGIN ####")
	fmt.Println(resp)
	fmt.Println(msg_info)
	fmt.Print("#### RESP END ####\n\n")
}

func parse_string(msg *std_msgs_msg.String, info *rclgo.MessageInfo, err error) {
	if err != nil {
		return
	}
	fmt.Println("#### MSG BEGIN ####")
	fmt.Println(msg.Data)
	fmt.Print("#### MSG END ####\n\n")

	go call_service()

}

func run() error {
	rclArgs, _, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}

	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()

	node, err := rclgo.NewNode("subscriber", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	sub, err := std_msgs_msg.NewStringSubscription(
		node,
		"hello",
		nil,
		parse_string,
	)
	if err != nil {
		return fmt.Errorf("failed to create subscriber: %v", err)
	}
	defer sub.Close()

	client, err = std_srvs_srv.NewEmptyClient(node, "call_me", nil)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ws, err := rclgo.NewWaitSet()
	if err != nil {
		return fmt.Errorf("failed to create waitset: %v", err)
	}
	defer ws.Close()
	ws.AddSubscriptions(sub.Subscription)
	return ws.Run(ctx)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
