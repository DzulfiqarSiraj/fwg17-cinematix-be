package models

import "github.com/putragabrielll/fwg17-cinematix-be/src/services"

func FindMovieCinemaByMovieId(movieId int) (services.MovieCinema, error) {
	sql := `
	SELECT
	"m"."id" AS "movieId",
	array_agg(DISTINCT jsonb_build_object(
	  'cinemaId', "c"."id",
	  'cinemaName', "c"."name" || '-' || "c"."grade",
	  'cinemaPrice', "c"."price",
	  'movieCinemaId', "mc"."id",
	  'cinemaImage', "c"."image"
	)) AS "cinema"
	FROM "cinema" AS "c"
	LEFT JOIN "movieCinema" "mc" ON "mc"."cinemaId"="c"."id"
	LEFT JOIN "movies" "m" ON "m"."id"="mc"."moviesId"
	WHERE "m"."id" = $1
	GROUP BY "m"."id"
	`
	data := services.MovieCinema{}
	err := db.Get(&data, sql, movieId)
	return data, err
}
