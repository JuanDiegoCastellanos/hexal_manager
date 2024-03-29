package entities

import "hexal_manager/util/helpers"

type SocialSpace struct {
	ID                   int64              `json:"id"`
	Name                 string             `json:"name"`
	ResidentialComplexID int64              `json:"residential_complex_id"`
	Capacity             int64              `json:"capacity"`
	State                helpers.EnumEntity `json:"state"`
	SocialSpaceType      helpers.EnumEntity `json:"social_space_type"`
}

func NewSocialSpace(
	id, residentialComplexID, capacity int64, name string,
	state, socialSpaceType helpers.EnumEntity) *SocialSpace {
	return &SocialSpace{
		ID:                   id,
		Name:                 name,
		State:                state,
		SocialSpaceType:      socialSpaceType,
		Capacity:             capacity,
		ResidentialComplexID: residentialComplexID,
	}
}
