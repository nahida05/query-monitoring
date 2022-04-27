package service

import "testing"

func TestQueryFilter_Validate(t *testing.T) {
	type fields struct {
		Type  string
		Sort  string
		Page  int
		Limit int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "with negative limit",
			fields: fields{
				Type:  "select",
				Sort:  "asc",
				Page:  1,
				Limit: -5,
			},
			wantErr: true,
		},
		{
			name: "with upper case type and sort",
			fields: fields{
				Type:  "SELECT",
				Sort:  "ASC",
				Page:  1,
				Limit: 10,
			},
			wantErr: false,
		}, {
			name: "with invalid type",
			fields: fields{
				Type:  "create",
				Sort:  "asc",
				Page:  1,
				Limit: 10,
			},
			wantErr: true,
		}, {
			name: "with invalid sort",
			fields: fields{
				Type:  "select",
				Sort:  "asc-desc",
				Page:  1,
				Limit: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := QueryFilter{
				Type:  tt.fields.Type,
				Sort:  tt.fields.Sort,
				Page:  tt.fields.Page,
				Limit: tt.fields.Limit,
			}
			if err := f.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
