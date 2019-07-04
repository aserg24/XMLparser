package structureXML

import "encoding/xml"

type ContentCategory struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Hru         string `xml:"hru"`
}

type ContentCategories struct {
	Text            string            `xml:",chardata"`
	ContentCategory []ContentCategory `xml:"content_category"`
}

type Genre struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id"`
	CategoryID  string `xml:"category_id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Hru         string `xml:"hru"`
}

type Genres struct {
	Text         string  `xml:",chardata"`
	ContentGenre []Genre `xml:"content_genre"`
}

type File struct {
	Text   string `xml:",chardata"`
	URL    string `xml:"url"`
	ID     string `xml:"id"`
	Format string `xml:"format"`
}

type Posters struct {
	Text string `xml:",chardata"`
	File []File `xml:"file"`
}

type Person struct {
	Text        string  `xml:",chardata"`
	ID          string  `xml:"id"`
	Posters     Posters `xml:"posters"`
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
	Text   string   `xml:",chardata"`
	Person []Person `xml:"person"`
}

type PersonType struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type PersonsType struct {
	Text       string       `xml:",chardata"`
	PersonType []PersonType `xml:"persons_type"`
}

type Award struct {
	Text          string `xml:",chardata"`
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
	Text  string  `xml:",chardata"`
	Award []Award `xml:"award"`
}

type PersonInfo struct {
	Text         string `xml:",chardata"`
	PersonID     string `xml:"person_id"`
	PersonTypeID string `xml:"person_type_id"`
	ExtraInfo    string `xml:"extra_info"`
}

type PersonsInfo struct {
	Text   string       `xml:",chardata"`
	Person []PersonInfo `xml:"person"`
}

type CompilationGenre struct {
	Text           string `xml:",chardata"`
	ContentGenreID string `xml:"content_genre_id"`
}

type CompilationGenres struct {
	Text  string             `xml:",chardata"`
	Genre []CompilationGenre `xml:"genre"`
}

type CompilationAward struct {
	Text          string `xml:",chardata"`
	CompilationID string `xml:"compilation_id"`
	AwardID       string `xml:"award_id"`
	Year          string `xml:"year"`
}

type CompilationAwards struct {
	Text  string             `xml:",chardata"`
	Award []CompilationAward `xml:"award"`
}

type Files struct {
	Text string `xml:",chardata"`
	File File   `xml:"file"`
}

type Poster struct {
	Text          string `xml:",chardata"`
	ID            string `xml:"id"`
	CompilationID string `xml:"compilation_id"`
	Files         Files  `xml:"files"`
}

type PostersInfo struct {
	Text   string `xml:",chardata"`
	Poster Poster `xml:"poster"`
}

type CompilationSeasonAdditionalData struct {
	Text                           string `xml:",chardata"`
	ID                             string `xml:"id"`
	Title                          string `xml:"title"`
	AdditionalDataTypeID           string `xml:"additional_data_type_id"`
	CompilationSeasonDescriptionID string `xml:"compilation_season_description_id"`
	Duration                       string `xml:"duration"`
}

type SeasonDescrAdditionalData struct {
	Text                            string                          `xml:",chardata"`
	CompilationSeasonAdditionalData CompilationSeasonAdditionalData `xml:"compilation_season_additional_data"`
}

type SeasonDescription2 struct {
	Text           string                    `xml:",chardata"`
	Season         string                    `xml:"season"`
	Description    string                    `xml:"description"`
	AdditionalData SeasonDescrAdditionalData `xml:"additional_data"`
}

type SeasonDescription struct {
	Text              string               `xml:",chardata"`
	SeasonDescription []SeasonDescription2 `xml:"season_description"`
}

type ReleaseDate struct {
	Text             string `xml:",chardata"`
	ReleaseDateBegin string `xml:"release_date_begin"`
	ReleaseDateEnd   string `xml:"release_date_end"`
}

type ReleaseDates struct {
	Text        string        `xml:",chardata"`
	ReleaseDate []ReleaseDate `xml:"release_date"`
}

type CompilationAdditionalData struct {
	Text                 string `xml:",chardata"`
	ID                   string `xml:"id"`
	Title                string `xml:"title"`
	AdditionalDataTypeID string `xml:"additional_data_type_id"`
	CompilationID        string `xml:"compilation_id"`
	Duration             string `xml:"duration"`
}

type CompAdditionalData struct {
	Text                      string                      `xml:",chardata"`
	CompilationAdditionalData []CompilationAdditionalData `xml:"compilation_additional_data"`
}

type Category struct {
	Text              string `xml:",chardata"`
	ContentCategoryID string `xml:"content_category_id"`
}

type Categories struct {
	Text     string     `xml:",chardata"`
	Category []Category `xml:"category"`
}

type Tima struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type ProductionCompany struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type Theme struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Audience struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type MoodInfo struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Quality struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type About struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Basis struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Place struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Subgenre struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Gender struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Character struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type CategoryProperty struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type Check struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type Restrictions struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type Compilation struct {
	Text              string             `xml:",chardata"`
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
	Persons           PersonsInfo        `xml:"persons"`
	Genres            CompilationGenres  `xml:"genres"`
	Awards            CompilationAwards  `xml:"awards"`
	Posters           PostersInfo        `xml:"posters"`
	SeasonDescription SeasonDescription  `xml:"season_description"`
	ReleaseDates      ReleaseDates       `xml:"release_dates"`
	AdditionalData    CompAdditionalData `xml:"additional_data"`
	Categories        Categories         `xml:"categories"`
	Tima              Tima               `xml:"tima"`
	ProductionCompany ProductionCompany  `xml:"production_company"`
	Theme             Theme              `xml:"theme"`
	Audience          Audience           `xml:"audience"`
	Mood              MoodInfo           `xml:"mood"`
	Quality           Quality            `xml:"quality"`
	About             About              `xml:"about"`
	Basis             Basis              `xml:"basis"`
	Place             Place              `xml:"place"`
	Fake              string             `xml:"fake"`
	Subgenre          Subgenre           `xml:"subgenre"`
	Gender            Gender             `xml:"gender"`
	Character         Character          `xml:"character"`
	CategoryProperty  CategoryProperty   `xml:"category_property"`
	Check             Check              `xml:"check"`
	Restrictions      Restrictions       `xml:"restrictions"`
}

type Compilations struct {
	Text        string        `xml:",chardata"`
	Compilation []Compilation `xml:"compilation"`
}

type Country []struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Name  string `xml:"name"`
}

type Countries struct {
	Text    string `xml:",chardata"`
	Country []Country
}

type Instrument []struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Info  string `xml:"info"`
}

type Instruments struct {
	Text       string       `xml:",chardata"`
	Instrument []Instrument `xml:"instrument"`
}

type Mood struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Info  string `xml:"info"`
}

type Moods struct {
	Text string `xml:",chardata"`
	Mood []Mood `xml:"mood"`
}

type RusVersion struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type RusVersions struct {
	Text       string       `xml:",chardata"`
	RusVersion []RusVersion `xml:"rus_version"`
}

type ContentFormat struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	TypeID      string `xml:"type_id"`
	Group       string `xml:"group"`
}

type ContentFormats struct {
	Text          string          `xml:",chardata"`
	ContentFormat []ContentFormat `xml:"content_format"`
}

type Lang struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
	Name  string `xml:"name"`
}

type Langs struct {
	Text string `xml:",chardata"`
	Lang []Lang `xml:"lang"`
}

type ContentType struct {
	Text        string `xml:",chardata"`
	ID          string `xml:"id"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
}

