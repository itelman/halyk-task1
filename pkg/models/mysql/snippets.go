package mysql

import (
	"proxy/pkg/models"
)

type SnippetModel struct{}

func (m *SnippetModel) Get(id int) (*models.Request, error) {
	s := &models.Request{}

	/*err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}*/

	return s, nil
}
