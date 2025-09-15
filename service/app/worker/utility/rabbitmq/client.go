package rabbitmq

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMQClient(ctx context.Context) (*RabbitMQClient, error) {
	address := g.Cfg().MustGet(ctx, "rabbitmq.default.address").String()
	if address == "" {
		return nil, fmt.Errorf("rabbitmq.default.address is empty")
	}

	g.Log().Info(ctx, "rabbitmq.default.address is ", address)

	// 建立链接
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq.default.address is invalid")
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("rabbitmq.default.address is invalid")
	}

	client := &RabbitMQClient{
		conn:    conn,
		channel: channel,
	}

	// 初始化交换机和队列
	err = client.initExchangeAndQueue(ctx)
	if err != nil {
		_ = client.Close()
		return nil, err
	}
	g.Log().Info(ctx, "Rabbitmq客户端初始化成功")
	return client, nil
}

func (r *RabbitMQClient) initExchangeAndQueue(ctx context.Context) error {
	exchangeName := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.orderDelayExchange").String()
	if exchangeName == "" {
		return fmt.Errorf("rabbitmq.default.exchange.orderDelayExchange is empty")
	}
	g.Log().Info(ctx, "rabbitmq.default.exchange.orderDelayExchange is ", exchangeName)

	// 声明延迟交换机
	err := r.channel.ExchangeDeclare(
		exchangeName,        // 交换机名称
		"x-delayed-message", // 交换机类型
		true,                // 持久化
		false,               // 自动删除
		false,               // 内部使用
		false,               // 不等待服务器响应
		amqp.Table{
			"x-delayed-type": "direct", // 指定延迟交换机的底层类型
		},
	)
	if err != nil {
		return fmt.Errorf("rabbitmq.default.exchange.orderDelayExchange is invalid: %v", err)
	}

	// 获取队列名称
	queueName := g.Cfg().MustGet(ctx, "rabbitmq.default.queue.orderTimeoutQueue").String()
	if queueName == "" {
		return fmt.Errorf("rabbitmq.default.queue.orderTimeoutQueue is empty: %v", err)
	}
	g.Log().Info(ctx, "rabbitmq.default.queue.orderTimeoutQueue is ", queueName)

	// 声明队列
	_, err = r.channel.QueueDeclare(
		queueName, // 交换机名称
		true,      // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // 不等待服务器响应
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("rabbitmq.default.queue.orderTimeoutQueue is invalid: %v", err)
	}
	g.Log().Info(ctx, "rabbitmq.default.queue.orderTimeoutQueue is ", queueName)

	// 绑定队列到交换机
	err = r.channel.QueueBind(
		queueName,
		"order.timeout",
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("rabbitmq.default.queue.orderTimeoutQueue is invalid: %v", err)
	}

	g.Log().Info(ctx, "rabbitmq.default.queue.orderTimeoutQueue is ", queueName)
	return nil
}

func (r *RabbitMQClient) Close() error {
	var err error
	if r.channel != nil {
		err = r.channel.Close()
		if err != nil {
			g.Log().Errorf(context.Background(), "rabbitmq.close.channel is invalid: %v", err)
		}
	}
	if r.conn != nil {
		if closeErr := r.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
			g.Log().Errorf(context.Background(), "rabbitmq.close.conn is invalid: %v", closeErr)
		}
	}
	return err
}

func (r *RabbitMQClient) TestConnection(ctx context.Context) error {
	testQueue := "test.connection.queue"
	_, err := r.channel.QueueDeclare(
		testQueue, // 队列名称
		false,     // 不持久化
		true,      // 自动删除
		false,     // 不独占
		false,     // 不等待服务器响应
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("rabbitmq.channel.QueueDeclare is invalid: %v", err)
	}

	err = r.channel.Publish(
		"",        // 使用默认交换机
		testQueue, // 路由键
		false,     // 不强制
		false,     // 不立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("测试信息 " + time.Now().Format(time.RFC3339)),
		})
	if err != nil {
		return fmt.Errorf("rabbitmq.channel.QueueDeclare is invalid: %v", err)
	}
	g.Log().Info(ctx, "rabbitmq.channel.QueueDeclare is ", testQueue)
	return nil
}
