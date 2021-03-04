package BingoApi

import (
	"os"
	"testing"
)

func TestBingoApi_NewsSearch(t *testing.T) {
	ClientKey := os.Getenv("CLIENT_KEY")

	type fields struct {
		ClientKey string
	}
	type args struct {
		q string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test basic call", fields{ClientKey}, args{"microsoft"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BingoApi{
				ClientKey: tt.fields.ClientKey,
			}
			got, err := b.NewsSearch(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewsSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || len(got.Value) <= 0 {
				t.Errorf("NewsSearch() got = %v, want %v", got, tt.args.q)
			}
		})
	}
}
