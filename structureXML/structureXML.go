package structureXML

import "encoding/xml"

type ContentCategory struct {
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Hru         string `xml:"hru"`
}

type ContentCategories struct {
	ContentCategory []ContentCategory `xml:"content_category"`
}

type Genre struct {
	ID          string `xml:"id"`
	CategoryID  string `xml:"category_id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Hru         string `xml:"hru"`
}

type Genres struct {
	ContentGenre []Genre `xml:"content_genre"`
}

type File struct {
	URL    string `xml:"url"`
	ID     string `xml:"id"`
	Format string `xml:"format"`
}

type Posters struct {
	File []File `xml:"file"`
}

type Person struct {
	ID          string  `xml:"id"`
	Posters     Posters `xml:"posters" pg:"jsonb"`
	Name        string  `xml:"name"`
	IsStar      string  `xml:"is_star"`
	EngTitle    string  `xml:"eng_title"`
	GenTitle    string  `xml:"gen_title"`
	Hru         string  `xml:"hru"`
	FakeHru     string  `xml:"fake_hru"`
	BlogPost    string  `xml:"blog_post"`
	BlogLink    string  `xml:"blog_link"`
	BlogAdded   string  `xml:"blog_added"`
	Bio         string  `xml:"bio"`
	Facts       string  `xml:"facts"`
	Rss         string  `xml:"rss"`
	KinopoiskID string  `xml:"kinopoisk_id"`
}

type Persons struct {
	Person []Person `xml:"person"`
}

type PersonType struct {
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type PersonsType struct {
	PersonType []PersonType `xml:"persons_type"`
}

type Award struct {
	ID            string `xml:"id"`
	Title         string `xml:"title"`
	Description   string `xml:"description"`
	Place         string `xml:"place"`
	URLToResource string `xml:"url_to_resource"`
	ImageID       string `xml:"image_id"`
	StartYear     string `xml:"start_year"`
	EndYear       string `xml:"end_year"`
}

type Awards struct {
	Award []Award `xml:"award"`
}

type PersonInfo struct {
	PersonID     string `xml:"person_id"`
	PersonTypeID string `xml:"person_type_id"`
	ExtraInfo    string `xml:"extra_info"`
}

type PersonsInfo struct {
	Person []PersonInfo `xml:"person"`
}

type CompilationGenre struct {
	ContentGenreID string `xml:"content_genre_id"`
}

type CompilationGenres struct {
	Genre []CompilationGenre `xml:"genre"`
}

type CompilationAward struct {
	CompilationID string `xml:"compilation_id"`
	AwardID       string `xml:"award_id"`
	Year          string `xml:"year"`
}

type CompilationAwards struct {
	Award []CompilationAward `xml:"award"`
}

type Files struct {
	File File `xml:"file"`
}

type Poster struct {
	ID            string `xml:"id"`
	CompilationID string `xml:"compilation_id"`
	Files         Files  `xml:"files"`
}

type PostersInfo struct {
	Poster Poster `xml:"poster"`
}

type CompilationSeasonAdditionalData struct {
	ID                             string `xml:"id"`
	Title                          string `xml:"title"`
	AdditionalDataTypeID           string `xml:"additional_data_type_id"`
	CompilationSeasonDescriptionID string `xml:"compilation_season_description_id"`
	Duration                       string `xml:"duration"`
}

type SeasonDescrAdditionalData struct {
	CompilationSeasonAdditionalData CompilationSeasonAdditionalData `xml:"compilation_season_additional_data"`
}

type SeasonDescription2 struct {
	Season         string                    `xml:"season"`
	Description    string                    `xml:"description"`
	AdditionalData SeasonDescrAdditionalData `xml:"additional_data"`
}

type SeasonDescription struct {
	SeasonDescription []SeasonDescription2 `xml:"season_description"`
}

type ReleaseDate struct {
	ReleaseDateBegin string `xml:"release_date_begin"`
	ReleaseDateEnd   string `xml:"release_date_end"`
}

type ReleaseDates struct {
	ReleaseDate []ReleaseDate `xml:"release_date"`
}

type CompilationAdditionalData struct {
	ID                   string `xml:"id"`
	Title                string `xml:"title"`
	AdditionalDataTypeID string `xml:"additional_data_type_id"`
	CompilationID        string `xml:"compilation_id"`
	Duration             string `xml:"duration"`
}

type CompAdditionalData struct {
	CompilationAdditionalData []CompilationAdditionalData `xml:"compilation_additional_data"`
}

type Category struct {
	ContentCategoryID string `xml:"content_category_id"`
}

type Categories struct {
	Category []Category `xml:"category"`
}

type Tima struct {
	Item []string `xml:"item"`
}

type ProductionCompany struct {
	Item string `xml:"item"`
}

type Theme struct {
	Item []string `xml:"item"`
}

type Audience struct {
	Item []string `xml:"item"`
}

type MoodInfo struct {
	Item []string `xml:"item"`
}

type Quality struct {
	Item []string `xml:"item"`
}

type About struct {
	Item []string `xml:"item"`
}

type Basis struct {
	Item []string `xml:"item"`
}

type Place struct {
	Item []string `xml:"item"`
}

type Subgenre struct {
	Item []string `xml:"item"`
}

type Gender struct {
	Item []string `xml:"item"`
}

type Character struct {
	Item string `xml:"item"`
}

type CategoryProperty struct {
	Item string `xml:"item"`
}

type Check struct {
	Item string `xml:"item"`
}

type Restrictions struct {
	Item string `xml:"item"`
}

type Compilation struct {
	ID                string             `xml:"id"`
	Title             string             `xml:"title"`
	Description       string             `xml:"description"`
	Restrict          string             `xml:"restrict"`
	OrigCountry       string             `xml:"orig_country"`
	KinopoiskID       string             `xml:"kinopoisk_id"`
	ImdbRating        string             `xml:"imdb_rating"`
	KinopoiskRating   string             `xml:"kinopoisk_rating"`
	UsaBoxOffice      string             `xml:"usa_box_office"`
	WorldBoxOffice    string             `xml:"world_box_office"`
	WeekendBoxOffice  string             `xml:"weekend_box_office"`
	ProductionBudget  string             `xml:"production_budget"`
	Hru               string             `xml:"hru"`
	FirstSeriesDate   string             `xml:"first_series_date"`
	LastSeriesDate    string             `xml:"last_series_date"`
	Persons           PersonsInfo        `xml:"persons" pg:"jsonb"`
	Genres            []string           `xml:"genres>genre>content_genre_id" pg:"array"`
	Awards            CompilationAwards  `xml:"awards" pg:"jsonb"`
	Posters           PostersInfo        `xml:"posters" pg:"jsonb"`
	SeasonDescription SeasonDescription  `xml:"season_description" pg:"jsonb"`
	ReleaseDates      ReleaseDates       `xml:"release_dates" pg:"jsonb"`
	AdditionalData    CompAdditionalData `xml:"additional_data" pg:"jsonb"`
	Categories        []string           `xml:"categories>category>content_category_id" pg:"array"`
	Tima              Tima               `xml:"tima" pg:"jsonb"`
	ProductionCompany ProductionCompany  `xml:"production_company" pg:"jsonb"`
	Theme             Theme              `xml:"theme" pg:"jsonb"`
	Audience          Audience           `xml:"audience" pg:"jsonb"`
	Mood              MoodInfo           `xml:"mood" pg:"jsonb"`
	Quality           Quality            `xml:"quality" pg:"jsonb"`
	About             About              `xml:"about" pg:"jsonb"`
	Basis             Basis              `xml:"basis" pg:"jsonb"`
	Place             Place              `xml:"place" pg:"jsonb"`
	Fake              string             `xml:"fake"`
	Subgenre          Subgenre           `xml:"subgenre" pg:"jsonb"`
	Gender            Gender             `xml:"gender" pg:"jsonb"`
	Character         Character          `xml:"character" pg:"jsonb"`
	CategoryProperty  CategoryProperty   `xml:"category_property" pg:"jsonb"`
	Check             Check              `xml:"check" pg:"jsonb"`
	Restrictions      Restrictions       `xml:"restrictions" pg:"jsonb"`
}

type Compilations struct {
	Compilation []Compilation `xml:"compilation"`
}

type Country struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Name  string `xml:"name"`
}

type Countries struct {
	Country []Country
}

type Instrument struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Info  string `xml:"info"`
}

type Instruments struct {
	Instrument []Instrument `xml:"instrument"`
}

type Mood struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Info  string `xml:"info"`
}

type Moods struct {
	Mood []Mood `xml:"mood"`
}

type RusVersion struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type RusVersions struct {
	RusVersion []RusVersion `xml:"rus_version"`
}

type ContentFormat struct {
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	TypeID      string `xml:"type_id"`
	Group       string `xml:"group"`
}

type ContentFormats struct {
	ContentFormat []ContentFormat `xml:"content_format"`
}

type Lang struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Name  string `xml:"name"`
}

type Langs struct {
	Lang []Lang `xml:"lang"`
}

type ContentType struct {
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type ContentTypes struct {
	ContentType []ContentType `xml:"content_type"`
}

type Property struct {
	ID       string `xml:"id"`
	Title    string `xml:"title"`
	XMLField string `xml:"xml_field"`
	Unique   string `xml:"unique"`
}

type Properties struct {
	Property []Property `xml:"property"`
}

type PropertyValue struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type PropertyValues struct {
	PropertyValue []PropertyValue `xml:"property_value"`
}

type PromoInfo struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type Promo struct {
	Promo []PromoInfo `xml:"promo"`
}

type ContentAdditionalDataType struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type AdditionalDataType struct {
	ContentAdditionalDataType []ContentAdditionalDataType `xml:"content_additional_data_type"`
}

type ContentGenreInfo struct {
	ContentID      string `xml:"content_id"`
	ContentGenreID string `xml:"content_genre_id"`
	Priority       string `xml:"priority"`
}

type ContentGenresInfo struct {
	ContentGenre []ContentGenreInfo `xml:"content_genre"`
}

type ContentAward struct {
	ContentID string `xml:"content_id"`
	AwardID   string `xml:"award_id"`
	Year      string `xml:"year"`
}

type ContentAwards struct {
	Text  string         `xml:",chardata"`
	Award []ContentAward `xml:"award"`
}

type ContentMoods struct {
	Mood []string `xml:"mood"`
}

type ContentCategoriesInfo struct {
	ContentCategory []string `xml:"content_category"`
}

type ContentRusVersions struct {
	RusVersion []string `xml:"rus_version"`
}

type ContentAdditionalData struct {
	ID                   string `xml:"id"`
	Title                string `xml:"title"`
	AdditionalDataTypeID string `xml:"additional_data_type_id"`
	ContentID            string `xml:"content_id"`
	Duration             string `xml:"duration"`
}

type AdditionalData struct {
	ContentAdditionalData []ContentAdditionalData `xml:"content_additional_data"`
}

type ContentFiles struct {
	File []File `xml:"file"`
}

type Thumb struct {
	ID        string       `xml:"id"`
	ContentID string       `xml:"content_id"`
	Files     ContentFiles `xml:"files"`
}

type Thumbs struct {
	Thumb Thumb `xml:"thumb"`
}

type Subscriptions struct {
	Subscription []string `xml:"subscription"`
}

type RightTypeList struct {
	RightType string `xml:"right_type"`
}

type RightHolder struct {
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type RightHolders struct {
	RightHolder []RightHolder `xml:"right_holder"`
}

type ContentCheck struct {
	Item []string `xml:"item"`
}

type ContentRestrictions struct {
	Item []string `xml:"item"`
}

type TestProperty struct {
	Item string `xml:"item"`
}

type LocalizationCheck struct {
	Item string `xml:"item"`
}

type ContentProductionCompany struct {
	Item []string `xml:"item"`
}

type ContentCharacter struct {
	Item []string `xml:"item"`
}

type Franchise struct {
	Item string `xml:"item"`
}

type Shild struct {
	Item string `xml:"item"`
}

type NewContent struct {
	Item string `xml:"item"`
}

type Content struct {
	ID                string                   `xml:"id"`
	Title             string                   `xml:"title"`
	OrigTitle         string                   `xml:"orig_title"`
	Compilation       string                   `xml:"compilation"`
	CopyrightOverlay  string                   `xml:"copyright_overlay"`
	Season            string                   `xml:"season"`
	Episode           string                   `xml:"episode"`
	Duration          string                   `xml:"duration"`
	PreviewID         string                   `xml:"preview_id"`
	Restrict          string                   `xml:"restrict"`
	PrSiteID          string                   `xml:"pr_site_id"`
	OriginCountry     string                   `xml:"origin_country"`
	ReleaseDate       string                   `xml:"release_date"`
	KinopoiskID       string                   `xml:"kinopoisk_id"`
	CompilationID     string                   `xml:"compilation_id"`
	PromoID           string                   `xml:"promo_id"`
	ImdbRating        string                   `xml:"imdb_rating"`
	KinopoiskRating   string                   `xml:"kinopoisk_rating"`
	UsaBoxOffice      string                   `xml:"usa_box_office"`
	WorldBoxOffice    string                   `xml:"world_box_office"`
	WeekendBoxOffice  string                   `xml:"weekend_box_office"`
	ProductionBudget  string                   `xml:"production_budget"`
	YoutubeTrailerURL string                   `xml:"youtube_trailer_url"`
	DateInsert        string                   `xml:"date_insert"`
	DescriptionAlt    string                   `xml:"description_alt"`
	SharingDisabled   string                   `xml:"sharing_disabled"`
	RightDateEnd      string                   `xml:"right_date_end"`
	DrmOnly           string                   `xml:"drm_only"`
	Instruments       string                   `xml:"instruments"`
	ContentGenres     ContentGenresInfo        `xml:"content_genres" pg:"jsonb"`
	Persons           PersonsInfo              `xml:"persons" pg:"jsonb"`
	Awards            ContentAwards            `xml:"awards" pg:"jsonb"`
	Moods             ContentMoods             `xml:"moods" pg:"jsonb"`
	ContentCategories ContentCategoriesInfo    `xml:"content_categories" pg:"jsonb"`
	Posters           PostersInfo              `xml:"posters" pg:"jsonb"`
	Trailers          string                   `xml:"trailers"`
	RusVersions       ContentRusVersions       `xml:"rus_versions" pg:"jsonb"`
	AdditionalData    AdditionalData           `xml:"additional_data" pg:"jsonb"`
	Thumbs            Thumbs                   `xml:"thumbs" pg:"jsonb"`
	Subscriptions     Subscriptions            `xml:"subscriptions" pg:"jsonb"`
	RightTypeList     RightTypeList            `xml:"right_type_list" pg:"jsonb"`
	RightType         string                   `xml:"right_type"`
	RightHolders      RightHolders             `xml:"right_holders" pg:"jsonb"`
	Mp4               string                   `xml:"mp4"`
	Tima              Tima                     `xml:"tima" pg:"jsonb"`
	Theme             Theme                    `xml:"theme" pg:"jsonb"`
	Audience          Audience                 `xml:"audience" pg:"jsonb"`
	Subgenre          Subgenre                 `xml:"subgenre" pg:"jsonb"`
	Mood              MoodInfo                 `xml:"mood" pg:"jsonb"`
	Quality           Quality                  `xml:"quality" pg:"jsonb"`
	About             About                    `xml:"about" pg:"jsonb"`
	Basis             Basis                    `xml:"basis" pg:"jsonb"`
	CategoryProperty  CategoryProperty         `xml:"category_property" pg:"jsonb"`
	Check             ContentCheck             `xml:"check" pg:"jsonb"`
	Place             Place                    `xml:"place" pg:"jsonb"`
	Restrictions      ContentRestrictions      `xml:"restrictions" pg:"jsonb"`
	TestProperty      TestProperty             `xml:"test_property" pg:"jsonb"`
	LocalizationCheck LocalizationCheck        `xml:"localization_check" pg:"jsonb"`
	Gender            Gender                   `xml:"gender" pg:"jsonb"`
	ProductionCompany ContentProductionCompany `xml:"production_company" pg:"jsonb"`
	Character         ContentCharacter         `xml:"character" pg:"jsonb"`
	Franchise         Franchise                `xml:"franchise" pg:"jsonb"`
	Shild             Shild                    `xml:"shild" pg:"jsonb"`
	NewContent        NewContent               `xml:"new_content" pg:"jsonb"`
}

type Contents struct {
	Content []Content `xml:"content"`
}

type MyShows struct {
	XMLName                 xml.Name           `xml:"xml"`
	ContentCategories       ContentCategories  `xml:"content_categories"`
	ContentCategory2s       ContentCategories  `xml:"content_category_2s"`
	ContentCategory3s       ContentCategories  `xml:"content_category_3s"`
	ContentGenres           Genres             `xml:"content_genres"`
	Persons                 Persons            `xml:"persons"` // :(
	PersonsType             PersonsType        `xml:"persons_type"`
	Awards                  Awards             `xml:"awards"`
	Compilations            Compilations       `xml:"compilations"` // :(
	Countries               Countries          `xml:"countries"`
	Instruments             Instruments        `xml:"instruments"`
	Moods                   Moods              `xml:"moods"`
	RusVersions             RusVersions        `xml:"rus_versions"`
	ContentFormats          ContentFormats     `xml:"content_formats"`
	Langs                   Langs              `xml:"langs"`
	ContentTypes            ContentTypes       `xml:"content_types"`
	Properties              Properties         `xml:"properties"`
	PropertyValues          PropertyValues     `xml:"property_values"`
	Promo                   Promo              `xml:"promo"`
	AdditionalDataType      AdditionalDataType `xml:"additional_data_type"`
	Contents                Contents           `xml:"contents"` // :(
	ActionLogSequenceNumber string             `xml:"action_log_sequence_number"`
}
