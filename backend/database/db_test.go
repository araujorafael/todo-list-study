package database

import (
	"reflect"
	"testing"
	"todo-list-study/backend/models"
)

func TestNewDatabase(t *testing.T) {
	tests := []struct {
		name string
		want *DatabaseImpl
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseImpl_AddTask(t *testing.T) {
	type fields struct {
		Data []models.Task
	}
	type args struct {
		task models.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []models.Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseImpl{
				Data: tt.fields.Data,
			}
			if got := d.AddTask(tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DatabaseImpl.AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseImpl_ListTasks(t *testing.T) {
	type fields struct {
		Data []models.Task
	}
	tests := []struct {
		name   string
		fields fields
		want   []models.Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseImpl{
				Data: tt.fields.Data,
			}
			if got := d.ListTasks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DatabaseImpl.ListTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseImpl_GetTask(t *testing.T) {
	type fields struct {
		Data []models.Task
	}
	type args struct {
		id uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Task
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DatabaseImpl{
				Data: tt.fields.Data,
			}
			got, err := d.GetTask(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DatabaseImpl.GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DatabaseImpl.GetTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
