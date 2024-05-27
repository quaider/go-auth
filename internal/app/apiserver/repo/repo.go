package repo

var client Factory

// Factory 定义了整个服务的repo合集（即facades）
type Factory interface {
	Users() UserRepo
}

func Client() Factory {
	return client
}

func SetClient(factory Factory) {
	client = factory
}
