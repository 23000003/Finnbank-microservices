package services

// Use this for resolvers business logic

import (
	"context"
	"finnbank/common/utils"
	t "finnbank/graphql-api/types"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NotificationService struct {
	db *pgxpool.Pool
	l  *utils.Logger
}

func NewNotificationService(db *pgxpool.Pool, logger *utils.Logger) *NotificationService {
	return &NotificationService{
		db: db,
		l:  logger,
	}
}


// GetAllNotificationByUserId, (Query)
func (s *NotificationService) GetAllNotificationByUserId(notifToID string, limit int) ([]t.Notification, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT notif_id, notif_type, notif_to_id, notif_from_name,
		       content, is_read, redirect_url, date_notified, date_read
		FROM notifications
		WHERE notif_to_id = $1
		ORDER BY date_notified DESC
		LIMIT $2
	`, notifToID, limit)
	if err != nil {
		s.l.Error("DB query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var notifications []t.Notification
	for rows.Next() {
		var notif t.Notification
		err := rows.Scan(
			&notif.NotifID, &notif.NotifType, &notif.NotifToID,
			&notif.NotifFromName, &notif.Content, &notif.IsRead,
			&notif.RedirectURL, &notif.DateNotified, &notif.DateRead,
		)
		if err != nil {
			s.l.Error("Scan failed: %v", err)
			continue
		}
		notifications = append(notifications, notif)
	}
	return notifications, nil
}

func (s *NotificationService) GetAllUnreadNotificationByUserId(notifToID string) (*t.UnreadAndTotalNotification, error) {
	var unreadAndTotal t.UnreadAndTotalNotification
	err := s.db.QueryRow(context.Background(), `
		SELECT COUNT(*)
		FROM notifications
		WHERE notif_to_id = $1 AND is_read = FALSE
	`, notifToID).Scan(&unreadAndTotal.UnreadNotification)

	if err != nil {
		s.l.Error("Unread Notif DB query failed: %v", err)
		return nil, err
	}

	err1 := s.db.QueryRow(context.Background(), `
		SELECT COUNT(*)
		FROM notifications
		WHERE notif_to_id = $1
	`, notifToID).Scan(&unreadAndTotal.TotalNotification)

	if err1 != nil {
		s.l.Error("Total Notif DB query failed: %v", err1)
		return nil, err1
	}


	return &unreadAndTotal, nil
}

// GetNotificationByUserId, (Query)
func (s *NotificationService) GetNotificationByUserId(notifID string) (*t.Notification, error) {
	query := `
		SELECT notif_id, notif_type, notif_to_id, notif_from_name,
		       content, is_read, redirect_url, date_notified, date_read
		FROM notifications
		WHERE notif_id = $1
	`

	var notif t.Notification
	err := s.db.QueryRow(context.Background(), query, notifID).Scan(
		&notif.NotifID, &notif.NotifType, &notif.NotifToID,
		&notif.NotifFromName, &notif.Content, &notif.IsRead,
		&notif.RedirectURL, &notif.DateNotified, &notif.DateRead,
	)

	if err != nil {
		s.l.Error("GetNotificationByUserId failed: %v", err)
		return nil, err
	}

	return &notif, nil
}

// GenerateNotification, (Mutation)
func (s *NotificationService) GenerateNotification(notif t.Notification, notifConn *websocket.Conn) (*t.Notification, error) {
	notifID := uuid.New().String()  // Generate UUID
	notif.NotifID = notifID         // Assign it to the model
	notif.DateNotified = time.Now() // Make sure this is set before insert

	query := `
		INSERT INTO notifications (
			notif_id, notif_type, notif_to_id, notif_from_name,
			content, is_read, redirect_url, date_notified
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING notif_id, date_notified
	`

	var returnedID string
	var dateNotified time.Time

	err := s.db.QueryRow(context.Background(), query,
		notifID,
		notif.NotifType,
		notif.NotifToID,
		notif.NotifFromName,
		notif.Content,
		notif.IsRead,
		notif.RedirectURL,
		notif.DateNotified,
	).Scan(&returnedID, &dateNotified)

	if err != nil {
		s.l.Error("GenerateNotification query failed: %v", err)
		return nil, err
	}

	notif.NotifID = returnedID
	notif.DateNotified = dateNotified


	sendNotif := t.Notification{
		NotifID:       notif.NotifID,
		NotifType:     notif.NotifType,
		NotifToID:     notif.NotifToID,
		NotifFromName: notif.NotifFromName,
		Content:       notif.Content,
		IsRead:        notif.IsRead,
		RedirectURL:   notif.RedirectURL,
		DateNotified:  notif.DateNotified,
	}

	// Send the notification to the WebSocket connection
	if err := notifConn.WriteJSON(sendNotif); err != nil {
		s.l.Error("Error sending notification: %v", err)
		return nil, err
	}

	return &notif, nil
}

// ReadNotificationByUserId (Mutation)
func (s *NotificationService) ReadNotificationByUserId(notifID string) error {
	now := time.Now()
	query := `
		UPDATE notifications
		SET is_read = TRUE, date_read = $1
		WHERE notif_id = $2
	`

	_, err := s.db.Exec(context.Background(), query, now, notifID)
	if err != nil {
		s.l.Error("ReadNotificationByUserId failed: %v", err)
		return err
	}

	return nil
}

