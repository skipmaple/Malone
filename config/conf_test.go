// Copyright © 2020. Drew Lee. All rights reserved.

package config

import (
	"github.com/spf13/viper"
	"testing"
)

func Test_bindingConfig(t *testing.T) {
	type args struct {
		cfg    *viper.Viper
		rawVal interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1", args: args{
				cfg:    cfgViper.Sub("dev").Sub("database"),
				rawVal: &Database,
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if Database.Name != "karlmalone" {
				t.Errorf("error-- dbname get: %s, ideal: karlMalone", Database.Name)
			}
		})
	}
}

func Test_viper(t *testing.T) {
	dbName := cfgViper.GetString("dev.database.Name")
	if dbName != "karlmalone" {
		t.Errorf(`error-- dbName: %v, ideal: karlmalone`, dbName)
	}

	logDir := cfgViper.GetString("dev.logger.Dir")
	if logDir != "./log" {
		t.Errorf(`error-- logDir: %v, ideal: ./log`, logDir)
	}

}
