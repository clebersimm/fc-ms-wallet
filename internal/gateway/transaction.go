package gateway

import "github.com/clebersimm/fc-ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transation *entity.Transaction) error
}
