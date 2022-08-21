package service

//import (
//	"reflect"
//	"testing"
//)
//
//func TestService_GeocodeOne(t *testing.T) {
//	type args struct {
//		url string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    map[string]int
//		wantErr bool
//	}{
//		{
//			name: "should succeed and return nil",
//			args: args{url: "https:/google.com"},
//			want: map[string]int{},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := CountWords(tt.args.url)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("CountWords() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("CountWords() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