type ContentTypes struct {
	Text        string        `xml:",chardata"`
	ContentType []ContentType `xml:"content_type"`
}

type Property struct {
	Text     string `xml:",chardata"`
	ID       string `xml:"id"`
	Title    string `xml:"title"`
	XMLField string `xml:"xml_field"`
	Unique   string `xml:"unique"`
}

type Properties struct {
	Text     string     `xml:",chardata"`
	Property []Property `xml:"property"`
}

type PropertyValue struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type PropertyValues struct {
	Text          string          `xml:",chardata"`
	PropertyValue []PropertyValue `xml:"property_value"`
}

type PromoInfo struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type Promo struct {
	Text  string      `xml:",chardata"`
	Promo []PromoInfo `xml:"promo"`
}

type ContentAdditionalDataType struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type AdditionalDataType struct {
	Text                      string                      `xml:",chardata"`
	ContentAdditionalDataType []ContentAdditionalDataType `xml:"content_additional_data_type"`
}

type ContentGenreInfo struct {
	Text           string `xml:",chardata"`
	ContentID      string `xml:"content_id"`
	ContentGenreID string `xml:"content_genre_id"`
	Priority       string `xml:"priority"`
}

type ContentGenresInfo struct {
	Text         string             `xml:",chardata"`
	ContentGenre []ContentGenreInfo `xml:"content_genre"`
}

type ContentAward struct {
	Text      string `xml:",chardata"`
	ContentID string `xml:"content_id"`
	AwardID   string `xml:"award_id"`
	Year      string `xml:"year"`
}

type ContentAwards struct {
	Text  string         `xml:",chardata"`
	Award []ContentAward `xml:"award"`
}

