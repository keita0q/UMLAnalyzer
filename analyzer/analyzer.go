package analyzer

import (
	"github.com/keita0q/UMLAnalyzer/model"
	"database/sql"
	"time"
	"fmt"
)

type Analyzer struct {
	peopleDau        *model.DAU
	createCardsDau   *model.DAU
	importCardsDau   *model.DAU
	updateProfileDau *model.DAU
}

type Config struct {
	DB *sql.DB
}

const PEOPLE_QUERY = "SELECT * FROM user_set WHERE name = 'people_dau'"
const CREATE_CARDS_QUERY = "SELECT * FROM user_set WHERE name = 'create_cards_dau'"
const IMPORT_CARDS_QUERY = "SELECT * FROM user_set WHERE name = 'import_cards_dau'"
const UPDATE_PROFILE_QUERY = "SELECT * FROM user_set WHERE name = 'update_profile_dau'"

type AnalyzeResult struct {

}

func New(aConfig *Config) (*Analyzer, error) {
	tPeopleDau, tError := createDau(aConfig.DB, PEOPLE_QUERY)
	if tError != nil {
		return nil, tError
	}
	tCreateCardsDau, tError := createDau(aConfig.DB, CREATE_CARDS_QUERY)
	if tError != nil {
		return nil, tError
	}
	tImportCardsDau, tError := createDau(aConfig.DB, IMPORT_CARDS_QUERY)
	if tError != nil {
		return nil, tError
	}
	tUpdateProfileDau, tError := createDau(aConfig.DB, UPDATE_PROFILE_QUERY)
	if tError != nil {
		return nil, tError
	}
	return &Analyzer{
		peopleDau: model.NewDAU(tPeopleDau),
		createCardsDau: model.NewDAU(tCreateCardsDau),
		importCardsDau: model.NewDAU(tImportCardsDau),
		updateProfileDau: model.NewDAU(tUpdateProfileDau),
	}, nil
}

func createDau(aDB *sql.DB, aQuery string) ([]model.UserSet, error) {
	tDau := []model.UserSet{}
	tRows, tError := aDB.Query(aQuery)
	if tError != nil {
		return nil, tError
	}
	for tRows.Next() {
		tUserSet := model.UserSet{}
		if tError := tRows.Scan(&tUserSet.Name, &tUserSet.UserIDs, &tUserSet.Size, &tUserSet.Date); tError != nil {
			return nil, tError
		}
		tDau = append(tDau, tUserSet)
	}
	return tDau, nil
}

func (aAnalyzer *Analyzer)AnalyzeAllPersistency(BaseDate time.Time, aPeriod int) *AnalyzeResult {
	tResult := &AnalyzeResult{}

	return tResult
}

