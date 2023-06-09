package actors

import (
	"github.com/rmfachran/miniproject2/entity"
	"github.com/rmfachran/miniproject2/repository/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"time"
)

type fields struct {
	actorRepo mocks.ActorInterfaceRepo
}

// actor use case test
func Test_useCaseActor_ApprovedAdmin(t *testing.T) {

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
		{
			name:    "success",
			fields:  fields{*mocks.NewActorInterfaceRepo(t)},
			args:    args{2},
			want:    []*entity.Actor{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: &tt.fields.actorRepo,
			}
			tt.fields.actorRepo.On("ApprovedAdmin", mock.Anything).Return([]*entity.Actor{}, nil)
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
				Password:   "test123",
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
			//pass,_:=middleware.HashPassword(tt.args.actor.Password)
			tt.fields.actorRepo.On("CreateAdmin", mock.Anything).Return(&tt.want, nil)
			got, err := uc.CreateAdmin(tt.args.actor)
			got.Password = "test123"
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
			name:   "success",
			fields: fields{*mocks.NewActorInterfaceRepo(t)},
			args:   args{2},
			want: entity.Actor{
				Username:   "john",
				Password:   "john123",
				RoleId:     2,
				IsVerified: "true",
				IsActive:   "true",
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
			tt.fields.actorRepo.On("GetAdmin", tt.args.id).Return(&tt.want, nil)
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
		{
			name:   "success",
			fields: fields{*mocks.NewActorInterfaceRepo(t)},
			args: args{
				first_name: "john",
				last_name:  "tor",
				email:      "john@example.com",
				page:       1,
				pageSize:   2,
			},
			want:    []*entity.Customer{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: &tt.fields.actorRepo,
			}
			tt.fields.actorRepo.On("GetCustomers",
				tt.args.first_name,
				tt.args.last_name,
				tt.args.email,
				tt.args.page,
				tt.args.pageSize).Return([]*entity.Customer{}, nil)
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
		{
			name:   "success",
			fields: fields{*mocks.NewActorInterfaceRepo(t)},
			args: args{
				username: "john123",
				password: "test123",
			},
			want: &entity.Actor{
				Username:   "john123",
				Password:   "test123",
				RoleId:     2,
				IsVerified: "true",
				IsActive:   "true",
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
			tt.fields.actorRepo.On("LoginAdmin", mock.Anything).Return(tt.want, nil)
			got, err := uc.LoginAdmin(tt.args.username, tt.args.password)
			//got.Password = "test123"
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
		{
			name:   "success",
			fields: fields{*mocks.NewActorInterfaceRepo(t)},
			args: args{
				username: "john123",
				password: "test123",
			},
			want: &entity.Actor{
				Username:   "john123",
				Password:   "test123",
				RoleId:     2,
				IsVerified: "true",
				IsActive:   "true",
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
			tt.fields.actorRepo.On("LoginSuperAdmin", mock.Anything).Return(tt.want, nil)
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

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			fields:  fields{*mocks.NewActorInterfaceRepo(t)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := useCaseActor{
				actorRepo: &tt.fields.actorRepo,
			}
			tt.fields.actorRepo.On("SaveCustomersFromAPI", "https://reqres.in/api/users?page=2").Return(nil)
			if err := uc.SaveCustomersFromAPI(); (err != nil) != tt.wantErr {
				t.Errorf("SaveCustomersFromAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_useCaseActor_UpdateAdmin(t *testing.T) {

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
		{
			name:   "success",
			fields: fields{*mocks.NewActorInterfaceRepo(t)},
			args: args{
				id: 2,
				adm: ActorParam{
					ID:         2,
					Username:   "admin",
					Password:   "admin123",
					RoleId:     2,
					IsVerified: "true",
					IsActive:   "true",
					CreatedAt:  time.Time{},
					UpdatedAt:  time.Time{},
				},
			},
			want: entity.Actor{
				ID:         2,
				Username:   "admin",
				Password:   "admin123",
				RoleId:     2,
				IsVerified: "true",
				IsActive:   "true",
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
			tt.fields.actorRepo.On("GetAdmin", mock.Anything).Return(&tt.want, nil)
			tt.fields.actorRepo.On("UpdateAdmin", mock.Anything, mock.Anything).Return(&tt.want, nil)
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
