package parser

import (
	"errors"

	"github.com/AplaProject/go-apla/packages/conf/syspar"
	"github.com/AplaProject/go-apla/packages/consts"
	"github.com/AplaProject/go-apla/packages/service"
	"github.com/AplaProject/go-apla/packages/utils"
	"github.com/AplaProject/go-apla/packages/utils/tx"
)

var (
	messageNetworkStopping = "Attention! The network is stopped!"

	errNetworkStopping = errors.New("Network is stopping")
)

type StopNetworkParser struct {
	*Parser

	cert *utils.Cert
}

func (p *StopNetworkParser) Init() error {
	return nil
}

func (p *StopNetworkParser) Validate() error {
	if err := p.validate(); err != nil {
		p.GetLogger().WithError(err).Error("validating tx")
		return err
	}

	return nil
}

func (p *StopNetworkParser) validate() error {
	data := p.TxPtr.(*consts.StopNetwork)

	cert, err := utils.ParseCert(data.StopNetworkCert)
	if err != nil {
		return err
	}

	fbdata, err := syspar.GetFirstBlockData()
	if err != nil {
		return err
	}

	if err = cert.Validate(fbdata.StopNetworkCertBundle); err != nil {
		return err
	}

	p.cert = cert
	return nil
}

func (p *StopNetworkParser) Action() error {
	// Allow execute transaction, if the certificate was used
	if p.cert.EqualBytes(consts.UsedStopNetworkCerts...) {
		return nil
	}

	// Set the node in a pause state
	service.PauseNodeActivity(service.PauseTypeStopingNetwork)

	p.GetLogger().Warn(messageNetworkStopping)
	return errNetworkStopping
}

func (p *StopNetworkParser) Rollback() error {
	return nil
}

func (p StopNetworkParser) Header() *tx.Header {
	return nil
}
