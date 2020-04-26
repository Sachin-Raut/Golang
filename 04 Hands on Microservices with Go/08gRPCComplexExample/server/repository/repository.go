/*
documents -> GoWorkspace -> src -> hands-on-microservices -> 08gRPCComplexExample -> 
server -> repository -> repository.go
*/

/* 

//this creates greet.pb.go file, in terminal navigate to proto folder & execute following

$ protoc greet.proto --go_out=plugins=grpc:.

*/

package repository

import (
	"errors"
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq" //pq is
	"hands-on-microservices/08gRPCComplexExample/server/entities"
)

//PsqlRepo is
type PsqlRepo struct {
	db *sql.DB
}

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "a"
	dbname = "wta"
)

//NewWTARepository is
func NewWTARepository() *PsqlRepo {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err.Error())
	}

	repo := &PsqlRepo{}

	repo.db = db
	fmt.Println("success=======")
	return repo
}

//CloseWTARepository is
func (repo *PsqlRepo) CloseWTARepository(){
	repo.db.Close()
}

//GetPlayerWithHighestRanking is
func (repo *PsqlRepo) GetPlayerWithHighestRanking(playerID uint32) (*entities.PlayerWithRanking, error){
	row := repo.db.QueryRow(`select players.player_id, players.first_name, players.last_name, players.isRightHanded, players.birth_date,players.country_code, rankings.ranking_date, rankings.ranking, rankings.ranking_points from players, rankings where players.player_id = rankings.player_id and players.player_id=$1 order by rankings.ranking, rankings.ranking_date limit 1;`,playerID)

	p := entities.PlayerWithRanking{}

	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.IsRightHanded, &p.BirthDate, &p.CountryCode, &p.RankingDate, &p.RankingNumber, &p.RankingPoints)

	if err != nil {
		return nil, errors.New("[GetPlayerWithHighestRanking] error on query to db"+err.Error())
	}
	return &p, nil
}

//GetRankingsByPlayerID is
func (repo *PsqlRepo) GetRankingsByPlayerID(playerID uint32) ([]*entities.Ranking, error){
	var results []*entities.Ranking

	rows, err := repo.db.Query(`select rankings.ranking_date, rankings.ranking, rankings.ranking_points from rankings where rankings.player_id=$1 order by rankings.ranking_date asc;`,playerID)

	if err != nil {
		return nil, errors.New("[GetRankingByPlayerID] error on query to DB "+err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		r := &entities.Ranking{}

		err := rows.Scan(&r.RankingDate, &r.RankingNumber, &r.RankingPoints)

		if err != nil {
			return nil, errors.New("[GetRankingByPlayerID] error on query to db"+err.Error())
		}

		results = append(results, r)
	}

	err = rows.Err()

	if err != nil {
		return nil, errors.New("[GetRankingByPlayerID] error on query to DB"+err.Error())
	}

	return results, nil
}

//GetPlayer is
func (repo *PsqlRepo) GetPlayer(playerID uint32)(*entities.Player, error){
	var p entities.Player

	row := repo.db.QueryRow(`select players.player_id, players.first_name, players.last_name, players.isRightHanded, players.birth_date, 
	players.country_code from players where players.player_id=$1;`,playerID)

	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.IsRightHanded, &p.BirthDate, &p.CountryCode)

	if err != nil {
		return nil, errors.New("[GetPlayers] error on query to db "+err.Error())
	}

	return &p, nil
}
