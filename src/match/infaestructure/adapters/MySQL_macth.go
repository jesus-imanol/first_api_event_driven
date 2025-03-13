package adapters

import (
	"apiInvitation/src/core"
	"apiInvitation/src/match/domain/entities"
	"apiInvitation/src/match/domain/models"
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
func (mysql *MySQL) Send(match *entities.Match) error {
	query := "INSERT INTO user_matches (sender_id, receiver_id, status) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, match.SenderUser, match.ReceiverUser, match.Status)
	if err!= nil {
		log.Printf("[MySQL] - Error al ejecutar la consulta: %v", err)
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
			match.Id = int32(lastInsertID)
		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}
func (mysql *MySQL) GetUserMatchesWithDetails(userId int32) ([]*models.MatchWithDetails, error) {
	query := `
		SELECT 
			m.match_id, 
			m.sender_id, 
			m.receiver_id, 
			m.status,
			u1.full_name as sender_name,
			u2.full_name as receiver_name,
			u1.profile_picture as sender_picture,
			u2.profile_picture as receiver_picture
		FROM 
			user_matches m
		JOIN 
			users u1 ON m.sender_id = u1.id
		JOIN 
			users u2 ON m.receiver_id = u2.id
		WHERE 
			m.sender_id = ? OR m.receiver_id = ?
		ORDER BY 
			m.updated_at DESC
	`
	
	rows, err := mysql.conn.FetchRows(query, userId, userId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	
	var matchesWithDetails []*models.MatchWithDetails
	for rows.Next() {
		match := models.MatchWithDetails{
			Match: entities.Match{},
		}
		
		err := rows.Scan(
			&match.Id, 
			&match.SenderUser, 
			&match.ReceiverUser, 
			&match.Status,
			&match.SenderName,
			&match.ReceiverName,
			&match.SenderPicture,
			&match.ReceiverPicture,
		)
		
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		
		matchesWithDetails = append(matchesWithDetails, &match)
	}
	
	return matchesWithDetails, nil
}