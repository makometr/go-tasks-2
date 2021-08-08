package main

import (
	"bufio"
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test_readDataFromFile(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readDataFromFile(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("readDataFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readDataFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func Test_doGrep(t *testing.T) {
	type args struct {
		data   []string
		params parametres
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			name: "default Find USD",
			args: args{
				data:   readFile("testFiles/default_1.txt"),
				params: parametres{pattern: "USD"},
			},
			wantOut: strings.Join(readFile("testFiles/default_1_ans.txt"), "\n"),
		},
		{
			name: "default Find last eur (eur$)",
			args: args{
				data:   readFile("testFiles/default_2.txt"),
				params: parametres{pattern: "eur$"},
			},
			wantOut: strings.Join(readFile("testFiles/default_2_ans.txt"), "\n"),
		},
		{
			name: "default_1 Count",
			args: args{
				data:   readFile("testFiles/default_1.txt"),
				params: parametres{pattern: "u", toCount: true},
			},
			wantOut: "3\n",
		},
		{
			name: "default_1 Count_ignore_case",
			args: args{
				data:   readFile("testFiles/default_2.txt"),
				params: parametres{pattern: "usd", toCount: true, ignoreCase: true},
			},
			wantOut: "4\n",
		},
		{
			name: "default_1 Count_ignore_case",
			args: args{
				data:   readFile("testFiles/default_2.txt"),
				params: parametres{pattern: "usd", toCount: true, ignoreCase: true, invert: true},
			},
			wantOut: "2\n",
		},
		{
			name: "fixed no pattern",
			args: args{
				data:   readFile("testFiles/default_2.txt"),
				params: parametres{pattern: "byn$", fixed: true},
			},
			wantOut: "",
		},
		// check output flags -A -B -C -n
		{
			name: "fixed no pattern",
			args: args{
				data:   readFile("testFiles/out.txt"),
				params: parametres{pattern: "bitoc$", after: 2, before: 2},
			},
			wantOut: strings.Join(readFile("testFiles/out_c_ans.txt"), "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			doGrep(tt.args.data, tt.args.params, out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("doGrep() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
