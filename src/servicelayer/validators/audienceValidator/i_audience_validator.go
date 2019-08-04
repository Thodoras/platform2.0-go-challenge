package audienceValidator

import "platform2.0-go-challenge/src/models"

type IAudienceValidator interface {
	ValidateAudience(audience models.Audience) error
}
