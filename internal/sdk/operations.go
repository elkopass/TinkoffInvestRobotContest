package sdk

import (
	"github.com/elkopass/TinkoffInvestRobotContest/internal/loggy"
	pb "github.com/elkopass/TinkoffInvestRobotContest/internal/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type OperationsInterface interface {
	// Метод получения списка операций по счёту.
	GetOperations(accountID string, from, to *timestamp.Timestamp, state pb.OperationState, figi string) ([]*pb.Operation, error)
	// Метод получения портфеля по счёту.
	GetPortfolio(accountID string) (*pb.PortfolioResponse, error)
	// Метод получения списка позиций по счёту.
	GetPositions(accountID string) (*pb.PositionsResponse, error)
	// Метод получения доступного остатка для вывода средств.
	GetWithdrawLimits(accountID string) (*pb.WithdrawLimitsResponse, error)
}

type OperationsService struct {
	client pb.OperationsServiceClient
}

func NewOperationsService() *OperationsService {
	conn, err := createClientConn()
	if err != nil {
		loggy.GetLogger().Sugar().Fatal(err.Error())
	}

	client := pb.NewOperationsServiceClient(conn)
	return &OperationsService{client: client}
}

func (os OperationsService) GetOperations(accountID string, from, to *timestamp.Timestamp, state pb.OperationState, figi string) ([]*pb.Operation, error) {
	ctx, cancel := createRequestContext()
	defer cancel()

	res, err := os.client.GetOperations(ctx, &pb.OperationsRequest{
		AccountId: accountID,
		From: from,
		To: to,
		State: state,
		Figi: figi,
	})
	if err != nil {
		return nil, err
	}

	return res.Operations, nil
}

func (os OperationsService) GetPortfolio(accountID string) (*pb.PortfolioResponse, error) {
	ctx, cancel := createRequestContext()
	defer cancel()

	res, err := os.client.GetPortfolio(ctx, &pb.PortfolioRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (os OperationsService) GetPositions(accountID string) (*pb.PositionsResponse, error) {
	ctx, cancel := createRequestContext()
	defer cancel()

	res, err := os.client.GetPositions(ctx, &pb.PositionsRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (os OperationsService) GetWithdrawLimits(accountID string) (*pb.WithdrawLimitsResponse, error) {
	ctx, cancel := createRequestContext()
	defer cancel()

	res, err := os.client.GetWithdrawLimits(ctx, &pb.WithdrawLimitsRequest{
		AccountId: accountID,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
