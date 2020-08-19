package main

import (
    "database/sql"
    "fmt"
)

type province struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
}


func (p *province) getProvince(db *sql.DB) error {
    statement := fmt.Sprintf("SELECT provinciaid AS id, provincia AS name FROM provincia WHERE provinciaid=%d", p.ID)
    return db.QueryRow(statement).Scan(&p.ID,&p.Name)
}

func getProvinces(db *sql.DB, start, count int) ([]province, error) {
    statement := fmt.Sprintf("SELECT provinciaid AS id, provincia AS name FROM provincia LIMIT %d OFFSET %d", count, start)
    rows, err := db.Query(statement)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var provinces []province

    for rows.Next() {
        var p province
        if err := rows.Scan(&p.ID, &p.Name); err != nil {
            return nil, err
        }
        provinces = append(provinces, p)
    }

    return provinces, nil
}
