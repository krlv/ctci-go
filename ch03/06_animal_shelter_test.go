package ch03

import (
	"reflect"
	"testing"
)

func TestShelter_DequeueAny(t *testing.T) {
	shelter := NewShelter()
	shelter.Enqueue(Cat{animal{id: 1, name: "Garfield"}})
	shelter.Enqueue(Dog{animal{id: 2, name: "Belka"}})

	dogShelter := NewShelter()
	dogShelter.Enqueue(Dog{animal{id: 1, name: "Belka"}})
	dogShelter.Enqueue(Dog{animal{id: 2, name: "Strelka"}})

	catShelter := NewShelter()
	catShelter.Enqueue(Cat{animal{id: 1, name: "Garfield"}})
	catShelter.Enqueue(Cat{animal{id: 2, name: "Grumpy"}})

	tests := []struct {
		name    string
		shelter *Shelter
		want    Animal
		wantErr bool
	}{
		{
			name:    "returns the oldest animal from the shelter",
			shelter: shelter,
			want:    Cat{animal{id: 1, name: "Garfield"}},
		},
		{
			name:    "returns the oldest dog from the shelter",
			shelter: dogShelter,
			want:    Dog{animal{id: 1, name: "Belka"}},
		},
		{
			name:    "returns the oldest cat from the shelter",
			shelter: catShelter,
			want:    Cat{animal{id: 1, name: "Garfield"}},
		},
		{
			name:    "returns error if the shelter is empty",
			shelter: NewShelter(),
			want:    Dog{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shelter.DequeueAny()
			if (err != nil) != tt.wantErr {
				t.Errorf("DequeueAny() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DequeueAny() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShelter_DequeueCat(t *testing.T) {
	shelter := NewShelter()
	shelter.Enqueue(Cat{animal{id: 1, name: "Garfield"}})
	shelter.Enqueue(Dog{animal{id: 2, name: "Belka"}})

	dogShelter := NewShelter()
	dogShelter.Enqueue(Dog{animal{id: 1, name: "Belka"}})
	dogShelter.Enqueue(Dog{animal{id: 2, name: "Strelka"}})

	tests := []struct {
		name    string
		shelter *Shelter
		want    Cat
		wantErr bool
	}{
		{
			name:    "returns the oldest cat from the shelter",
			shelter: shelter,
			want:    Cat{animal{id: 1, name: "Garfield"}},
		},
		{
			name:    "returns error if there are no cats in the shelter",
			shelter: dogShelter,
			wantErr: true,
		},
		{
			name:    "returns error if the shelter is empty",
			shelter: NewShelter(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shelter.DequeueCat()
			if (err != nil) != tt.wantErr {
				t.Errorf("DequeueCat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DequeueCat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShelter_DequeueDog(t *testing.T) {
	shelter := NewShelter()
	shelter.Enqueue(Cat{animal{id: 1, name: "Garfield"}})
	shelter.Enqueue(Dog{animal{id: 2, name: "Belka"}})

	catShelter := NewShelter()
	catShelter.Enqueue(Cat{animal{id: 1, name: "Garfield"}})
	catShelter.Enqueue(Cat{animal{id: 2, name: "Grumpy"}})

	tests := []struct {
		name    string
		shelter *Shelter
		want    Dog
		wantErr bool
	}{
		{
			name:    "returns the oldest dog from the shelter",
			shelter: shelter,
			want:    Dog{animal{id: 2, name: "Belka"}},
		},
		{
			name:    "returns error if there are no dogs in the shelter",
			shelter: catShelter,
			wantErr: true,
		},
		{
			name:    "returns error if the shelter is empty",
			shelter: NewShelter(),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shelter.DequeueDog()
			if (err != nil) != tt.wantErr {
				t.Errorf("DequeueDog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DequeueDog() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_animal_IsOlder(t *testing.T) {
	type fields struct {
		id   int
		name string
	}
	type args struct {
		o Animal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "returns true if the animal is older",
			fields: fields{id: 1, name: "Garfield"},
			args:   args{o: animal{id: 2}},
			want:   true,
		},
		{
			name:   "returns false if the animal is younger",
			fields: fields{id: 2, name: "Grumpy"},
			args:   args{o: animal{id: 1}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := animal{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := a.IsOlder(tt.args.o); got != tt.want {
				t.Errorf("IsOlder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_animal_Name(t *testing.T) {
	type fields struct {
		id   int
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns the name of the animal",
			fields: fields{name: "Garfield"},
			want:   "Garfield",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := animal{
				id:   tt.fields.id,
				name: tt.fields.name,
			}
			if got := a.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
