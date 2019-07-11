DROP SCHEMA public CASCADE;
CREATE SCHEMA public;
CREATE TABLE content_categories
(
	"id" text,
	"title" text,
	"description" text,
	"hru" text
);
CREATE TABLE moods
(
	"id" text,
	"title" text,
	"info" text
);
CREATE TABLE content_formats
(
	"id" text,
	"title" text,
	"description" text,
	"type_id" text,
	"group" text
);
CREATE TABLE langs
(
	"id" text,
	"title" text,
	"name" text
);
CREATE TABLE promo
(
	"id" text,
	"title" text
);
CREATE TABLE persons
(
	"id" text,
	"posters" jsonb,
	"name" text,
	"is_star" text,
	"eng_title" text,
	"gen_title" text,
	"hru" text,
	"fake_hru" text,
	"blog_post" text,
	"blog_link" text,
	"blog_added" text,
	"bio" text,
	"facts" text,
	"rss" text,
	"kinopoisk_id" text
);
CREATE TABLE countries
(
	"id" text,
	"title" text,
	"name" text
);
CREATE TABLE rus_versions
(
	"id" text,
	"title" text
);
CREATE TABLE property_values
(
	"id" text,
	"title" text
);
CREATE TABLE additional_data_type
(
	"id" text,
	"title" text
);
CREATE TABLE content_genres
(
	"id" text,
	"category_id" text,
	"title" text,
	"description" text,
	"hru" text
);
CREATE TABLE instruments
(
	"id" text,
	"title" text,
	"info" text
);
CREATE TABLE content_types
(
	"id" text,
	"title" text,
	"description" text
);
CREATE TABLE properties
(
	"id" text,
	"title" text,
	"xml_field" text,
	"unique" text
);
CREATE TABLE contents
(
	"id" text,
	"title" text,
	"orig_title" text,
	"compilation" text,
	"copyright_overlay" text,
	"season" text,
	"episode" text,
	"duration" text,
	"preview_id" text,
	"restrict" text,
	"pr_site_id" text,
	"origin_country" text,
	"release_date" text,
	"kinopoisk_id" text,
	"compilation_id" text,
	"promo_id" text,
	"imdb_rating" text,
	"kinopoisk_rating" text,
	"usa_box_office" text,
	"world_box_office" text,
	"weekend_box_office" text,
	"production_budget" text,
	"youtube_trailer_url" text,
	"date_insert" text,
	"description_alt" text,
	"sharing_disabled" text,
	"right_date_end" text,
	"drm_only" text,
	"instruments" text,
	"content_genres" jsonb,
	"persons" jsonb,
	"awards" jsonb,
	"moods" jsonb,
	"content_categories" jsonb,
	"posters" jsonb,
	"trailers" text,
	"rus_versions" jsonb,
	"additional_data" jsonb,
	"thumbs" jsonb,
	"subscriptions" jsonb,
	"right_type_list" jsonb,
	"right_type" text,
	"right_holders" jsonb,
	"mp4" text,
	"tima" jsonb,
	"theme" jsonb,
	"audience" jsonb,
	"subgenre" jsonb,
	"mood" jsonb,
	"quality" jsonb,
	"about" jsonb,
	"basis" jsonb,
	"category_property" jsonb,
	"check" jsonb,
	"place" jsonb,
	"restrictions" jsonb,
	"test_property" jsonb,
	"localization_check" jsonb,
	"gender" jsonb,
	"production_company" jsonb,
	"character" jsonb,
	"franchise" jsonb,
	"shild" jsonb,
	"new_content" jsonb
);
CREATE TABLE content_category_2s
(
	"id" text,
	"title" text,
	"description" text,
	"hru" text
);
CREATE TABLE content_category_3s
(
	"id" text,
	"title" text,
	"description" text,
	"hru" text
);
CREATE TABLE persons_type
(
	"id" text,
	"title" text,
	"description" text
);
CREATE TABLE awards
(
	"id" text,
	"title" text,
	"description" text,
	"place" text,
	"url_to_resource" text,
	"image_id" text,
	"start_year" text,
	"end_year" text
);
CREATE TABLE compilations
(
	"id" text,
	"title" text,
	"description" text,
	"restrict" text,
	"orig_country" text,
	"kinopoisk_id" text,
	"imdb_rating" text,
	"kinopoisk_rating" text,
	"usa_box_office" text,
	"world_box_office" text,
	"weekend_box_office" text,
	"production_budget" text,
	"hru" text,
	"first_series_date" text,
	"last_series_date" text,
	"persons" jsonb,
	"genres" text[],
	"awards" jsonb,
	"posters" jsonb,
	"season_description" jsonb,
	"release_dates" jsonb,
	"additional_data" jsonb,
	"categories" text[],
	"tima" jsonb,
	"production_company" jsonb,
	"theme" jsonb,
	"audience" jsonb,
	"mood" jsonb,
	"quality" jsonb,
	"about" jsonb,
	"basis" jsonb,
	"place" jsonb,
	"fake" text,
	"subgenre" jsonb,
	"gender" jsonb,
	"character" jsonb,
	"category_property" jsonb,
	"check" jsonb,
	"restrictions" jsonb
);
