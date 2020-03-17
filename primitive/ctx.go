/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
 * Define the ctx key and value type.
 */
package primitive

import (
	"context"
	"math"
)

type CtxKey int

const (
	method CtxKey = iota
	msgCtx
	orderlyCtx
	concurrentlyCtx

	// method name in  producer
	SendSync   = "SendSync"
	SendOneway = "SendOneway"
	SendAsync  = "SendAsync"
	// method name in consumer
	ConsumerPush = "ConsumerPush"
	ConsumerPull = "ConsumerPull"
)

type ConsumeMessageContext struct {
	ConsumerGroup string
	Msgs          []*MessageExt
	MQ            *MessageQueue
	Success       bool
	Status        string
	// mqTractContext
	Properties map[string]string
}

// WithMethod set call method name
func WithMethod(ctx context.Context, m string) context.Context {
	return context.WithValue(ctx, method, m)
}

// GetMethod get call method name
func GetMethod(ctx context.Context) string {
	return ctx.Value(method).(string)
}

// WithConsumerCtx set ConsumeMessageContext in PushConsumer
func WithConsumerCtx(ctx context.Context, c *ConsumeMessageContext) context.Context {
	return context.WithValue(ctx, msgCtx, c)
}

// GetConsumerCtx get ConsumeMessageContext, only legal in PushConsumer. so should add bool return param indicate
// whether exist.
func GetConsumerCtx(ctx context.Context) (*ConsumeMessageContext, bool) {
	c, exist := ctx.Value(msgCtx).(*ConsumeMessageContext)
	return c, exist
}

type ConsumeOrderlyContext struct {
	MQ                            MessageQueue
	AutoCommit                    bool
	SuspendCurrentQueueTimeMillis int
}

func NewConsumeOrderlyContext() *ConsumeOrderlyContext {
	return &ConsumeOrderlyContext{
		AutoCommit:                    true,
		SuspendCurrentQueueTimeMillis: -1,
	}
}

func WithOrderlyCtx(ctx context.Context, c *ConsumeOrderlyContext) context.Context {
	return context.WithValue(ctx, orderlyCtx, c)
}

func GetOrderlyCtx(ctx context.Context) (*ConsumeOrderlyContext, bool) {
	c, exist := ctx.Value(orderlyCtx).(*ConsumeOrderlyContext)
	return c, exist
}

type ConsumeConcurrentlyContext struct {
	MQ                        MessageQueue
	DelayLevelWhenNextConsume int
	AckIndex                  int32
}

func NewConsumeConcurrentlyContext() *ConsumeConcurrentlyContext {
	return &ConsumeConcurrentlyContext{
		AckIndex: math.MaxInt32,
	}
}

func WithConcurrentlyCtx(ctx context.Context, c *ConsumeConcurrentlyContext) context.Context {
	return context.WithValue(ctx, concurrentlyCtx, c)
}

func GetConcurrentlyCtx(ctx context.Context) (*ConsumeConcurrentlyContext, bool) {
	c, exist := ctx.Value(concurrentlyCtx).(*ConsumeConcurrentlyContext)
	return c, exist
}