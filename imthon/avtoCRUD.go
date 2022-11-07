package main 

import (
	"database/sql"
	"fmt"
	"time"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) DBManager {
	return DBManager{db}
}

type Avtomobil struct {
	ID int64
	CategoryName string
	Name string
	Price float64
	ImageUrl string
	CreatedAt time.Time
	Images []*ProductImage
}

type AvtomobilImage struct {
	ID int64
	ImageUrl string
	SequenceNumber int32
}

type GetAvtomobilParams struct {
	Limit int32
	Page int32
	Search string
}

type GetAvtomobilResponse struct {
	Avtomobil []*Avtomobil
	Count int32
}

func (A *DBManager) CreateAvtomobil(Avtomobil *Avtomobil) (int64, error) {
	var AvtomobilID int64
	query := `
		INSERT INTO AVTOMOBIL (
				category_name,
				name,
				price,
				image_url
			)   VALUES($1, $2, $3, $4)
			RETURNING id
	`

	row := A.db.QueryRow(
		query,
		Avtomobil.CategoryName,
		Avtomobil.Name,
		Avtomobil.Price,
		Avtomobil.ImageUrl,
	)

	err := row.Scan(&AvtomobilID)
	if err != nil {
		return 0, err
	}

	queryInsertImage := `
		INSERT INTO avtomobil_images (
				Avtomobil_name,
				image_url,
				sequence_number
			) VALUES($1, $2, $3) 
	`

	for _, image := range Avtomobil.Images {
		_, err := A.db.Exec(
			queryInsertImage,
			AvtomobilID.
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			return 0, err
		}
	}
	return AvtomobilID, nil
}

func (A *DBManager) GetAvtomobil(id, int64) (*Avtomobil, error) {
	var Avtomobil Avtomobil

	Avtomobil.Images = make([]*AvtomobilImage, 0)

	query := `
		SELECT
			a.id,
			a.category_name,
			c.name,
			p.name,
			p.price,
			p.image_url,
			p.created_at
		FROM avtomobil p
		INNER JOIN categories c on c.id=p.category_name
		WHERE p.id=$1
	`

	row := A.db.QueryRow(query, id)

	err := row.Scan(
		&Avtomobil.ID,
		&Avtomobil.CategoryName,
		&Avtomobil.Name,
		&Avtomobil.Price,
		&Avtomobil.ImageUrl,
		&Avtomobil.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	queryImages := `
		SELECT
			id,
			image_url,
			sequence_number
		FROM avtomobil_images
		WHERE avtomobil_name=$1
	`

	rows, err := m.db.Query(queryImages, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var image AvtomobilImage

		err := rows.Scan(
			&image.ID,
			&image.ImageUrl,
			&image.SequenceNumber,
		)
		if err != nil {
			return nil, err
		}
		Avtomobil.image = append(Avtomobil.Images, *&image)
	}
	return &Avtomobil, nil
}

func (A *DBManager) GetAllAvtomobils(params *GetAvtomobilParams)
(*GetAvtomobilResponse , error) {
	var result GetAvtomobilResponse

	result.Avtomobil = make([]*Avtomobil, 0)

	filter := ""
	if params.Search != "" {
		filter = fmt.Sprintf("WHERE name ilike '%s'", "%"+params.Search+"%")
	}

	query := `
		SELECT
			p.id,
			p.category_name,
			c.name,
			p.name,
			p.image_url,
			p.created_at
		FROM avtomobil p
		INNER JOIN categories c on c.id=p.category_name
		` + filter + `
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
		`
	
	offset := (params.Page - 1) * params.Limit
	rows, err := A.db.Query(query, params.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var Avtomobil Avtomobil

		err := rows.Scan(
			&avtomobil.ID,
			&avtomobil.CategoryName,
			&avtomobil.Name,
			&avtomobil.Price,
			&avtomobil.ImageUrl,
			&avtomobil.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Avtomobil = append(result.Avtomobil, &avtomobil)
	}
	return &result, nil
}

func (a *DBManager) UpdateAvtomobil(avtomobil *Avtomobil) error {
	query := `
		UPDATE avtomobils SET
			category_id=$1,
			name=$2,
			price=$3,
			image_url=$4
		WHERE name=$5
	`

	result, err := a.db.Exec(
		query,
		avtomobil.CategoryName,
		avtomobil_name,
		avtombil.price,
		avtomobil.ImageUrl,
		avtomobil.ID,
	)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.errNoRows
	}
	queryDeleteImages := `DELETE FROM avtomobil_images WHERE avtomobil_id=$1`
	_, err = a.db.Exec(queryDeleteImages, avtombil.ID)
	if err != nil {
		return err
	}

	queryInsertImage := `
		INSERT INTO avtomobil_images (
			avtomobil_id,
			image_url,
			sequence_number
		) VALUES($1, $2, $3)
	`

	for _, image := range avtomobil.Images {
		_, err := a.db.Exec(
			queryInsertImage,
			avtombil.ID,
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *DBManager) DeleteAvtomobil(id int64) error {
	queryDeleteImages := `DELETE FROM avtomobil_images WHERE avtomobil_id=$1`
	_, err := a.db.Exec(queryDeleteImages, id)
	if err != nil {
		return err
	}

	queryDelete := `DELETE FROM avtomobil WHERE id=$1`
	result, err := a.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}
	if rowsCount == 0 {
		return sql.errNoRows
	}

	return nil
}