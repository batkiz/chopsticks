package chopsticks

import (
	"math"
	"testing"
)

func TestError(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Error(tt.args.a); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHumanReadableFileSize(t *testing.T) {
	type args struct {
		length int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "0 ", args: args{length: 0}, want: "0 B"},
		{name: "27 ", args: args{length: 27}, want: "27 B"},
		{name: "999 ", args: args{length: 999}, want: "999 B"},
		{name: "1000 ", args: args{length: 1000}, want: "1000 B"},
		{name: "1023 ", args: args{length: 1023}, want: "1023 B"},
		{name: "1024 ", args: args{length: 1024}, want: "1.0 KiB"},
		{name: "1728 ", args: args{length: 1728}, want: "1.7 KiB"},
		{name: "1855425871872 ", args: args{length: 1855425871872}, want: "1.7 TiB"},
		{name: "math.MaxInt64 ", args: args{length: math.MaxInt64}, want: "8.0 EiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanReadableFileSize(tt.args.length); got != tt.want {
				t.Errorf("HumanReadableFileSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Info(tt.args.a); got != tt.want {
				t.Errorf("Info() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuccess(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Success(tt.args.a); got != tt.want {
				t.Errorf("Success() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Warn(tt.args.a); got != tt.want {
				t.Errorf("Warn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCacheDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCacheDir(); got != tt.want {
				t.Errorf("getCacheDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConfig(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfig(tt.args.name); got != tt.want {
				t.Errorf("getConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getConfigHome(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfigHome(); got != tt.want {
				t.Errorf("getConfigHome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGlobalDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGlobalDir(); got != tt.want {
				t.Errorf("getGlobalDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getScoopDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getScoopDir(); got != tt.want {
				t.Errorf("getScoopDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadConfig(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadConfig(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loadConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setConfig(t *testing.T) {
	type args struct {
		name  string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setConfig(tt.args.name, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("setConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("setConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