type ContentMoods struct {
	Text string   `xml:",chardata"`
	Mood []string `xml:"mood"`
}

type ContentCategoriesInfo struct {
	Text            string   `xml:",chardata"`
	ContentCategory []string `xml:"content_category"`
}

type ContentRusVersions struct {
	Text       string   `xml:",chardata"`
	RusVersion []string `xml:"rus_version"`
}

type ContentAdditionalData struct {
	Text                 string `xml:",chardata"`
	ID                   string `xml:"id"`
	Title                string `xml:"title"`
	AdditionalDataTypeID string `xml:"additional_data_type_id"`
	ContentID            string `xml:"content_id"`
	Duration             string `xml:"duration"`
}

type AdditionalData struct {
	Text                  string                  `xml:",chardata"`
	ContentAdditionalData []ContentAdditionalData `xml:"content_additional_data"`
}

type ContentFiles struct {
	Text string `xml:",chardata"`
	File []File `xml:"file"`
}

type Thumb struct {
	Text      string       `xml:",chardata"`
	ID        string       `xml:"id"`
	ContentID string       `xml:"content_id"`
	Files     ContentFiles `xml:"files"`
}

type Thumbs struct {
	Text  string `xml:",chardata"`
	Thumb Thumb  `xml:"thumb"`
}

type Subscriptions struct {
	Text         string   `xml:",chardata"`
	Subscription []string `xml:"subscription"`
}

type RightTypeList struct {
	Text      string `xml:",chardata"`
	RightType string `xml:"right_type"`
}

type RightHolder struct {
	Text  string `xml:",chardata"`
	ID    string `xml:"id"`
	Title string `xml:"title"`
}

type RightHolders struct {
	Text        string        `xml:",chardata"`
	RightHolder []RightHolder `xml:"right_holder"`
}

type ContentCheck struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type ContentRestrictions struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type TestProperty struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type LocalizationCheck struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type ContentProductionCompany struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type ContentCharacter struct {
	Text string   `xml:",chardata"`
	Item []string `xml:"item"`
}

type Franchise struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type Shild struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type NewContent struct {
	Text string `xml:",chardata"`
	Item string `xml:"item"`
}

type Content struct {
	Text              string                   `xml:",chardata"`
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
	ContentGenres     ContentGenresInfo        `xml:"content_genres"`
	Persons           PersonsInfo              `xml:"persons"`
	Awards            ContentAwards            `xml:"awards"`
	Moods             ContentMoods             `xml:"moods"`
	ContentCategories ContentCategoriesInfo    `xml:"content_categories"`
	Posters           PostersInfo              `xml:"posters"`
	Trailers          string                   `xml:"trailers"`
	RusVersions       ContentRusVersions       `xml:"rus_versions"`
	AdditionalData    AdditionalData           `xml:"additional_data"`
	Thumbs            Thumbs                   `xml:"thumbs"`
	Subscriptions     Subscriptions            `xml:"subscriptions"`
	RightTypeList     RightTypeList            `xml:"right_type_list"`
	RightType         string                   `xml:"right_type"`
	RightHolders      RightHolders             `xml:"right_holders"`
	Mp4               string                   `xml:"mp4"`
	Tima              Tima                     `xml:"tima"`
	Theme             Theme                    `xml:"theme"`
	Audience          Audience                 `xml:"audience"`
	Subgenre          Subgenre                 `xml:"subgenre"`
	Mood              MoodInfo                 `xml:"mood"`
	Quality           Quality                  `xml:"quality"`
	About             About                    `xml:"about"`
	Basis             Basis                    `xml:"basis"`
	CategoryProperty  CategoryProperty         `xml:"category_property"`
	Check             ContentCheck             `xml:"check"`
	Place             Place                    `xml:"place"`
	Restrictions      ContentRestrictions      `xml:"restrictions"`
	TestProperty      TestProperty             `xml:"test_property"`
	LocalizationCheck LocalizationCheck        `xml:"localization_check"`
	Gender            Gender                   `xml:"gender"`
	ProductionCompany ContentProductionCompany `xml:"production_company"`
	Character         ContentCharacter         `xml:"character"`
	Franchise         Franchise                `xml:"franchise"`
	Shild             Shild                    `xml:"shild"`
	NewContent        NewContent               `xml:"new_content"`
}

type Contents struct {
	Text    string    `xml:",chardata"`
	Content []Content `xml:"content"`
}

// XML was generated 2019-07-02 14:25:57 by DESKTOP-70UBMEV\Arog on DESKTOP-70UBMEV.
type MyShows struct {
	XMLName                 xml.Name           `xml:"xml"`
	Text                    string             `xml:",chardata"`
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
