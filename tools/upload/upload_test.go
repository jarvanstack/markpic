package upload

import "testing"

func TestUploaderImpl_Upload(t *testing.T) {
	u := NewUploader()
	type args struct {
		localPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				localPath: "/mnt/c/Users/dengjiawen/Downloads/ddd1.md ",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := u.Upload(tt.args.localPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("UploaderImpl.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("UploaderImpl.Upload() = %v", got)
		})
	}
}
