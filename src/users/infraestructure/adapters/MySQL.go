package adapters

import (
	"apiInvitation/src/core"
	"apiInvitation/src/users/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}
func NewMySQL() (*MySQL, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}, nil
}

func (mysql *MySQL) Register(user *entities.User) error {
	query := `
		INSERT INTO users (
			full_name, 
			email, 
			password_hash, 
			gender, 
			match_preference, 
			city, 
			state, 
			interests, 
			status_message, 
			profile_picture
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := mysql.conn.ExecutePreparedQuery(
		query,
		user.FullName,
		user.Email,
		user.PasswordHash,
		user.Gender,
		user.MatchPreference,
		user.City,
		user.State,
		user.Interests,
		user.StatusMessage,
		user.ProfilePicture,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
			lastInsertID, err := result.LastInsertId()
			if err != nil {
				fmt.Println(err)
				return err
			}
			user.Id = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}

func (mysql *MySQL) Update(id int32, fullname string, email string, passwordHash string,gender string, matchPreference string, city string, state string, interests string, statusMessage string) error {
	query := "UPDATE users SET name =?, lastname =?, password =?, role =? WHERE user_id =?"
    result, err := mysql.conn.ExecutePreparedQuery(query, fullname, email, passwordHash,gender, matchPreference, city, state,interests,statusMessage, id)
    if err!= nil {
        fmt.Println(err)
        return err
    }
    if result!= nil {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
        } else {
            log.Printf("[MySQL] - Ninguna fila fue afectada.")
        }
    } else {
        log.Printf("[MySQL] - Resultado de la consulta es nil.")
    }
    return nil
}
func (mysql *MySQL) GetAll() ([]*entities.User, error) {
	query := "SELECT * FROM users WHERE deleted = 0"
    rows, err := mysql.conn.FetchRows(query)
    if err!= nil {
        fmt.Println(err)
        return nil, err
    }
    defer rows.Close()
    var users []*entities.User
    var deleted bool
	var createdAt string
	var updatedAt string
    for rows.Next() {
        user := entities.User{}
        err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash,&user.Gender,&user.MatchPreference,&user.City,&user.State,&user.Interests,&user.StatusMessage, &user.ProfilePicture,&createdAt, &updatedAt ,&deleted)
		fmt.Println(user.PasswordHash)
        if err!= nil {
            fmt.Println(err)
            return nil, err
        }
        users = append(users, &user)
    }
    return users, nil
}
func (mysql *MySQL) Delete(id int32) error {
	query := "UPDATE users SET deleted = 1 WHERE user_id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("[MySQL] - Error al ejecutar la consulta: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("[MySQL] - Error al obtener las filas afectadas: %v", err)
		return err
	}
	if rowsAffected == 0 {
		log.Printf("[MySQL] - Ninguna fila fue afectada. Producto con ID %d no encontrado.", id)
		return fmt.Errorf("producto con ID %d no encontrado", id)
	}

	log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQL) GetById(id int32) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	rows, err := mysql.conn.FetchRows(query, id)
    if err!= nil {
        fmt.Println(err)
        return nil, err
    }
    defer rows.Close()
	var user entities.User
	var deleted bool
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash,&user.Gender,&user.MatchPreference,&user.City,&user.State,&user.Interests,&user.StatusMessage, &user.ProfilePicture, &deleted)
        if err!= nil {
            fmt.Println(err)
            return nil, err
        }
    }
	return &user, nil
}
func (mysql *MySQL) UploadPicture(id int32, urlPicture string) error {
	query := "UPDATE users SET profile_picture =? WHERE id =?"
    result, err := mysql.conn.ExecutePreparedQuery(query,urlPicture, id)
    if err!= nil {
        fmt.Println(err)
        return err
    }
    if result!= nil {
        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 1 {
            log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
        } else {
            log.Printf("[MySQL] - Ninguna fila fue afectada.")
        }
    } else {
        log.Printf("[MySQL] - Resultado de la consulta es nil.")
    }
    return nil
}