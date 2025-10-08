package repository

import (
    "database/sql"
    "go-fiber/app/model"
    "fmt"
    "log"
"time"
"errors" 
)

// func FindAllPekerjaan(db *sql.DB) ([]model.Pekerjaan, error) {
//     rows, err := db.Query(`
//         SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
//                lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
//                status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
//         FROM pekerjaan_alumni`)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var list []model.Pekerjaan
//     for rows.Next() {
//         var p model.Pekerjaan
//         if err := rows.Scan(
//             &p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
//             &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
//             &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
//         ); err != nil {
//             return nil, err
//         }
//         list = append(list, p)
//     }
//     return list, nil
// }

// func FindPekerjaanByID(db *sql.DB, id int) (*model.Pekerjaan, error) {
//     var p model.Pekerjaan
//     query := `SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
//                      lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
//                      status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
//               FROM pekerjaan_alumni WHERE id=$1 LIMIT 1`
//     err := db.QueryRow(query, id).
//         Scan(&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
//             &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
//             &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt)
//     if err != nil {
//         return nil, err
//     }
//     return &p, nil
// }

type PekerjaanRepository struct {
	DB *sql.DB
}

func CreatePekerjaan(db *sql.DB, p model.Pekerjaan) (*model.Pekerjaan, error) {
    query := `INSERT INTO pekerjaan_alumni 
              (alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
               tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan) 
              VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
              RETURNING id, created_at, updated_at`
    err := db.QueryRow(query,
        p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
        p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja,
        p.StatusPekerjaan, p.DeskripsiPekerjaan).
        Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &p, nil
}

func UpdatePekerjaan(db *sql.DB, id int, p model.Pekerjaan) (*model.Pekerjaan, error) {
    query := `UPDATE pekerjaan_alumni 
              SET alumni_id=$1, nama_perusahaan=$2, posisi_jabatan=$3, bidang_industri=$4,
                  lokasi_kerja=$5, gaji_range=$6, tanggal_mulai_kerja=$7, tanggal_selesai_kerja=$8,
                  status_pekerjaan=$9, deskripsi_pekerjaan=$10, updated_at=NOW()
              WHERE id=$11 RETURNING id, created_at, updated_at`
    err := db.QueryRow(query,
        p.AlumniID, p.NamaPerusahaan, p.PosisiJabatan, p.BidangIndustri,
        p.LokasiKerja, p.GajiRange, p.TanggalMulaiKerja, p.TanggalSelesaiKerja,
        p.StatusPekerjaan, p.DeskripsiPekerjaan, id).
        Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
    if err != nil {
        return nil, err
    }
    return &p, nil
}

func DeletePekerjaan(db *sql.DB, id int) error {
    _, err := db.Exec(`DELETE FROM pekerjaan_alumni WHERE id=$1`, id)
    return err
}

