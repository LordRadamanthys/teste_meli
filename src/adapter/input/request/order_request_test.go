package request

import (
	"errors"
	"testing"
)

func TestValidateRequest(t *testing.T) {
	tests := []struct {
		name                  string
		order                 OrderRequest
		expectedErrorResponse error
	}{
		{
			name:                  "Validate itens exceed 100 ",
			order:                 mockRequest(),
			expectedErrorResponse: errors.New("the number of items cannot exceed 100"),
		},
		{
			name:                  "Itens cannot be empty",
			order:                 OrderRequest{[]ItemRequest{}},
			expectedErrorResponse: errors.New("items cannot be empty"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := tt.order
			v := resp.ValidateRequest()
			if v.Error() != tt.expectedErrorResponse.Error() {
				t.Errorf("Error() = %v, want %v", v, tt.expectedErrorResponse)
			}
		})
	}
}

func mockRequest() OrderRequest {
	return OrderRequest{
		[]ItemRequest{
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
			{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {},
		},
	}
}
