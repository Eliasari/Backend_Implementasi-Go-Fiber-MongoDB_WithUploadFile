// package model

// import "time"

// type Alumni struct {
// 	ID         int       `json:"id"`
// 	NIM        string    `json:"nim"`
// 	Nama       string    `json:"nama"`
// 	Jurusan    string    `json:"jurusan"`
// 	Angkatan   int       `json:"angkatan"`
// 	TahunLulus int       `json:"tahun_lulus"`
// 	Email      string    `json:"email"`
// 	NoTelepon  string    `json:"no_telepon"`
// 	Alamat     string    `json:"alamat"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// 	UserID     int       `json:"user_id"`
// }

// type GetAlumniRepo struct {
// 	ID         int       `json:"id"`
// 	NIM        string    `json:"nim"`
// 	Nama       string    `json:"nama"`
// 	Jurusan    string    `json:"jurusan"`
// 	Angkatan   int       `json:"angkatan"`
// 	TahunLulus int       `json:"tahun_lulus"`
// 	Email      string    `json:"email"`
// 	NoTelepon  string    `json:"no_telepon"`
// 	Alamat     string    `json:"alamat"`
// 	CreatedAt  time.Time `json:"created_at"`
// 	UpdatedAt  time.Time `json:"updated_at"`
// }

package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Alumni struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NIM        string             `bson:"nim" json:"nim"`
	Nama       string             `bson:"nama" json:"nama"`
	Jurusan    string             `bson:"jurusan" json:"jurusan"`
	Angkatan   int                `bson:"angkatan" json:"angkatan"`
	TahunLulus int                `bson:"tahun_lulus" json:"tahun_lulus"`
	Email      string             `bson:"email" json:"email"`
	NoTelepon  string             `bson:"no_telepon" json:"no_telepon"`
	Alamat     string             `bson:"alamat" json:"alamat"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
	UserID     primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type UpdateAlumni struct {
	NIM        string    `json:"nim" bson:"nim,omitempty"`
	Nama       string    `json:"nama" bson:"nama,omitempty"`
	Jurusan    string    `json:"jurusan" bson:"jurusan,omitempty"`
	Angkatan   int       `json:"angkatan" bson:"angkatan,omitempty"`
	TahunLulus int       `json:"tahun_lulus" bson:"tahun_lulus,omitempty"`
	Email      string    `json:"email" bson:"email,omitempty"`
	NoTelepon  string    `json:"no_telepon" bson:"no_telepon,omitempty"`
	Alamat     string    `json:"alamat" bson:"alamat,omitempty"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at,omitempty"`
}

type GetAlumni struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	NIM        string             `bson:"nim" json:"nim"`
	Nama       string             `bson:"nama" json:"nama"`
	Jurusan    string             `bson:"jurusan" json:"jurusan"`
	Angkatan   int                `bson:"angkatan" json:"angkatan"`
	TahunLulus int                `bson:"tahun_lulus" json:"tahun_lulus"`
	Email      string             `bson:"email" json:"email"`
	NoTelepon  string             `bson:"no_telepon" json:"no_telepon"`
	Alamat     string             `bson:"alamat" json:"alamat"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type AlumniResponse struct {
	Data []Alumni `json:"data"`
	Meta MetaInfo `json:"meta"`
}

