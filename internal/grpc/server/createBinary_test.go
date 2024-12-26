package server

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/keeper/internal/cfg"
	"github.com/rodeorm/keeper/internal/grpc/meta"
	"github.com/rodeorm/keeper/internal/grpc/proto"
	"github.com/rodeorm/keeper/mocks"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestCreateBinary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageBinary := mocks.NewMockBinaryStorager(ctrl)
	storageBinary.EXPECT().AddBinaryByUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	storageUser := mocks.NewMockUserStorager(ctrl)
	storageUser.EXPECT().VerifyUserOTP(gomock.Any(), gomock.Any(), gomock.Any()).Return(true).AnyTimes()

	grpcSrv := grpcServer{cfg: &cfg.Server{BinaryStorager: storageBinary, UserStorager: storageUser}}

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
		req          *proto.CreateBinaryRequest
		expectedCode codes.Code
	}

	ts := test{
		req: &proto.CreateBinaryRequest{
			Binary: &proto.Binary{
				Value: []byte("это тест"),
				Meta:  "Это тестовая мета",
			},
		},
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
	c.Verify(ctx, &proto.VerifyRequest{}, grpc.Header(&header))
	token := meta.GetTokenFromMeta(header)
	ctx = meta.AddTokenToCtx(ctx, token)
	_, err = c.CreateBinary(ctx, ts.req)
	st, _ := status.FromError(err)

	assert.Equal(t, ts.expectedCode, st.Code())

}
