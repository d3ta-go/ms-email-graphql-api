package schema

import "testing"

func TestLoadGraphQLSchema(t *testing.T) {
	type args struct {
		schemaDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Test OK",
			args:    args{schemaDir: "./schema.graphql"},
			wantErr: false,
		},
		{
			name:    "Test Error",
			args:    args{schemaDir: "./schema.graphql/notfound"},
			wantErr: true,
		},
		{
			name:    "Test Error 2",
			args:    args{schemaDir: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := LoadGraphQLSchema(tt.args.schemaDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadGraphQLSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
