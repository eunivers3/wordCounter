package service

import (
	"reflect"
	"testing"
)

func TestService_CountWords(t *testing.T) {
	type fields struct {
		parser Parser
	}
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				parser: tt.fields.parser,
			}
			got, err := s.CountWords(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountWords() got = %v, want %v", got, tt.want)
			}
		})
	}
}
