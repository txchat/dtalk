package service

import (
	"bytes"
	"context"
	"github.com/rs/zerolog"
	"github.com/txchat/imparse"
	"github.com/txchat/imparse/chat"
)

func (s *Service) Push(ctx context.Context, key, from string, body []byte) (int64, uint64, error) {
	log := zerolog.Ctx(ctx).With().Str("Uid", from).Str("ConnId", key).Logger()

	log.Debug().Msg("start create frame")
	frame, err := s.parser.NewFrame(key, from, bytes.NewReader(body))
	if err != nil {
		log.Error().Stack().Err(err).Msg("NewFrame error")
		return 0, 0, err
	}
	log.Debug().Msg("start check frame")
	if err := s.rcpAnswer.Check(ctx, &checker, frame); err != nil {
		log.Error().Stack().Err(err).Msg("Check error")
		return 0, 0, err
	}
	log.Debug().Msg("start frame filter")
	createTime, err := s.rcpAnswer.Filter(ctx, frame)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Filter error")
		return 0, 0, err
	}
	log.Debug().Msg("start frame transport")
	if err := s.rcpAnswer.Transport(ctx, frame); err != nil {
		log.Error().Stack().Err(err).Msg("Transport error")
		return 0, 0, err
	}
	log.Debug().Msg("start frame ack")
	mid, err := s.rcpAnswer.Ack(ctx, frame)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Ack error")
		return 0, 0, err
	}
	log.Debug().Msg("deal msg send success")
	return mid, createTime, nil
}

//gRPC调用的内部推送通道，不经过check
func (s *Service) InnerPush(ctx context.Context, key, from, target string, pushType imparse.Channel, body []byte) (int64, error) {
	log := zerolog.Ctx(ctx).With().Str("Uid", from).Str("ConnId", key).Str("Push Type", pushType.String()).Logger()

	log.Debug().Msg("start create frame")
	frame, err := s.parser.NewFrame(key, from, bytes.NewReader(body), chat.WithTarget(target), chat.WithTransmissionMethod(pushType))
	if err != nil {
		log.Error().Stack().Err(err).Msg("NewFrame error")
		return 0, err
	}

	log.Debug().Msg("start frame filter")
	if _, err := s.rcpAnswer.Filter(ctx, frame); err != nil {
		log.Error().Stack().Err(err).Msg("Filter error")
		return 0, err
	}
	log.Debug().Msg("start frame transport")
	if err := s.rcpAnswer.Transport(ctx, frame); err != nil {
		log.Error().Stack().Err(err).Msg("Transport error")
		return 0, err
	}
	log.Debug().Msg("start frame ack")
	mid, err := s.rcpAnswer.Ack(ctx, frame)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Ack error")
		return 0, err
	}
	return mid, nil
}
