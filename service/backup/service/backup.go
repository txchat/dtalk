package service

import (
	xerror "github.com/txchat/dtalk/pkg/error"
	"github.com/txchat/dtalk/service/backup/model"
	"time"
)

func (s *Service) EditMnemonic(addr, mne string) error {
	err := s.dao.UpdateMnemonic(&model.AddrBackup{
		Address:    addr,
		Mnemonic:   mne,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	return nil
}

func (s *Service) AddressRetrieve(addr string) (*model.AddrBackup, error) {
	item, err := s.dao.Query(model.Address, addr)
	if err != nil {
		return nil, xerror.NewError(xerror.QueryFailed)
	}
	return item, nil
}

func (s *Service) PhoneIsBound(area, phone string) (bool, error) {
	item, err := s.dao.Query(model.Phone, phone)
	if err != nil {
		return false, xerror.NewError(xerror.QueryFailed)
	}
	if item == nil {
		itemRelate, err := s.dao.QueryRelate(model.Phone, phone)
		if err != nil {
			return false, xerror.NewError(xerror.QueryFailed)
		}
		return itemRelate != nil, nil
	}
	return item != nil, nil
}

func (s *Service) SendPhoneCode(codeType, phone string) error {
	//发送短信验证码
	params := map[string]string{
		model.ParamMobile:   phone,
		model.ParamCodeType: s.cfg.SMS.CodeTypes[codeType],
	}
	_, err := s.smsValidate.Send(params)
	return err
}

func (s *Service) PhoneBinding(addr, area, phone, code, mnemonic string) error {
	params := map[string]string{
		model.ParamMobile:   phone,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.SMS.CodeTypes[model.Quick],
	}
	//验证
	err := s.smsValidate.ValidateCode(params)
	if err != nil {
		return xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	//绑定
	err = s.dao.UpdateAddrBackup(model.Phone, &model.AddrBackup{
		Address:    addr,
		Area:       area,
		Phone:      phone,
		Mnemonic:   mnemonic,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	return nil
}

func (s *Service) PhoneBindingV2(addr, area, phone, code, mnemonic string) error {
	params := map[string]string{
		model.ParamMobile:   phone,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.SMS.CodeTypes[model.Bind],
	}
	//验证
	err := s.smsValidate.ValidateCode(params)
	if err != nil {
		return xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	//绑定
	err = s.dao.UpdateAddrBackup(model.Phone, &model.AddrBackup{
		Address:    addr,
		Area:       area,
		Phone:      phone,
		Mnemonic:   mnemonic,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	return nil
}

//关联手机号
func (s *Service) PhoneRelate(addr, area, phone, mnemonic string) error {
	//绑定
	err := s.dao.UpdateAddrRelate(model.Phone, &model.AddrRelate{
		Address:    addr,
		Area:       area,
		Phone:      phone,
		Mnemonic:   mnemonic,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	return nil
}

func (s *Service) PhoneExport(addr, area, phone, code string) (bool, error) {
	params := map[string]string{
		model.ParamMobile:   phone,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.SMS.CodeTypes[model.Export],
	}
	//验证
	err := s.smsValidate.ValidateCode(params)
	if err != nil {
		return false, xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	//查询
	//item, err := s.dao.Query(model.Phone, phone)
	//if err != nil {
	//	return false, xerror.NewError(xerror.QueryFailed)
	//}
	return true, nil
}

func (s *Service) PhoneRetrieve(area, phone, code string) (*model.AddrBackup, error) {
	//验证
	params := map[string]string{
		model.ParamMobile:   phone,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.Email.CodeTypes[model.Quick],
	}
	err := s.smsValidate.ValidateCode(params)
	if err != nil {
		return nil, xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	item, err := s.dao.Query(model.Phone, phone)
	if err != nil {
		return nil, xerror.NewError(xerror.QueryFailed)
	}
	if item == nil {
		//查询 是否关联
		item, err := s.dao.QueryRelate(model.Phone, phone)
		if err != nil || item == nil {
			return nil, err
		}
		newAddr := &model.AddrBackup{
			Address:    item.Address,
			Area:       item.Area,
			Phone:      item.Phone,
			Email:      item.Email,
			Mnemonic:   item.Mnemonic,
			PrivateKey: item.PrivateKey,
			UpdateTime: item.UpdateTime,
			CreateTime: item.CreateTime,
		}
		err = s.dao.UpdateAddrBackup(model.Phone, newAddr)
		if err != nil {
			return nil, err
		}
		return newAddr, nil
	}
	return item, nil
}

//--------------------email-------------//
func (s *Service) EmailIsBound(email string) (bool, error) {
	item, err := s.dao.Query(model.Email, email)
	if err != nil {
		return false, xerror.NewError(xerror.QueryFailed)
	}
	return item != nil, nil
}

func (s *Service) EmailBinding(addr, email, code, mnemonic string) error {
	params := map[string]string{
		model.ParamEmail:    email,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.Email.CodeTypes[model.Quick],
	}
	//验证
	err := s.emailValidate.ValidateCode(params)
	if err != nil {
		return xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	//绑定
	err = s.dao.UpdateAddrBackup(model.Email, &model.AddrBackup{
		Address:    addr,
		Email:      email,
		Mnemonic:   mnemonic,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	return nil
}

func (s *Service) EmailBindingV2(addr, email, code, mnemonic string) error {
	params := map[string]string{
		model.ParamEmail:    email,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.Email.CodeTypes[model.Bind],
	}
	//验证
	err := s.emailValidate.ValidateCode(params)
	if err != nil {
		return xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	//绑定
	err = s.dao.UpdateAddrBackup(model.Email, &model.AddrBackup{
		Address:    addr,
		Email:      email,
		Mnemonic:   mnemonic,
		UpdateTime: time.Now(),
	})
	if err != nil {
		return xerror.NewError(xerror.ExecFailed)
	}
	return nil
}

func (s *Service) EmailExport(addr, email, code string) (bool, error) {
	params := map[string]string{
		model.ParamEmail:    email,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.Email.CodeTypes[model.Export],
	}
	//验证
	err := s.emailValidate.ValidateCode(params)
	if err != nil {
		return false, xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	////查询
	//item, err := s.dao.Query(model.Email, email)
	//if err != nil {
	//	return false, xerror.NewError(xerror.QueryFailed)
	//}
	return true, nil
}

func (s *Service) SendEmailCode(codeType, email string) error {
	//发送短信验证码
	params := map[string]string{
		model.ParamEmail:    email,
		model.ParamCodeType: s.cfg.Email.CodeTypes[codeType],
	}
	_, err := s.emailValidate.Send(params)
	return err
}

func (s *Service) EmailRetrieve(email, code string) (*model.AddrBackup, error) {
	//验证
	params := map[string]string{
		model.ParamEmail:    email,
		model.ParamCode:     code,
		model.ParamCodeType: s.cfg.Email.CodeTypes[model.Quick],
	}
	err := s.emailValidate.ValidateCode(params)
	if err != nil {
		return nil, xerror.NewError(xerror.VerifyCodeError).SetExtMessage(err.Error())
	}
	item, err := s.dao.Query(model.Email, email)
	if err != nil {
		return nil, xerror.NewError(xerror.QueryFailed)
	}
	return item, nil
}

func (s *Service) GetAddress(req *model.GetAddressRequest) (res *model.GetAddressResponse, err error) {
	defer func() {
		if err != nil &&
			err.Error() != xerror.NewError(xerror.ParamsError).Error() &&
			err.Error() != xerror.NewError(xerror.QueryNotExist).Error() {
			s.log.Error("GetAddress", "err", err, "req", req)
		}
	}()
	query := model.NewQuery(req.Query)
	addrBackup, err := s.dao.Query(uint32(query.Tp), query.Query)
	if err != nil {
		if err == model.ErrQueryType {
			return nil, xerror.NewError(xerror.ParamsError)
		}
		return nil, xerror.NewError(xerror.QueryFailed).SetExtMessage(err.Error())
	}
	if addrBackup == nil {
		return nil, xerror.NewError(xerror.QueryNotExist)
	}
	return &model.GetAddressResponse{Address: addrBackup.Address}, nil
}
