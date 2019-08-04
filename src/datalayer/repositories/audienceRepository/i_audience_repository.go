package audienceRepository

import "platform2.0-go-challenge/src/models"

type IAudienceRepository interface {
	GetAudiences(userID int) ([]models.Audience, error)
	GetAudiencesPaginated(userID, limit, offset int) ([]models.Audience, error)
	AddAudience(audience models.Audience) (int, error)
	EditAudience(audience models.Audience) (int64, error)
	DeleteAudience(id int) (int64, error)
}
