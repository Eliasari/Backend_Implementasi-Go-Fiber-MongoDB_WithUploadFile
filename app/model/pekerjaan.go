// package model

// import "time"

// type Pekerjaan struct {
// 	ID                 int        `json:"id"`
// 	AlumniID           int        `json:"alumni_id"`
// 	NamaPerusahaan     string     `json:"nama_perusahaan"`
// 	PosisiJabatan      string     `json:"posisi_jabatan"`
// 	BidangIndustri     string     `json:"bidang_industri"`
// 	LokasiKerja        string     `json:"lokasi_kerja"`
// 	GajiRange          string     `json:"gaji_range"`
// 	TanggalMulaiKerja  time.Time  `json:"tanggal_mulai_kerja"`
// 	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja"`
// 	StatusPekerjaan    string     `json:"status_pekerjaan"`
// 	DeskripsiPekerjaan string     `json:"deskripsi_pekerjaan"`
// 	CreatedAt          time.Time  `json:"created_at"`
// 	UpdatedAt          time.Time  `json:"updated_at"`
// }

// // ==========================================
// // TRASH STRUCT (data yang sudah dihapus)
// // ==========================================
// type PekerjaanTrash struct {
// 	ID                int        `json:"id"`
// 	AlumniID          int        `json:"alumni_id"`
// 	NamaPerusahaan    string     `json:"nama_perusahaan"`
// 	PosisiJabatan     string     `json:"posisi_jabatan"`
// 	BidangIndustri    string     `json:"bidang_industri"`
// 	LokasiKerja       string     `json:"lokasi_kerja"`
// 	GajiRange         string     `json:"gaji_range"`
// 	TanggalMulaiKerja string     `json:"tanggal_mulai_kerja"`
// 	TanggalSelesaiKerja *string  `json:"tanggal_selesai_kerja"`
// 	StatusPekerjaan   string     `json:"status_pekerjaan"`
// 	DeskripsiPekerjaan string    `json:"deskripsi_pekerjaan"`
// 	CreatedAt         time.Time  `json:"created_at"`
// 	UpdatedAt         time.Time  `json:"updated_at"`
// 	IsDelete          *time.Time `json:"is_delete"`
// }


package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pekerjaan struct {
	ID                  primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	AlumniID            primitive.ObjectID  `json:"alumni_id,omitempty" bson:"alumni_id,omitempty"`
	NamaPerusahaan      string              `json:"nama_perusahaan" bson:"nama_perusahaan"`
	PosisiJabatan       string              `json:"posisi_jabatan" bson:"posisi_jabatan"`
	BidangIndustri      string              `json:"bidang_industri" bson:"bidang_industri"`
	LokasiKerja         string              `json:"lokasi_kerja" bson:"lokasi_kerja"`
	GajiRange           string              `json:"gaji_range" bson:"gaji_range"`
	TanggalMulaiKerja   time.Time           `json:"tanggal_mulai_kerja" bson:"tanggal_mulai_kerja"`
	TanggalSelesaiKerja *time.Time          `json:"tanggal_selesai_kerja,omitempty" bson:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string              `json:"status_pekerjaan" bson:"status_pekerjaan"`
	DeskripsiPekerjaan  string              `json:"deskripsi_pekerjaan" bson:"deskripsi_pekerjaan"`
	CreatedAt           time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at" bson:"updated_at"`
}

type UpdatePekerjaan struct {
	NamaPerusahaan      string     `json:"nama_perusahaan,omitempty" bson:"nama_perusahaan,omitempty"`
	PosisiJabatan       string     `json:"posisi_jabatan,omitempty" bson:"posisi_jabatan,omitempty"`
	BidangIndustri      string     `json:"bidang_industri,omitempty" bson:"bidang_industri,omitempty"`
	LokasiKerja         string     `json:"lokasi_kerja,omitempty" bson:"lokasi_kerja,omitempty"`
	GajiRange           string     `json:"gaji_range,omitempty" bson:"gaji_range,omitempty"`
	TanggalMulaiKerja   *time.Time `json:"tanggal_mulai_kerja,omitempty" bson:"tanggal_mulai_kerja,omitempty"`
	TanggalSelesaiKerja *time.Time `json:"tanggal_selesai_kerja,omitempty" bson:"tanggal_selesai_kerja,omitempty"`
	StatusPekerjaan     string     `json:"status_pekerjaan,omitempty" bson:"status_pekerjaan,omitempty"`
	DeskripsiPekerjaan  string     `json:"deskripsi_pekerjaan,omitempty" bson:"deskripsi_pekerjaan,omitempty"`
	UpdatedAt           time.Time  `json:"updated_at" bson:"updated_at"`
}
