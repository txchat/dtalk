package rpcidgen

import (
	"context"

	"github.com/txchat/dtalk/app/services/generator/generatorclient"
)

type IDGenerator struct {
	client generatorclient.Generator
}

func NewIDGenerator(client generatorclient.Generator) *IDGenerator {
	return &IDGenerator{
		client: client,
	}
}

func (c *IDGenerator) GetID(ctx context.Context) (int64, error) {
	reply, err := c.client.GetID(ctx, &generatorclient.GetIDReq{})
	return reply.GetId(), err
}
