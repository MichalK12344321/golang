package ssh

type VSSHFactoryProvider struct{}

func NewVSSHFactoryProvider() *VSSHFactoryProvider {
	return &VSSHFactoryProvider{}
}

func (p *VSSHFactoryProvider) GetSSHFactory() func() SSH {
	return func() SSH {
		return NewSSHClient()
	}
}

var _ SSHFactoryProvider = new(VSSHFactoryProvider)
