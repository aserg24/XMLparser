package structureXML

import "encoding/xml"

// XML was generated 2019-07-04 18:01:39 by DESKTOP-70UBMEV\Arog on DESKTOP-70UBMEV.
type XML1 struct {
	XMLName           xml.Name `xml:"xml"`
	Text              string   `xml:",chardata"`
	ContentCategories struct {
		Text            string `xml:",chardata"`
		ContentCategory []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Hru         string `xml:"hru"`
		} `xml:"content_category"`
	} `xml:"content_categories"`
	ContentCategory2s struct {
		Text            string `xml:",chardata"`
		ContentCategory []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Hru         string `xml:"hru"`
		} `xml:"content_category"`
	} `xml:"content_category_2s"`
	ContentCategory3s struct {
		Text            string `xml:",chardata"`
		ContentCategory []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Hru         string `xml:"hru"`
		} `xml:"content_category"`
	} `xml:"content_category_3s"`
	ContentGenres struct {
		Text         string `xml:",chardata"`
		ContentGenre []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			CategoryID  string `xml:"category_id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Hru         string `xml:"hru"`
		} `xml:"content_genre"`
	} `xml:"content_genres"`
	Persons struct {
		Text   string `xml:",chardata"`
		Person []struct {
			Text    string `xml:",chardata"`
			ID      string `xml:"id"`
			Posters struct {
				Text string `xml:",chardata"`
				File []struct {
					Text   string `xml:",chardata"`
					URL    string `xml:"url"`
					ID     string `xml:"id"`
					Format string `xml:"format"`
				} `xml:"file"`
			} `xml:"posters"`
			Name        string `xml:"name"`
			IsStar      string `xml:"is_star"`
			EngTitle    string `xml:"eng_title"`
			GenTitle    string `xml:"gen_title"`
			Hru         string `xml:"hru"`
			FakeHru     string `xml:"fake_hru"`
			BlogPost    string `xml:"blog_post"`
			BlogLink    string `xml:"blog_link"`
			BlogAdded   string `xml:"blog_added"`
			Bio         string `xml:"bio"`
			Facts       string `xml:"facts"`
			Rss         string `xml:"rss"`
			KinopoiskID string `xml:"kinopoisk_id"`
		} `xml:"person"`
	} `xml:"persons"`
	PersonsType struct {
		Text       string `xml:",chardata"`
		PersonType []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
		} `xml:"person_type"`
	} `xml:"persons_type"`
	Awards struct {
		Text  string `xml:",chardata"`
		Award []struct {
			Text          string `xml:",chardata"`
			ID            string `xml:"id"`
			Title         string `xml:"title"`
			Description   string `xml:"description"`
			Place         string `xml:"place"`
			URLToResource string `xml:"url_to_resource"`
			ImageID       string `xml:"image_id"`
			StartYear     string `xml:"start_year"`
			EndYear       string `xml:"end_year"`
		} `xml:"award"`
	} `xml:"awards"`
	Compilations struct {
		Text        string `xml:",chardata"`
		Compilation []struct {
			Text             string `xml:",chardata"`
			ID               string `xml:"id"`
			Title            string `xml:"title"`
			Description      string `xml:"description"`
			Restrict         string `xml:"restrict"`
			OrigCountry      string `xml:"orig_country"`
			KinopoiskID      string `xml:"kinopoisk_id"`
			ImdbRating       string `xml:"imdb_rating"`
			KinopoiskRating  string `xml:"kinopoisk_rating"`
			UsaBoxOffice     string `xml:"usa_box_office"`
			WorldBoxOffice   string `xml:"world_box_office"`
			WeekendBoxOffice string `xml:"weekend_box_office"`
			ProductionBudget string `xml:"production_budget"`
			Hru              string `xml:"hru"`
			FirstSeriesDate  string `xml:"first_series_date"`
			LastSeriesDate   string `xml:"last_series_date"`
			Persons          struct {
				Text   string `xml:",chardata"`
				Person []struct {
					Text         string `xml:",chardata"`
					PersonID     string `xml:"person_id"`
					PersonTypeID string `xml:"person_type_id"`
					ExtraInfo    string `xml:"extra_info"`
				} `xml:"person"`
			} `xml:"persons"`
			Genres struct {
				Text  string `xml:",chardata"`
				Genre []struct {
					Text           string `xml:",chardata"`
					ContentGenreID string `xml:"content_genre_id"`
				} `xml:"genre"`
			} `xml:"genres"`
			Awards struct {
				Text  string `xml:",chardata"`
				Award []struct {
					Text          string `xml:",chardata"`
					CompilationID string `xml:"compilation_id"`
					AwardID       string `xml:"award_id"`
					Year          string `xml:"year"`
				} `xml:"award"`
			} `xml:"awards"`
			Posters struct {
				Text   string `xml:",chardata"`
				Poster struct {
					Text          string `xml:",chardata"`
					ID            string `xml:"id"`
					CompilationID string `xml:"compilation_id"`
					Files         struct {
						Text string `xml:",chardata"`
						File struct {
							Text   string `xml:",chardata"`
							URL    string `xml:"url"`
							ID     string `xml:"id"`
							Format string `xml:"format"`
						} `xml:"file"`
					} `xml:"files"`
				} `xml:"poster"`
			} `xml:"posters"`
			SeasonDescription struct {
				Text              string `xml:",chardata"`
				SeasonDescription []struct {
					Text           string `xml:",chardata"`
					Season         string `xml:"season"`
					Description    string `xml:"description"`
					AdditionalData struct {
						Text                            string `xml:",chardata"`
						CompilationSeasonAdditionalData struct {
							Text                           string `xml:",chardata"`
							ID                             string `xml:"id"`
							Title                          string `xml:"title"`
							AdditionalDataTypeID           string `xml:"additional_data_type_id"`
							CompilationSeasonDescriptionID string `xml:"compilation_season_description_id"`
							Duration                       string `xml:"duration"`
						} `xml:"compilation_season_additional_data"`
					} `xml:"additional_data"`
				} `xml:"season_description"`
			} `xml:"season_description"`
			ReleaseDates struct {
				Text        string `xml:",chardata"`
				ReleaseDate []struct {
					Text             string `xml:",chardata"`
					ReleaseDateBegin string `xml:"release_date_begin"`
					ReleaseDateEnd   string `xml:"release_date_end"`
				} `xml:"release_date"`
			} `xml:"release_dates"`
			AdditionalData struct {
				Text                      string `xml:",chardata"`
				CompilationAdditionalData []struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"id"`
					Title                string `xml:"title"`
					AdditionalDataTypeID string `xml:"additional_data_type_id"`
					CompilationID        string `xml:"compilation_id"`
					Duration             string `xml:"duration"`
				} `xml:"compilation_additional_data"`
			} `xml:"additional_data"`
			Categories struct {
				Text     string `xml:",chardata"`
				Category []struct {
					Text              string `xml:",chardata"`
					ContentCategoryID string `xml:"content_category_id"`
				} `xml:"category"`
			} `xml:"categories"`
			Tima struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"tima"`
			ProductionCompany struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"production_company"`
			Theme struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"theme"`
			Audience struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"audience"`
			Mood struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"mood"`
			Quality struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"quality"`
			About struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"about"`
			Basis struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"basis"`
			Place struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"place"`
			Fake     string `xml:"fake"`
			Subgenre struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"subgenre"`
			Gender struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"gender"`
			Character struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"character"`
			CategoryProperty struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"category_property"`
			Check struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"check"`
			Restrictions struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"restrictions"`
		} `xml:"compilation"`
	} `xml:"compilations"`
	Countries struct {
		Text    string `xml:",chardata"`
		Country []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
			Name  string `xml:"name"`
		} `xml:"country"`
	} `xml:"countries"`
	Instruments struct {
		Text       string `xml:",chardata"`
		Instrument []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
			Info  string `xml:"info"`
		} `xml:"instrument"`
	} `xml:"instruments"`
	Moods struct {
		Text string `xml:",chardata"`
		Mood []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
			Info  string `xml:"info"`
		} `xml:"mood"`
	} `xml:"moods"`
	RusVersions struct {
		Text       string `xml:",chardata"`
		RusVersion []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
		} `xml:"rus_version"`
	} `xml:"rus_versions"`
	ContentFormats struct {
		Text          string `xml:",chardata"`
		ContentFormat []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			TypeID      string `xml:"type_id"`
			Group       string `xml:"group"`
		} `xml:"content_format"`
	} `xml:"content_formats"`
	Langs struct {
		Text string `xml:",chardata"`
		Lang []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
			Name  string `xml:"name"`
		} `xml:"lang"`
	} `xml:"langs"`
	ContentTypes struct {
		Text        string `xml:",chardata"`
		ContentType []struct {
			Text        string `xml:",chardata"`
			ID          string `xml:"id"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
		} `xml:"content_type"`
	} `xml:"content_types"`
	Properties struct {
		Text     string `xml:",chardata"`
		Property []struct {
			Text     string `xml:",chardata"`
			ID       string `xml:"id"`
			Title    string `xml:"title"`
			XMLField string `xml:"xml_field"`
			Unique   string `xml:"unique"`
		} `xml:"property"`
	} `xml:"properties"`
	PropertyValues struct {
		Text          string `xml:",chardata"`
		PropertyValue []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
		} `xml:"property_value"`
	} `xml:"property_values"`
	Promo struct {
		Text  string `xml:",chardata"`
		Promo []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
		} `xml:"promo"`
	} `xml:"promo"`
	AdditionalDataType struct {
		Text                      string `xml:",chardata"`
		ContentAdditionalDataType []struct {
			Text  string `xml:",chardata"`
			ID    string `xml:"id"`
			Title string `xml:"title"`
		} `xml:"content_additional_data_type"`
	} `xml:"additional_data_type"`
	Contents struct {
		Text    string `xml:",chardata"`
		Content []struct {
			Text              string `xml:",chardata"`
			ID                string `xml:"id"`
			Title             string `xml:"title"`
			OrigTitle         string `xml:"orig_title"`
			Compilation       string `xml:"compilation"`
			CopyrightOverlay  string `xml:"copyright_overlay"`
			Season            string `xml:"season"`
			Episode           string `xml:"episode"`
			Duration          string `xml:"duration"`
			PreviewID         string `xml:"preview_id"`
			Restrict          string `xml:"restrict"`
			PrSiteID          string `xml:"pr_site_id"`
			OriginCountry     string `xml:"origin_country"`
			ReleaseDate       string `xml:"release_date"`
			KinopoiskID       string `xml:"kinopoisk_id"`
			CompilationID     string `xml:"compilation_id"`
			PromoID           string `xml:"promo_id"`
			ImdbRating        string `xml:"imdb_rating"`
			KinopoiskRating   string `xml:"kinopoisk_rating"`
			UsaBoxOffice      string `xml:"usa_box_office"`
			WorldBoxOffice    string `xml:"world_box_office"`
			WeekendBoxOffice  string `xml:"weekend_box_office"`
			ProductionBudget  string `xml:"production_budget"`
			YoutubeTrailerURL string `xml:"youtube_trailer_url"`
			DateInsert        string `xml:"date_insert"`
			DescriptionAlt    string `xml:"description_alt"`
			SharingDisabled   string `xml:"sharing_disabled"`
			RightDateEnd      string `xml:"right_date_end"`
			DrmOnly           string `xml:"drm_only"`
			Instruments       string `xml:"instruments"`
			ContentGenres     struct {
				Text         string `xml:",chardata"`
				ContentGenre []struct {
					Text           string `xml:",chardata"`
					ContentID      string `xml:"content_id"`
					ContentGenreID string `xml:"content_genre_id"`
					Priority       string `xml:"priority"`
				} `xml:"content_genre"`
			} `xml:"content_genres"`
			Persons struct {
				Text   string `xml:",chardata"`
				Person []struct {
					Text         string `xml:",chardata"`
					PersonID     string `xml:"person_id"`
					PersonTypeID string `xml:"person_type_id"`
					ExtraInfo    string `xml:"extra_info"`
				} `xml:"person"`
			} `xml:"persons"`
			Awards struct {
				Text  string `xml:",chardata"`
				Award []struct {
					Text      string `xml:",chardata"`
					ContentID string `xml:"content_id"`
					AwardID   string `xml:"award_id"`
					Year      string `xml:"year"`
				} `xml:"award"`
			} `xml:"awards"`
			Moods struct {
				Text string   `xml:",chardata"`
				Mood []string `xml:"mood"`
			} `xml:"moods"`
			ContentCategories struct {
				Text            string   `xml:",chardata"`
				ContentCategory []string `xml:"content_category"`
			} `xml:"content_categories"`
			Posters struct {
				Text   string `xml:",chardata"`
				Poster struct {
					Text      string `xml:",chardata"`
					ID        string `xml:"id"`
					ContentID string `xml:"content_id"`
					Files     struct {
						Text string `xml:",chardata"`
						File struct {
							Text   string `xml:",chardata"`
							URL    string `xml:"url"`
							ID     string `xml:"id"`
							Format string `xml:"format"`
						} `xml:"file"`
					} `xml:"files"`
				} `xml:"poster"`
			} `xml:"posters"`
			Trailers    string `xml:"trailers"`
			RusVersions struct {
				Text       string   `xml:",chardata"`
				RusVersion []string `xml:"rus_version"`
			} `xml:"rus_versions"`
			AdditionalData struct {
				Text                  string `xml:",chardata"`
				ContentAdditionalData []struct {
					Text                 string `xml:",chardata"`
					ID                   string `xml:"id"`
					Title                string `xml:"title"`
					AdditionalDataTypeID string `xml:"additional_data_type_id"`
					ContentID            string `xml:"content_id"`
					Duration             string `xml:"duration"`
				} `xml:"content_additional_data"`
			} `xml:"additional_data"`
			Thumbs struct {
				Text  string `xml:",chardata"`
				Thumb struct {
					Text      string `xml:",chardata"`
					ID        string `xml:"id"`
					ContentID string `xml:"content_id"`
					Files     struct {
						Text string `xml:",chardata"`
						File []struct {
							Text   string `xml:",chardata"`
							URL    string `xml:"url"`
							ID     string `xml:"id"`
							Format string `xml:"format"`
						} `xml:"file"`
					} `xml:"files"`
				} `xml:"thumb"`
			} `xml:"thumbs"`
			Subscriptions struct {
				Text         string   `xml:",chardata"`
				Subscription []string `xml:"subscription"`
			} `xml:"subscriptions"`
			RightTypeList struct {
				Text      string `xml:",chardata"`
				RightType string `xml:"right_type"`
			} `xml:"right_type_list"`
			RightType    string `xml:"right_type"`
			RightHolders struct {
				Text        string `xml:",chardata"`
				RightHolder []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id"`
					Title string `xml:"title"`
				} `xml:"right_holder"`
			} `xml:"right_holders"`
			Mp4  string `xml:"mp4"`
			Tima struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"tima"`
			Theme struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"theme"`
			Audience struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"audience"`
			Subgenre struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"subgenre"`
			Mood struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"mood"`
			Quality struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"quality"`
			About struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"about"`
			Basis struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"basis"`
			CategoryProperty struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"category_property"`
			Check struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"check"`
			Place struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"place"`
			Restrictions struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"restrictions"`
			TestProperty struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"test_property"`
			LocalizationCheck struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"localization_check"`
			Gender struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"gender"`
			ProductionCompany struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"production_company"`
			Character struct {
				Text string   `xml:",chardata"`
				Item []string `xml:"item"`
			} `xml:"character"`
			Franchise struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"franchise"`
			Shild struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"shild"`
			NewContent struct {
				Text string `xml:",chardata"`
				Item string `xml:"item"`
			} `xml:"new_content"`
		} `xml:"content"`
	} `xml:"contents"`
	ActionLogSequenceNumber string `xml:"action_log_sequence_number"`
}
