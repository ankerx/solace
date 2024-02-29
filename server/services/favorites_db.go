package services

import (
	"task/system"
)

type Favorite struct {
	ID      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Mal_id  int    `json:"mal_id"`
	// ...
}

func dest(fav *Favorite) []interface{} {
	return []interface{}{
		&fav.ID,
		&fav.Created,
		&fav.Updated,
		&fav.Title,
		&fav.Image,
		&fav.Mal_id,
		// ...
	}
}

func selectAllFavorites() ([]Favorite, error) {
	rows, err := system.Db.Query("select * from favorites")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favs []Favorite = make([]Favorite, 0)
	for rows.Next() {
		fav := Favorite{}
		err = rows.Scan(dest(&fav)...)
		if err != nil {
			return nil, err
		}
		favs = append(favs, fav)
	}
	return favs, nil
}
func addToFavorites(mal_id int, title, image string) error {
	query := "INSERT INTO favorites (mal_id, title, image, id) VALUES (?,?,?,?)"
	id := new(int)
	_, err := system.Db.Exec(query, mal_id, title, image, id)
	if err != nil {
		return err
	}
	return nil
}

func removeFromFavorites(mal_id int) error {
	query := "DELETE FROM favorites WHERE mal_id=?"
	_, err := system.Db.Exec(query, mal_id)
	if err != nil {
		return err
	}
	return nil
}
