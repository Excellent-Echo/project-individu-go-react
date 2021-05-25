package sportlist

import "projectpenyewaanlapangan/entity"

func FormatSportlist(sportlist entity.SportList) entity.SportList {
	var formatSportlist = entity.SportList{
		ID:        sportlist.ID,
		SportName: sportlist.SportName,
	}

	return formatSportlist
}
