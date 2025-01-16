package server

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"github.com/rodeorm/keeper/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockUserStorager(ctrl)

	storage.EXPECT().VerifyUserOTP(gomock.Any(), gomock.Any(), gomock.Any()).Return(true).AnyTimes()

	grpcSrv := grpcServer{cfg: &cfg.Server{UserStorager: storage}}

	grpcSrv.srv = grpc.NewServer()
	defer grpcSrv.srv.Stop()

	proto.RegisterKeeperServiceServer(grpcSrv.srv, &grpcSrv)

	go func() {
		lis, err := net.Listen("tcp", ":3200")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := grpcSrv.srv.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}()

	type test struct {
		name         string
		req          *proto.VerifyRequest
		expectedCode codes.Code
	}

	ts := test{name: "проверка на успешную авторизацию",
		req:          &proto.VerifyRequest{Login: "login", OTP: "success"},
		expectedCode: codes.OK,
	}

	conn, err := grpc.NewClient(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewKeeperServiceClient(conn)
	ctx := context.Background()
	var header metadata.MD

	resp, err := c.Verify(ctx, ts.req, grpc.Header(&header))
	if err != nil {
		log.Println("Ошибка при вызове Verify:", err)
		t.FailNow()
	}
	st, _ := status.FromError(err)
	log.Printf("Результаты %s: %v", ts.name, resp)

	assert.Equal(t, ts.expectedCode, st.Code())

}
