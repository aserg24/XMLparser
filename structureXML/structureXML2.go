package structureXML

import "encoding/xml"

// XML was generated 2019-07-03 21:03:55 by DESKTOP-70UBMEV\Arog on DESKTOP-70UBMEV.
type XML struct {
	XMLName           xml.Name `xml:"xml"`
	Text              string   `xml:",chardata"`
	ContentCategories struct {
		Text            string `xml:",chardata"`
		ContentCategory []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			Hru struct {
				Text string `xml:",chardata"`
			} `xml:"hru"`
		} `xml:"content_category"`
	} `xml:"content_categories"`
	ContentCategory2s struct {
		Text            string `xml:",chardata"`
		ContentCategory []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			Hru struct {
				Text string `xml:",chardata"`
			} `xml:"hru"`
		} `xml:"content_category"`
	} `xml:"content_category_2s"`
	ContentCategory3s struct {
		Text            string `xml:",chardata"`
		ContentCategory []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			Hru struct {
				Text string `xml:",chardata"`
			} `xml:"hru"`
		} `xml:"content_category"`
	} `xml:"content_category_3s"`
	ContentGenres struct {
		Text         string `xml:",chardata"`
		ContentGenre []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			CategoryID struct {
				Text string `xml:",chardata"`
			} `xml:"category_id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			Hru struct {
				Text string `xml:",chardata"`
			} `xml:"hru"`
		} `xml:"content_genre"`
	} `xml:"content_genres"`
	Persons struct {
		Text   string `xml:",chardata"`
		Person []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Posters struct {
				Text string `xml:",chardata"`
				File []struct {
					Text string `xml:",chardata"`
					URL  struct {
						Text string `xml:",chardata"`
					} `xml:"url"`
					ID struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					Format struct {
						Text string `xml:",chardata"`
					} `xml:"format"`
				} `xml:"file"`
			} `xml:"posters"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
			IsStar struct {
				Text string `xml:",chardata"`
			} `xml:"is_star"`
			EngTitle struct {
				Text string `xml:",chardata"`
			} `xml:"eng_title"`
			GenTitle struct {
				Text string `xml:",chardata"`
			} `xml:"gen_title"`
			Hru struct {
				Text string `xml:",chardata"`
			} `xml:"hru"`
			FakeHru struct {
				Text string `xml:",chardata"`
			} `xml:"fake_hru"`
			BlogPost struct {
				Text string `xml:",chardata"`
			} `xml:"blog_post"`
			BlogLink struct {
				Text string `xml:",chardata"`
			} `xml:"blog_link"`
			BlogAdded struct {
				Text string `xml:",chardata"`
			} `xml:"blog_added"`
			Bio struct {
				Text string `xml:",chardata"`
			} `xml:"bio"`
			Facts struct {
				Text string `xml:",chardata"`
			} `xml:"facts"`
			Rss struct {
				Text string `xml:",chardata"`
			} `xml:"rss"`
			KinopoiskID struct {
				Text string `xml:",chardata"`
			} `xml:"kinopoisk_id"`
		} `xml:"person"`
	} `xml:"persons"`
	PersonsType struct {
		Text       string `xml:",chardata"`
		PersonType []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
		} `xml:"person_type"`
	} `xml:"persons_type"`
	Awards struct {
		Text  string `xml:",chardata"`
		Award []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			Place struct {
				Text string `xml:",chardata"`
			} `xml:"place"`
			URLToResource struct {
				Text string `xml:",chardata"`
			} `xml:"url_to_resource"`
			ImageID struct {
				Text string `xml:",chardata"`
			} `xml:"image_id"`
			StartYear struct {
				Text string `xml:",chardata"`
			} `xml:"start_year"`
			EndYear struct {
				Text string `xml:",chardata"`
			} `xml:"end_year"`
		} `xml:"award"`
	} `xml:"awards"`
	Compilations struct {
		Text        string `xml:",chardata"`
		Compilation []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			Restrict struct {
				Text string `xml:",chardata"`
			} `xml:"restrict"`
			OrigCountry struct {
				Text string `xml:",chardata"`
			} `xml:"orig_country"`
			KinopoiskID struct {
				Text string `xml:",chardata"`
			} `xml:"kinopoisk_id"`
			ImdbRating struct {
				Text string `xml:",chardata"`
			} `xml:"imdb_rating"`
			KinopoiskRating struct {
				Text string `xml:",chardata"`
			} `xml:"kinopoisk_rating"`
			UsaBoxOffice struct {
				Text string `xml:",chardata"`
			} `xml:"usa_box_office"`
			WorldBoxOffice struct {
				Text string `xml:",chardata"`
			} `xml:"world_box_office"`
			WeekendBoxOffice struct {
				Text string `xml:",chardata"`
			} `xml:"weekend_box_office"`
			ProductionBudget struct {
				Text string `xml:",chardata"`
			} `xml:"production_budget"`
			Hru struct {
				Text string `xml:",chardata"`
			} `xml:"hru"`
			FirstSeriesDate struct {
				Text string `xml:",chardata"`
			} `xml:"first_series_date"`
			LastSeriesDate struct {
				Text string `xml:",chardata"`
			} `xml:"last_series_date"`
			Persons struct {
				Text   string `xml:",chardata"`
				Person []struct {
					Text     string `xml:",chardata"`
					PersonID struct {
						Text string `xml:",chardata"`
					} `xml:"person_id"`
					PersonTypeID struct {
						Text string `xml:",chardata"`
					} `xml:"person_type_id"`
					ExtraInfo struct {
						Text string `xml:",chardata"`
					} `xml:"extra_info"`
				} `xml:"person"`
			} `xml:"persons"`
			Genres struct {
				Text  string `xml:",chardata"`
				Genre []struct {
					Text           string `xml:",chardata"`
					ContentGenreID struct {
						Text string `xml:",chardata"`
					} `xml:"content_genre_id"`
				} `xml:"genre"`
			} `xml:"genres"`
			Awards struct {
				Text  string `xml:",chardata"`
				Award []struct {
					Text          string `xml:",chardata"`
					CompilationID struct {
						Text string `xml:",chardata"`
					} `xml:"compilation_id"`
					AwardID struct {
						Text string `xml:",chardata"`
					} `xml:"award_id"`
					Year struct {
						Text string `xml:",chardata"`
					} `xml:"year"`
				} `xml:"award"`
			} `xml:"awards"`
			Posters struct {
				Text   string `xml:",chardata"`
				Poster struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					CompilationID struct {
						Text string `xml:",chardata"`
					} `xml:"compilation_id"`
					Files struct {
						Text string `xml:",chardata"`
						File struct {
							Text string `xml:",chardata"`
							URL  struct {
								Text string `xml:",chardata"`
							} `xml:"url"`
							ID struct {
								Text string `xml:",chardata"`
							} `xml:"id"`
							Format struct {
								Text string `xml:",chardata"`
							} `xml:"format"`
						} `xml:"file"`
					} `xml:"files"`
				} `xml:"poster"`
			} `xml:"posters"`
			SeasonDescription struct {
				Text              string `xml:",chardata"`
				SeasonDescription []struct {
					Text   string `xml:",chardata"`
					Season struct {
						Text string `xml:",chardata"`
					} `xml:"season"`
					Description struct {
						Text string `xml:",chardata"`
					} `xml:"description"`
					AdditionalData struct {
						Text                            string `xml:",chardata"`
						CompilationSeasonAdditionalData struct {
							Text string `xml:",chardata"`
							ID   struct {
								Text string `xml:",chardata"`
							} `xml:"id"`
							Title struct {
								Text string `xml:",chardata"`
							} `xml:"title"`
							AdditionalDataTypeID struct {
								Text string `xml:",chardata"`
							} `xml:"additional_data_type_id"`
							CompilationSeasonDescriptionID struct {
								Text string `xml:",chardata"`
							} `xml:"compilation_season_description_id"`
							Duration struct {
								Text string `xml:",chardata"`
							} `xml:"duration"`
						} `xml:"compilation_season_additional_data"`
					} `xml:"additional_data"`
				} `xml:"season_description"`
			} `xml:"season_description"`
			ReleaseDates struct {
				Text        string `xml:",chardata"`
				ReleaseDate []struct {
					Text             string `xml:",chardata"`
					ReleaseDateBegin struct {
						Text string `xml:",chardata"`
					} `xml:"release_date_begin"`
					ReleaseDateEnd struct {
						Text string `xml:",chardata"`
					} `xml:"release_date_end"`
				} `xml:"release_date"`
			} `xml:"release_dates"`
			AdditionalData struct {
				Text                      string `xml:",chardata"`
				CompilationAdditionalData []struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					Title struct {
						Text string `xml:",chardata"`
					} `xml:"title"`
					AdditionalDataTypeID struct {
						Text string `xml:",chardata"`
					} `xml:"additional_data_type_id"`
					CompilationID struct {
						Text string `xml:",chardata"`
					} `xml:"compilation_id"`
					Duration struct {
						Text string `xml:",chardata"`
					} `xml:"duration"`
				} `xml:"compilation_additional_data"`
			} `xml:"additional_data"`
			Categories struct {
				Text     string `xml:",chardata"`
				Category []struct {
					Text              string `xml:",chardata"`
					ContentCategoryID struct {
						Text string `xml:",chardata"`
					} `xml:"content_category_id"`
				} `xml:"category"`
			} `xml:"categories"`
			Tima struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"tima"`
			ProductionCompany struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"production_company"`
			Theme struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"theme"`
			Audience struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"audience"`
			Mood struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"mood"`
			Quality struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"quality"`
			About struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"about"`
			Basis struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"basis"`
			Place struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"place"`
			Fake struct {
				Text string `xml:",chardata"`
			} `xml:"fake"`
			Subgenre struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"subgenre"`
			Gender struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"gender"`
			Character struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"character"`
			CategoryProperty struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"category_property"`
			Check struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"check"`
			Restrictions struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"restrictions"`
		} `xml:"compilation"`
	} `xml:"compilations"`
	Countries struct {
		Text    string `xml:",chardata"`
		Country []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
		} `xml:"country"`
	} `xml:"countries"`
	Instruments struct {
		Text       string `xml:",chardata"`
		Instrument []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Info struct {
				Text string `xml:",chardata"`
			} `xml:"info"`
		} `xml:"instrument"`
	} `xml:"instruments"`
	Moods struct {
		Text string `xml:",chardata"`
		Mood []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Info struct {
				Text string `xml:",chardata"`
			} `xml:"info"`
		} `xml:"mood"`
	} `xml:"moods"`
	RusVersions struct {
		Text       string `xml:",chardata"`
		RusVersion []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
		} `xml:"rus_version"`
	} `xml:"rus_versions"`
	ContentFormats struct {
		Text          string `xml:",chardata"`
		ContentFormat []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
			TypeID struct {
				Text string `xml:",chardata"`
			} `xml:"type_id"`
			Group struct {
				Text string `xml:",chardata"`
			} `xml:"group"`
		} `xml:"content_format"`
	} `xml:"content_formats"`
	Langs struct {
		Text string `xml:",chardata"`
		Lang []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"name"`
		} `xml:"lang"`
	} `xml:"langs"`
	ContentTypes struct {
		Text        string `xml:",chardata"`
		ContentType []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			Description struct {
				Text string `xml:",chardata"`
			} `xml:"description"`
		} `xml:"content_type"`
	} `xml:"content_types"`
	Properties struct {
		Text     string `xml:",chardata"`
		Property []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			XMLField struct {
				Text string `xml:",chardata"`
			} `xml:"xml_field"`
			Unique struct {
				Text string `xml:",chardata"`
			} `xml:"unique"`
		} `xml:"property"`
	} `xml:"properties"`
	PropertyValues struct {
		Text          string `xml:",chardata"`
		PropertyValue []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
		} `xml:"property_value"`
	} `xml:"property_values"`
	Promo struct {
		Text  string `xml:",chardata"`
		Promo []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
		} `xml:"promo"`
	} `xml:"promo"`
	AdditionalDataType struct {
		Text                      string `xml:",chardata"`
		ContentAdditionalDataType []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
		} `xml:"content_additional_data_type"`
	} `xml:"additional_data_type"`
	Contents struct {
		Text    string `xml:",chardata"`
		Content []struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"id"`
			Title struct {
				Text string `xml:",chardata"`
			} `xml:"title"`
			OrigTitle struct {
				Text string `xml:",chardata"`
			} `xml:"orig_title"`
			Compilation struct {
				Text string `xml:",chardata"`
			} `xml:"compilation"`
			CopyrightOverlay struct {
				Text string `xml:",chardata"`
			} `xml:"copyright_overlay"`
			Season struct {
				Text string `xml:",chardata"`
			} `xml:"season"`
			Episode struct {
				Text string `xml:",chardata"`
			} `xml:"episode"`
			Duration struct {
				Text string `xml:",chardata"`
			} `xml:"duration"`
			PreviewID struct {
				Text string `xml:",chardata"`
			} `xml:"preview_id"`
			Restrict struct {
				Text string `xml:",chardata"`
			} `xml:"restrict"`
			PrSiteID struct {
				Text string `xml:",chardata"`
			} `xml:"pr_site_id"`
			OriginCountry struct {
				Text string `xml:",chardata"`
			} `xml:"origin_country"`
			ReleaseDate struct {
				Text string `xml:",chardata"`
			} `xml:"release_date"`
			KinopoiskID struct {
				Text string `xml:",chardata"`
			} `xml:"kinopoisk_id"`
			CompilationID struct {
				Text string `xml:",chardata"`
			} `xml:"compilation_id"`
			PromoID struct {
				Text string `xml:",chardata"`
			} `xml:"promo_id"`
			ImdbRating struct {
				Text string `xml:",chardata"`
			} `xml:"imdb_rating"`
			KinopoiskRating struct {
				Text string `xml:",chardata"`
			} `xml:"kinopoisk_rating"`
			UsaBoxOffice struct {
				Text string `xml:",chardata"`
			} `xml:"usa_box_office"`
			WorldBoxOffice struct {
				Text string `xml:",chardata"`
			} `xml:"world_box_office"`
			WeekendBoxOffice struct {
				Text string `xml:",chardata"`
			} `xml:"weekend_box_office"`
			ProductionBudget struct {
				Text string `xml:",chardata"`
			} `xml:"production_budget"`
			YoutubeTrailerURL struct {
				Text string `xml:",chardata"`
			} `xml:"youtube_trailer_url"`
			DateInsert struct {
				Text string `xml:",chardata"`
			} `xml:"date_insert"`
			DescriptionAlt struct {
				Text string `xml:",chardata"`
			} `xml:"description_alt"`
			SharingDisabled struct {
				Text string `xml:",chardata"`
			} `xml:"sharing_disabled"`
			RightDateEnd struct {
				Text string `xml:",chardata"`
			} `xml:"right_date_end"`
			DrmOnly struct {
				Text string `xml:",chardata"`
			} `xml:"drm_only"`
			Instruments struct {
				Text string `xml:",chardata"`
			} `xml:"instruments"`
			ContentGenres struct {
				Text         string `xml:",chardata"`
				ContentGenre []struct {
					Text      string `xml:",chardata"`
					ContentID struct {
						Text string `xml:",chardata"`
					} `xml:"content_id"`
					ContentGenreID struct {
						Text string `xml:",chardata"`
					} `xml:"content_genre_id"`
					Priority struct {
						Text string `xml:",chardata"`
					} `xml:"priority"`
				} `xml:"content_genre"`
			} `xml:"content_genres"`
			Persons struct {
				Text   string `xml:",chardata"`
				Person []struct {
					Text     string `xml:",chardata"`
					PersonID struct {
						Text string `xml:",chardata"`
					} `xml:"person_id"`
					PersonTypeID struct {
						Text string `xml:",chardata"`
					} `xml:"person_type_id"`
					ExtraInfo struct {
						Text string `xml:",chardata"`
					} `xml:"extra_info"`
				} `xml:"person"`
			} `xml:"persons"`
			Awards struct {
				Text  string `xml:",chardata"`
				Award []struct {
					Text      string `xml:",chardata"`
					ContentID struct {
						Text string `xml:",chardata"`
					} `xml:"content_id"`
					AwardID struct {
						Text string `xml:",chardata"`
					} `xml:"award_id"`
					Year struct {
						Text string `xml:",chardata"`
					} `xml:"year"`
				} `xml:"award"`
			} `xml:"awards"`
			Moods struct {
				Text string `xml:",chardata"`
				Mood []struct {
					Text string `xml:",chardata"`
				} `xml:"mood"`
			} `xml:"moods"`
			ContentCategories struct {
				Text            string `xml:",chardata"`
				ContentCategory []struct {
					Text string `xml:",chardata"`
				} `xml:"content_category"`
			} `xml:"content_categories"`
			Posters struct {
				Text   string `xml:",chardata"`
				Poster struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					ContentID struct {
						Text string `xml:",chardata"`
					} `xml:"content_id"`
					Files struct {
						Text string `xml:",chardata"`
						File struct {
							Text string `xml:",chardata"`
							URL  struct {
								Text string `xml:",chardata"`
							} `xml:"url"`
							ID struct {
								Text string `xml:",chardata"`
							} `xml:"id"`
							Format struct {
								Text string `xml:",chardata"`
							} `xml:"format"`
						} `xml:"file"`
					} `xml:"files"`
				} `xml:"poster"`
			} `xml:"posters"`
			Trailers struct {
				Text string `xml:",chardata"`
			} `xml:"trailers"`
			RusVersions struct {
				Text       string `xml:",chardata"`
				RusVersion []struct {
					Text string `xml:",chardata"`
				} `xml:"rus_version"`
			} `xml:"rus_versions"`
			AdditionalData struct {
				Text                  string `xml:",chardata"`
				ContentAdditionalData []struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					Title struct {
						Text string `xml:",chardata"`
					} `xml:"title"`
					AdditionalDataTypeID struct {
						Text string `xml:",chardata"`
					} `xml:"additional_data_type_id"`
					ContentID struct {
						Text string `xml:",chardata"`
					} `xml:"content_id"`
					Duration struct {
						Text string `xml:",chardata"`
					} `xml:"duration"`
				} `xml:"content_additional_data"`
			} `xml:"additional_data"`
			Thumbs struct {
				Text  string `xml:",chardata"`
				Thumb struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					ContentID struct {
						Text string `xml:",chardata"`
					} `xml:"content_id"`
					Files struct {
						Text string `xml:",chardata"`
						File []struct {
							Text string `xml:",chardata"`
							URL  struct {
								Text string `xml:",chardata"`
							} `xml:"url"`
							ID struct {
								Text string `xml:",chardata"`
							} `xml:"id"`
							Format struct {
								Text string `xml:",chardata"`
							} `xml:"format"`
						} `xml:"file"`
					} `xml:"files"`
				} `xml:"thumb"`
			} `xml:"thumbs"`
			Subscriptions struct {
				Text         string `xml:",chardata"`
				Subscription []struct {
					Text string `xml:",chardata"`
				} `xml:"subscription"`
			} `xml:"subscriptions"`
			RightTypeList struct {
				Text      string `xml:",chardata"`
				RightType struct {
					Text string `xml:",chardata"`
				} `xml:"right_type"`
			} `xml:"right_type_list"`
			RightType struct {
				Text string `xml:",chardata"`
			} `xml:"right_type"`
			RightHolders struct {
				Text        string `xml:",chardata"`
				RightHolder []struct {
					Text string `xml:",chardata"`
					ID   struct {
						Text string `xml:",chardata"`
					} `xml:"id"`
					Title struct {
						Text string `xml:",chardata"`
					} `xml:"title"`
				} `xml:"right_holder"`
			} `xml:"right_holders"`
			Mp4 struct {
				Text string `xml:",chardata"`
			} `xml:"mp4"`
			Tima struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"tima"`
			Theme struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"theme"`
			Audience struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"audience"`
			Subgenre struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"subgenre"`
			Mood struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"mood"`
			Quality struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"quality"`
			About struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"about"`
			Basis struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"basis"`
			CategoryProperty struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"category_property"`
			Check struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"check"`
			Place struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"place"`
			Restrictions struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"restrictions"`
			TestProperty struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"test_property"`
			LocalizationCheck struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"localization_check"`
			Gender struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"gender"`
			ProductionCompany struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"production_company"`
			Character struct {
				Text string `xml:",chardata"`
				Item []struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"character"`
			Franchise struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"franchise"`
			Shild struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"shild"`
			NewContent struct {
				Text string `xml:",chardata"`
				Item struct {
					Text string `xml:",chardata"`
				} `xml:"item"`
			} `xml:"new_content"`
		} `xml:"content"`
	} `xml:"contents"`
	ActionLogSequenceNumber struct {
		Text string `xml:",chardata"`
	} `xml:"action_log_sequence_number"`
}
