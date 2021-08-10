package main

import (
	"bufio"
	"os"
)

func readFileOrPanic(filename string) []string {
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

// func Test_doGrep(t *testing.T) {
// 	type args struct {
// 		data   []string
// 		params parametres
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantOut string
// 	}{
// 		{
// 			name: "default Find USD",
// 			args: args{
// 				data:   readFileOrPanic("testFiles/default_1.txt"),
// 				params: parametres{pattern: "USD"},
// 			},
// 			wantOut: strings.Join(readFileOrPanic("testFiles/default_1_ans.txt"), "\n"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			out := &bytes.Buffer{}
// 			doGrep(tt.args.data, tt.args.params, out)
// 			if gotOut := out.String(); gotOut != tt.wantOut {
// 				t.Errorf("doGrep() = %v, want %v", gotOut, tt.wantOut)
// 			}
// 		})
// 	}
// }
