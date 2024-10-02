package main

import (
	"context"
	"fmt"
	"os"
	"time"

	std_msgs_msg "paultio/ros-go/msgs/std_msgs/msg"
	std_srvs_srv "paultio/ros-go/msgs/std_srvs/srv"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

func handle_service(info *rclgo.ServiceInfo, req *std_srvs_srv.Empty_Request, sender std_srvs_srv.EmptyServiceResponseSender) {
	fmt.Println("#### SERVICE BEGIN ####")
	fmt.Println("Request received")
	fmt.Println("#### SERVICE END ####")
	resp := std_srvs_srv.Empty_Response{}
	err := sender.SendResponse(&resp)
	if err != nil {
		fmt.Println("Failed to send response:", err)
	}
}

func run() error {
	rclArgs, restArgs, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}

	// spinCtx, cancelSpin := context.WithCancel(context.Background())
	// defer cancelSpin()

	// ctx, err := rclgo.NewContext(0, nil)
	// if err != nil {
	// 	return fmt.Errorf("error while creating context: %v", err)
	// }

	// node, err := ctx.NewNode("service_node", "")
	// if err != nil {
	// 	return fmt.Errorf("error while creating node: %v", err)
	// }

	// // _, err = example_srvs_srv.NewAddTwoIntsService(node, "add", nil, func(info *rclgo.ServiceInfo, req *example_srvs_srv.AddTwoInts_Request, sender example_srvs_srv.AddTwoIntsServiceResponseSender) {
	// // 	res := example_srvs_srv.NewAddTwoInts_Response()
	// // 	res.Sum = req.A + req.B
	// // 	sender.SendResponse(res)
	// // })

	// _, err = std_srvs_srv.NewEmptyService(node, "call_me", nil, handle_service)

	// if err != nil {
	// 	return fmt.Errorf("error occured starting service: %v", err)
	// }

	// go ctx.Spin(spinCtx)

	// select {}

	spinCtx, cancelSpin := context.WithCancel(context.Background())
	defer cancelSpin()

	ctx, err := rclgo.NewContext(0, nil)
	if err != nil {
		return fmt.Errorf("error while creating context: %v", err)
	}

	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()

	node, err := ctx.NewNode("publisher", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	pub, err := std_msgs_msg.NewStringPublisher(node, "hello", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}
	defer pub.Close()

	_, err = std_srvs_srv.NewEmptyService(node, "call_me", nil, handle_service)
	if err != nil {
		return fmt.Errorf("failed to create service: %v", err)
	}

	msg := std_msgs_msg.NewString()
	if len(restArgs) > 0 {
		msg.Data = restArgs[0]
	} else {
		msg.Data = "gopher"
	}

	go ctx.Spin(spinCtx)

	for {
		node.Logger().Infof("Publishing: %#v", msg)
		if err := pub.Publish(msg); err != nil {
			return fmt.Errorf("failed to publish: %v", err)
		}
		time.Sleep(time.Second)

	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
