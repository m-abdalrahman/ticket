package models

import "ticket/internal/util/struc"

func (m *Model) NewComment(comment *struc.Comment) error {
	return m.db.Create(comment).Error
}
