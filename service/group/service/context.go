package service

//type SvcCtxOpt func(ctx context.Context) (context.Context, error)

//// SetContext 设置业务 ctx, opt 需要按照顺序填写
//func (s *Service) SetContext(ctx context.Context, options ...SvcCtxOpt) (context.Context, error) {
//	var err error
//	for _, opt := range options {
//		ctx, err = opt(ctx)
//		if err != nil {
//			return nil, err
//		}
//	}
//	return ctx, nil
//}
//
//func (s *Service) WithGroupInfoCtxOpt(ctx context.Context) (context.Context, error) {
//	groupId, err := biz.NewGroupIdFromContext(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	group, err := s.GetGroupInfoByGroupId(ctx, groupId)
//	if err != nil {
//		return nil, err
//	}
//	ctx = group.WithContext(ctx)
//
//	return ctx, nil
//}
//
//func (s *Service) WithPersonCtxOpt(ctx context.Context) (context.Context, error) {
//	groupId, err := biz.NewGroupIdFromContext(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	memberId, _ := ctx.Value(api.Address).(string)
//	if memberId == "" {
//		return ctx, nil
//	}
//
//	person, err := s.GetPersonByMemberIdAndGroupId(ctx, memberId, groupId)
//	if err != nil {
//		return nil, err
//	}
//	ctx = person.WithContext(ctx)
//
//	return ctx, nil
//}
//
//func (s *Service) WithAdminCtxOpt(ctx context.Context) (context.Context, error) {
//	person, err := biz.NewGroupMemberFromContext(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := person.IsAdmin(); err != nil {
//		return nil, err
//	}
//
//	return ctx, nil
//}
//
//func (s *Service) WithOwnerCtxOpt(ctx context.Context) (context.Context, error) {
//	person, err := biz.NewGroupMemberFromContext(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := person.IsOwner(); err != nil {
//		return nil, err
//	}
//
//	return ctx, nil
//}
