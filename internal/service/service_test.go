package service

import (
	googleClientMock "github.com/eunicebjm/gc/internal/mocks/google/geocoder"
	"github.com/eunicebjm/gc/pkg/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type Mocks struct {
	serviceToken string
	googleClient *googleClientMock.MockClient
}

func createMocks(ctrl *gomock.Controller) Mocks {
	return Mocks{
		serviceToken: "service_token",
		googleClient: googleClientMock.NewMockClient(ctrl),
	}
}

func initMockService(t *testing.T) (Mocks, *Service) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := createMocks(ctrl)

	svc, err := NewService(
		mocks.googleClient,
	)
	require.NoError(t, err)
	assert.NotEmpty(t, svc)

	return mocks, svc
}

func TestService_GeocodeOne(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Result
		wantErr bool
	}{
		{
			name: "should succeed and return nil",
			args: args{address: "some-address"},
			want: models.Result{
				ID:               "some-id",
				FormattedAddress: "some-address",
				Geometry: models.Geometry{Location: models.Location{
					Latitude:  0,
					Longitude: 0,
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock, svc := initMockService(t)

			gomock.InOrder(
				mock.googleClient.EXPECT().
					GeocodeOne(tt.args.address).
					Return(tt.want, nil),
			)

			got, err := svc.GeocodeOne(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeocodeOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeocodeOne() got = %v, want %v", got, tt.want)
			}
		})
	}
}