func FindPekerjaanByAlumniID(db *sql.DB, alumniID int) ([]model.Pekerjaan, error) {
    rows, err := db.Query(`
        SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
               lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
               status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at
        FROM pekerjaan_alumni
        WHERE alumni_id = $1`, alumniID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var list []model.Pekerjaan
    for rows.Next() {
        var p model.Pekerjaan
        if err := rows.Scan(
            &p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
            &p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
            &p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        list = append(list, p)
    }
    return list, nil
}

func GetPekerjaanRepo(db *sql.DB, search, sortBy, order string, limit, offset int) ([]model.Pekerjaan, error) {
	allowedSort := map[string]bool{"id": true, "nama_perusahaan": true, "posisi_jabatan": true, "tanggal_mulai_kerja": true}
	if !allowedSort[sortBy] {
		sortBy = "id"
	}
	if order != "asc" && order != "desc" {
		order = "asc"
	}

	query := fmt.Sprintf(`
		SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri, lokasi_kerja, gaji_range,
		       tanggal_mulai_kerja, tanggal_selesai_kerja, status_pekerjaan, deskripsi_pekerjaan,
		       created_at, updated_at
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, sortBy, order)

	log.Println("SQL:", query)
	log.Println("Params:", "%"+search+"%", limit, offset)

	rows, err := db.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Pekerjaan
	for rows.Next() {
		var p model.Pekerjaan
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		list = append(list, p)
	}

	return list, nil
}

func CountPekerjaanRepo(db *sql.DB, search string) (int, error) {
	var total int
	query := `
		SELECT COUNT(*) 
		FROM pekerjaan_alumni
		WHERE nama_perusahaan ILIKE $1 OR posisi_jabatan ILIKE $1 OR bidang_industri ILIKE $1 OR lokasi_kerja ILIKE $1`
	err := db.QueryRow(query, "%"+search+"%").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}


func NewPekerjaanRepository(db *sql.DB) *PekerjaanRepository {
    return &PekerjaanRepository{DB: db}
}

func (r *PekerjaanRepository) SoftDelete(id int, userID int, isAdmin bool) error {
    now := time.Now()

    if isAdmin {
        _, err := r.DB.Exec(`UPDATE pekerjaan_alumni SET is_delete = $1 WHERE id = $2`, now, id)
        return err
    }

    res, err := r.DB.Exec(`
        UPDATE pekerjaan_alumni p
        SET is_delete = $1
        FROM alumni a
        WHERE p.alumni_id = a.id AND p.id = $2 AND a.user_id = $3
    `, now, id, userID)
    if err != nil {
        return err
    }

    rows, _ := res.RowsAffected()
    if rows == 0 {
        return sql.ErrNoRows
    }
    return nil
}


func (r *PekerjaanRepository) GetTrashed(userID int, isAdmin bool) ([]model.Pekerjaan, error) {
	var rows *sql.Rows
	var err error

	if isAdmin {
		rows, err = r.DB.Query(`
			SELECT id, alumni_id, nama_perusahaan, posisi_jabatan, bidang_industri,
			       lokasi_kerja, gaji_range, tanggal_mulai_kerja, tanggal_selesai_kerja,
			       status_pekerjaan, deskripsi_pekerjaan, created_at, updated_at, is_delete
			FROM pekerjaan_alumni
			WHERE is_delete IS NOT NULL
		`)
	} else {
		rows, err = r.DB.Query(`
			SELECT p.id, p.alumni_id, p.nama_perusahaan, p.posisi_jabatan, p.bidang_industri,
			       p.lokasi_kerja, p.gaji_range, p.tanggal_mulai_kerja, p.tanggal_selesai_kerja,
			       p.status_pekerjaan, p.deskripsi_pekerjaan, p.created_at, p.updated_at, p.is_delete
			FROM pekerjaan_alumni p
			JOIN alumni a ON p.alumni_id = a.id
			WHERE p.is_delete IS NOT NULL AND a.user_id = $1
		`, userID)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Pekerjaan
	for rows.Next() {
		var p model.Pekerjaan
		if err := rows.Scan(
			&p.ID, &p.AlumniID, &p.NamaPerusahaan, &p.PosisiJabatan, &p.BidangIndustri,
			&p.LokasiKerja, &p.GajiRange, &p.TanggalMulaiKerja, &p.TanggalSelesaiKerja,
			&p.StatusPekerjaan, &p.DeskripsiPekerjaan, &p.CreatedAt, &p.UpdatedAt, &p.IsDelete,
		); err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func (r *PekerjaanRepository) Restore(id, userID int, isAdmin bool) error {
	query := `
		UPDATE pekerjaan_alumni pa
		SET is_delete = NULL
		WHERE pa.id = $1
		AND (
			EXISTS (
				SELECT 1 FROM alumni a
				WHERE a.id = pa.alumni_id
				AND a.user_id = $2
			)
			OR $3 = TRUE
		)
	`
	res, err := r.DB.Exec(query, id, userID, isAdmin)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("tidak diizinkan restore pekerjaan ini")
	}

	return nil
}


func (r *PekerjaanRepository) HardDelete(id, userID int, isAdmin bool) error {
	query := `
		DELETE FROM pekerjaan_alumni pa
		WHERE pa.id = $1
		AND (
			EXISTS (
				SELECT 1 FROM alumni a
				WHERE a.id = pa.alumni_id
				AND a.user_id = $2
			)
			OR $3 = TRUE
		)
	`

	res, err := r.DB.Exec(query, id, userID, isAdmin)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("tidak diizinkan hard delete pekerjaan ini")
	}

	return nil
}

