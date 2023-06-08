package actors

import (
	"github.com/rmfachran/miniproject2/entity"
	"github.com/rmfachran/miniproject2/repository"
	"github.com/rmfachran/miniproject2/repository/mocks"
	"reflect"
	"testing"
	"time"
)

type fields struct {
	actorRepo mocks.ActorInterfaceRepo
}

func Test_useCaseActor_ApprovedAdmin(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	//newFields := mocks
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			got, err := uc.ApprovedAdmin(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApprovedAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ApprovedAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseActor_CreateAdmin(t *testing.T) {

	type args struct {
		actor ActorParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "success",
			fields: fields{*mocks.NewActorInterfaceRepo(t)},
			args: args{ActorParam{
				Username:   "john123",
				Password:   "test123",
				RoleId:     2,
				IsVerified: "false",
				IsActive:   "false",
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			}},
			want: entity.Actor{
				Username:   "john123",
				Password:   "$2a$10$gubPmqqGEe1ZWlZ9Td.fWO3lBwjSJGe.dXvfMHjphncWifxZkJ2Qa",
				RoleId:     2,
				IsVerified: "false",
				IsActive:   "false",
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: &tt.fields.actorRepo,
			}
			tt.fields.actorRepo.On("CreateAdmin", &tt.want).Return(&entity.Actor{}, nil)
			got, err := uc.CreateAdmin(tt.args.actor)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseActor_GetAdminById(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			fields:  fields{},
			args:    args{},
			want:    entity.Actor{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			got, err := uc.GetAdminById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAdminById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAdminById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseActor_GetCustomers(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	type args struct {
		first_name string
		last_name  string
		email      string
		page       int
		pageSize   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			got, err := uc.GetCustomers(tt.args.first_name, tt.args.last_name, tt.args.email, tt.args.page, tt.args.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseActor_LoginAdmin(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			got, err := uc.LoginAdmin(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseActor_LoginSuperAdmin(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			got, err := uc.LoginSuperAdmin(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginSuperAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginSuperAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_useCaseActor_SaveCustomersFromAPI(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			if err := uc.SaveCustomersFromAPI(); (err != nil) != tt.wantErr {
				t.Errorf("SaveCustomersFromAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCaseActor_UpdateAdmin(t *testing.T) {
	type fields struct {
		actorRepo repository.ActorInterfaceRepo
	}
	type args struct {
		id  uint
		adm ActorParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Actor
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: tt.fields.actorRepo,
			}
			got, err := uc.UpdateAdmin(tt.args.id, tt.args.adm)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateAdmin() got = %v, want %v", got, tt.want)
			}
		})
	}
}
