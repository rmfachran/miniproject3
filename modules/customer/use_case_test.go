package customer

import (
	"github.com/rmfachran/miniproject2/entity"
	"github.com/rmfachran/miniproject2/repository/mocks"
	"reflect"
	"testing"
)

type fields struct {
	customerRepo mocks.CustomerInterfaceRepo
}

func Test_useCaseCustomer_CreateCustomer(t *testing.T) {
	type args struct {
		customer CustomerParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		{
			// TODO: Add test cases.
			name:   "Success",
			fields: fields{*mocks.NewCustomerInterfaceRepo(t)},
			args: args{customer: CustomerParam{
				Email:     "akshdkasd",
				FirstName: "adasdasd",
				LastName:  "asd",
				Avatar:    "asdasds",
			}},
			want: entity.Customer{
				Email:     "akshdkasd",
				FirstName: "adasdasd",
				LastName:  "asd",
				Avatar:    "asdasds",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: &tt.fields.customerRepo,
			}
			tt.fields.customerRepo.On("CreateCustomer", &tt.want).Return(&tt.want, nil)
			got, err := uc.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCustomer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseCustomer_DeleteCustomerById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{*mocks.NewCustomerInterfaceRepo(t)},
			args:    args{2},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: &tt.fields.customerRepo,
			}
			tt.fields.customerRepo.On("GetCustomerById", tt.args.id).Return(&entity.Customer{}, nil)
			tt.fields.customerRepo.On("DeleteById", tt.args.id, &entity.Customer{}).Return(nil)
			if err := uc.DeleteCustomerById(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCustomerById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCaseCustomer_GetCustomerById(t *testing.T) {

	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		{
			// TODO: Add test cases.
			name:   "",
			fields: fields{*mocks.NewCustomerInterfaceRepo(t)},
			args:   args{2},
			want: entity.Customer{
				ID:        2,
				Email:     "john@example.com",
				FirstName: "john",
				LastName:  "tor",
				Avatar:    "asdasdasd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: &tt.fields.customerRepo,
			}
			tt.fields.customerRepo.On("GetCustomerById", tt.args.id).Return(&tt.want, nil)
			got, err := uc.GetCustomerById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseCustomer_UpdateCustomerById(t *testing.T) {

	type args struct {
		id   uint
		cust CustomerParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "success",
			fields: fields{*mocks.NewCustomerInterfaceRepo(t)},
			args: args{
				id: 0,
				cust: CustomerParam{
					ID:        0,
					Email:     "john@example.com",
					FirstName: "john",
					LastName:  "tor",
					Avatar:    "asdasdasd",
				},
			},
			want: entity.Customer{
				ID:        0,
				Email:     "john@example.com",
				FirstName: "john",
				LastName:  "tor",
				Avatar:    "asdasdasd",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseCustomer{
				customerRepo: &tt.fields.customerRepo,
			}
			tt.fields.customerRepo.On("GetCustomerById", tt.args.id).Return(&tt.want, nil)
			tt.fields.customerRepo.On("UpdateCustomerById", tt.args.id, &tt.want).Return(&tt.want, nil)
			got, err := uc.UpdateCustomerById(tt.args.id, tt.args.cust)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCustomerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
