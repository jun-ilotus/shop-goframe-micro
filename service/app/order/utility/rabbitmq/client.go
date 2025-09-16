package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"sync"
	"time"
)

type OrderRabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	mutex   sync.Mutex
}

var (
	instance *OrderRabbitMQClient
	once     sync.Once
)

func GetOrderRabbitMQClient(ctx context.Context) (*OrderRabbitMQClient, error) {
	var err error
	once.Do(func() {
		instance, err = newOrderRabbitMQClient(ctx)
	})

	if err != nil {
		return nil, err
	}

	// 检查连接是否已关闭，如果已关闭则重新创建
	if instance.conn.IsClosed() {
		instance.mutex.Lock()
		defer instance.mutex.Unlock()

		// 双重检查，防止多个goroutine同时进入此逻辑
		if instance.conn.IsClosed() {
			instance, err = newOrderRabbitMQClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("fail to create new OrderRabbitMQClient: %v", err)
			}
		}
	}
	return instance, nil
}

func newOrderRabbitMQClient(ctx context.Context) (*OrderRabbitMQClient, error) {
	address := g.Cfg().MustGet(ctx, "rabbitmq.default.address").String()
	if address == "" {
		return nil, fmt.Errorf("rabbitmq.default.address is empty")
	}
	g.Log().Info(ctx, "rabbitmq.address is ", address)

	// 建立连接
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq.default.address is invalid: %v", err)
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("rabbitmq.default.address is invalid: %v", err)
	}

	// 设置QoS （服务质量）
	prefetchCount := g.Cfg().MustGet(ctx, "rabbitmq.default.consumer.prefetchCount").Int()
	if prefetchCount <= 0 {
		prefetchCount = 1
	}
	err = channel.Qos(prefetchCount, 0, false)
	if err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("rabbitmq.default.consumer.prefetchCount is invalid: %v", err)
	}

	client := &OrderRabbitMQClient{
		conn:    conn,
		channel: channel,
	}

	// 初始化交换机和队列
	err = client.initExchangeAndQueue(ctx)
	if err != nil {
		_ = client.Close()
		return nil, err
	}
	g.Log().Info(ctx, "rabbitmq.default.consumer.prefetchCount is ", prefetchCount)
	return client, nil
}

func (r *OrderRabbitMQClient) Close() error {
	var err error
	if r.conn != nil {
		err = r.conn.Close()
		if err != nil {
			g.Log().Errorf(context.Background(), "关闭Rabbitmq通道失败：%v", err)
		}
	}
	if r.conn != nil {
		if closeErr := r.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
			g.Log().Errorf(context.Background(), "关闭Rabbitmq连接失败：%v", closeErr)
		}
	}
	return err
}

func (r *OrderRabbitMQClient) initExchangeAndQueue(ctx context.Context) error {
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

	queueName := g.Cfg().MustGet(ctx, "rabbitmq.default.queue.orderTimeoutQueue").String()
	if queueName == "" {
		return fmt.Errorf("rabbitmq.default.queue.orderTimeoutQueue is empty")
	}
	g.Log().Info(ctx, "rabbitmq.default.queue.orderTimeoutQueue is ", queueName)

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

func (r *OrderRabbitMQClient) SendOrderTimeoutMessage(ctx context.Context, orderId int32, delayMs int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 获取交换机名称
	exchangeName := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.orderDelayExchange").String()
	if exchangeName == "" {
		return fmt.Errorf("rabbitmq.default.exchange.orderDelayExchange is empty")
	}

	message := map[string]interface{}{
		"orderId":   orderId,
		"type":      "order_timeout",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("rabbitmq.default.exchange.orderDelayExchange is invalid: %v", err)
	}

	err = r.channel.Publish(
		exchangeName,
		"order.timeout",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Headers: amqp.Table{
				"x-delay": "delayMs", // 延迟时间，毫秒
			},
			DeliveryMode: amqp.Persistent,
		})
	if err != nil {
		return fmt.Errorf("rabbitmq.default.exchange.orderTimeoutQueue is invalid: %v", err)
	}
	g.Log().Infof(ctx, "订单超时信息发送成功，订单ID: %d, 延迟: %d毫秒", orderId, delayMs)
	return nil
}

func GetOrderTimeoutDelay(ctx context.Context) int {
	timeout := g.Cfg().MustGet(ctx, "business.orderTimeout").Int()
	if timeout <= 0 {
		timeout = 30 * 60 * 1000
	}
	return timeout
}

func SendOrderTimeoutMessageStatic(ctx context.Context, orderId int32, delayMs int) error {
	client, err := GetOrderRabbitMQClient(ctx)
	if err != nil {
		return err
	}
	return client.SendOrderTimeoutMessage(ctx, orderId, delayMs)
}
