package ivi

import (
	"strings"
)

func (o Compilation) Line() []string {
	persons := getJSON(o.Persons)
	genres := "{" + strings.Join(o.Genres, ",") + "}"
	awards := getJSON(o.Awards)
	posters := getJSON(o.Posters)
	seasondescription := getJSON(o.SeasonDescription)
	releasedates := getJSON(o.ReleaseDates)
	additionaldata := getJSON(o.AdditionalData)
	categories := "{" + strings.Join(o.Categories, ",") + "}"
	tima := getJSON(o.Tima)
	productioncompany := getJSON(o.ProductionCompany)
	theme := getJSON(o.Theme)
	audience := getJSON(o.Audience)
	mood := getJSON(o.Mood)
	quality := getJSON(o.Quality)
	about := getJSON(o.About)
	basis := getJSON(o.Basis)
	place := getJSON(o.Place)
	subgenre := getJSON(o.Subgenre)
	gender := getJSON(o.Gender)
	character := getJSON(o.Character)
	categoryproperty := getJSON(o.CategoryProperty)
	check := getJSON(o.Check)
	restrictions := getJSON(o.Restrictions)
	return []string{o.ID, o.Title, o.Description, o.Restrict, o.OrigCountry, o.KinopoiskID, o.ImdbRating, o.KinopoiskRating, o.UsaBoxOffice, o.WorldBoxOffice, o.WeekendBoxOffice, o.ProductionBudget, o.Hru, o.FirstSeriesDate, o.LastSeriesDate, persons, genres, awards, posters, seasondescription, releasedates, additionaldata, categories, tima, productioncompany, theme, audience, mood, quality, about, basis, place, o.Fake, subgenre, gender, character, categoryproperty, check, restrictions}
}

func (o ContentCategory) Line() []string {
	return []string{o.ID, o.Title, o.Description, o.Hru}
}

func (o Genre) Line() []string {
	return []string{o.ID, o.CategoryID, o.Title, o.Description, o.Hru}
}

func (o Person) Line() []string {
	posters := getJSON(o.Posters)
	return []string{o.ID, posters, o.Name, o.IsStar, o.EngTitle, o.GenTitle, o.Hru, o.FakeHru, o.BlogPost, o.BlogLink, o.BlogAdded, o.Bio, o.Facts, o.Rss, o.KinopoiskID}
}

func (o Property) Line() []string {
	return []string{o.ID, o.Title, o.XMLField, o.Unique}
}

func (o Content) Line() []string {
	contentgenres := getJSON(o.ContentGenres)
	persons := getJSON(o.Persons)
	awards := getJSON(o.Awards)
	moods := getJSON(o.Moods)
	contentcategories := getJSON(o.ContentCategories)
	posters := getJSON(o.Posters)
	rusversions := getJSON(o.RusVersions)
	additionaldata := getJSON(o.AdditionalData)
	thumbs := getJSON(o.Thumbs)
	subscriptions := getJSON(o.Subscriptions)
	righttypelist := getJSON(o.RightTypeList)
	rightholders := getJSON(o.RightHolders)
	tima := getJSON(o.Tima)
	theme := getJSON(o.Theme)
	audience := getJSON(o.Audience)
	subgenre := getJSON(o.Subgenre)
	mood := getJSON(o.Mood)
	quality := getJSON(o.Quality)
	about := getJSON(o.About)
	basis := getJSON(o.Basis)
	categoryproperty := getJSON(o.CategoryProperty)
	check := getJSON(o.Check)
	place := getJSON(o.Place)
	restrictions := getJSON(o.Restrictions)
	testproperty := getJSON(o.TestProperty)
	localizationcheck := getJSON(o.LocalizationCheck)
	gender := getJSON(o.Gender)
	productioncompany := getJSON(o.ProductionCompany)
	character := getJSON(o.Character)
	franchise := getJSON(o.Franchise)
	shild := getJSON(o.Shild)
	newcontent := getJSON(o.NewContent)
	return []string{o.ID, o.Title, o.OrigTitle, o.Compilation, o.CopyrightOverlay, o.Season, o.Episode, o.Duration, o.PreviewID, o.Restrict, o.PrSiteID, o.OriginCountry, o.ReleaseDate, o.KinopoiskID, o.CompilationID, o.PromoID, o.ImdbRating, o.KinopoiskRating, o.UsaBoxOffice, o.WorldBoxOffice, o.WeekendBoxOffice, o.ProductionBudget, o.YoutubeTrailerURL, o.DateInsert, o.DescriptionAlt, o.SharingDisabled, o.RightDateEnd, o.DrmOnly, o.Instruments, contentgenres, persons, awards, moods, contentcategories, posters, o.Trailers, rusversions, additionaldata, thumbs, subscriptions, righttypelist, o.RightType, rightholders, o.Mp4, tima, theme, audience, subgenre, mood, quality, about, basis, categoryproperty, check, place, restrictions, testproperty, localizationcheck, gender, productioncompany, character, franchise, shild, newcontent}
}

func (o Award) Line() []string {
	return []string{o.ID, o.Title, o.Description, o.Place, o.URLToResource, o.ImageID, o.StartYear, o.EndYear}
}

func (o Country) Line() []string {
	return []string{o.ID, o.Title, o.Name}
}

func (o Lang) Line() []string {
	return []string{o.ID, o.Title, o.Name}
}

func (o ContentType) Line() []string {
	return []string{o.ID, o.Title, o.Description}
}

func (o PropertyValue) Line() []string {
	return []string{o.ID, o.Title}
}

func (o PromoInfo) Line() []string {
	return []string{o.ID, o.Title}
}

func (o PersonType) Line() []string {
	return []string{o.ID, o.Title, o.Description}
}

func (o RusVersion) Line() []string {
	return []string{o.ID, o.Title}
}

func (o ContentFormat) Line() []string {
	return []string{o.ID, o.Title, o.Description, o.TypeID, o.Group}
}

func (o Instrument) Line() []string {
	return []string{o.ID, o.Title, o.Info}
}

func (o Mood) Line() []string {
	return []string{o.ID, o.Title, o.Info}
}

func (o ContentAdditionalDataType) Line() []string {
	return []string{o.ID, o.Title}
}
