// Copyright © 2020. Drew Lee. All rights reserved.

package util

import "testing"

func TestValidatePwd(t *testing.T) {
	type args struct {
		plainPwd string
		salt     string
		pwd      string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"111222", "008081", "00b7691d86d96aebd21dd9e138f90840"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidatePwd(tt.args.plainPwd, tt.args.salt, tt.args.pwd); got != tt.want {
				t.Errorf("ValidatePwd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePwd(t *testing.T) {
	type args struct {
		plainPwd string
		salt     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"111222", ""}, "00b7691d86d96aebd21dd9e138f90840"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePwd(tt.args.plainPwd, tt.args.salt); got != tt.want {
				t.Errorf("MakePwd() = %v, want %v", got, tt.want)
			}
		})
	}
}
