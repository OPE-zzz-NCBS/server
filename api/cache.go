package api

import (
	"database/sql"
	"log"
	"net/http"
	"io/ioutil"
	_ "github.com/mattn/go-sqlite3"
	"github.com/OPENCBS/server/repo"
)

func GetCache(w http.ResponseWriter, r *APIRequest) {
	tempFile, err := ioutil.TempFile("./tmp", "")
	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	fileName := tempFile.Name()

	log.Println("start")
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	err = createCache(db)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	err = initCache(db, r)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	err = setPragma(db)
	if err != nil {
		sendInternalServerError(w, err)
		return
	}

	log.Println("end")

	http.ServeFile(w, r.Request, fileName)
}

func createCache(db *sql.DB) error {
	query := "create table economic_activities (" +
		"_id integer primary key, " +
		"name text not null, " +
		"parent_id integer not null)"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table branches (" +
		"_id integer primary key, " +
		"name text not null, " +
		"code text not null, " +
		"description text not null, " +
		"address text not null)"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table cities (" +
		"_id integer primary key, " +
		"name text not null, " +
		"district_id integer not null)"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table districts (" +
		"_id integer primary key, " +
		"name text not null, " + 
		"region_id integer not null)"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table regions (" +
		"_id integer primary key, " +
		"name text not null)"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table custom_fields (" +
		"_id integer primary key, " +
		"caption text not null, " +
		"type text not null, " +
		"owner text not null, " +
		"tab text not null, " +
		"is_unique integer not null, " +
		"required integer not null, " +
		"field_order integer not null, " +
		"extra text not null)"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table people (" +
		"id integer not null, " +
		"uuid text not null, " +
		"first_name text not null, " +
		"last_name text not null, " +
		"father_name text not null, " +
		"sex text not null, " +
		"birth_date text not null, " +
		"birth_place text not null, " +
		"identification_data text not null, " +
		"activity_id integer not null, " +
		"branch_id integer not null, " +
		"personal_phone text not null, " +
		"home_phone text not null, " +
		"email text not null, " +
		"city_1_id integer not null, " +
		"address_1 text not null, " +
		"postal_code_1 text not null, " +
		"city_2_id integer not null, " +
		"address_2 text not null, " +
		"postal_code_2 text not null)"
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	query = "create table custom_field_values (" +
		"field_id integer not null, " +
		"owner_id text not null, " +
		"value text not null)"
	_, err = db.Exec(query)
	if err != nil {
		return nil
	}

	return nil
}

func initCache(db *sql.DB, r *APIRequest) error {
	tx, err := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	if err = addEconomicActivties(tx, r); err != nil {
		return err
	}

	if err = addBranches(tx, r); err != nil {
		return err
	}

	if err = addCities(tx, r); err != nil {
		return err
	}

	if err = addDistricts(tx, r); err != nil {
		return err
	}

	if err = addCustomFields(tx, r); err != nil {
		return err
	}

	if err = addPeople(tx, r); err != nil {
		return err
	}

	return nil
}

func addEconomicActivties(tx *sql.Tx, r *APIRequest) error {
	repo := repo.NewEconomicActivityRepo(r.DbProvider)
	economicActivities, err := repo.GetAll()
	if err != nil {
		return err
	}

	query := "insert into economic_activities (_id, name, parent_id) values (?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, ea := range economicActivities {
		_, err = stmt.Exec(ea.Id, ea.Name, ea.ParentId)
		if err != nil {
			return err
		}
	}

	return nil
}

func addBranches(tx *sql.Tx, r *APIRequest) error {
	repo := repo.NewBranchRepo(r.DbProvider)
	branches, err := repo.GetAll()
	if err != nil {
		return err
	}

	query := "insert into branches (_id, name, code, description, address) values (?, ?, ?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, branch := range branches {
		_, err = stmt.Exec(branch.Id, branch.Name, branch.Code, branch.Description, branch.Address)
		if err != nil {
			return err
		}
	}

	return nil
}

func addCities(tx *sql.Tx, r *APIRequest) error {
	repo := repo.NewCityRepo(r.DbProvider)
	cities, err := repo.GetAll()
	if err != nil {
		return err
	}

	query := "insert into cities (_id, name, district_id) values (?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, city := range cities {
		_, err := stmt.Exec(city.Id, city.Name, city.DistrictId)
		if err != nil {
			return err
		}
	}

	return nil
}

func addDistricts(tx *sql.Tx, r *APIRequest) error {
	repo := repo.NewDistrictRepo(r.DbProvider)
	districts, err := repo.GetAll()
	if err != nil {
		return err
	}

	query := "insert into districts (_id, name, region_id) values (?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, district := range districts {
		_, err = stmt.Exec(district.Id, district.Name, district.RegionId)
		if err != nil {
			return err
		}
	}

	return nil
}

func addCustomFields(tx *sql.Tx, r *APIRequest) error {
	repo := repo.NewCustomFieldRepo(r.DbProvider)
	fields, err := repo.GetAll()
	if err != nil {
		return err
	}

	query := "insert into custom_fields (" +
		"_id, caption, type, owner, tab, is_unique, required, field_order, extra) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, field := range fields {
		var unique int
		var mandatory int
		if field.Unique {
			unique = 1
		} else {
			unique = 0
		}
		if field.Mandatory {
			mandatory = 1
		} else {
			unique = 0
		}
		_, err := stmt.Exec(
			field.Id,
			field.Caption,
			field.Type,
			field.Owner,
			field.Tab,
			unique,
			mandatory,
			field.Order,
			field.Extra,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func addPeople(tx *sql.Tx, r *APIRequest) error {
	repo := repo.NewPersonRepo(r.DbProvider)
	people, err := repo.GetPeople(0, 100000)
	if err != nil {
		return err
	}

	query := "insert into people (" +
		"id, uuid, first_name, last_name, father_name, sex, birth_date, birth_place, " +
		"identification_data, activity_id, branch_id, personal_phone, home_phone, " +
		"email, city_1_id, address_1, postal_code_1, city_2_id, address_2, postal_code_2) " +
		"values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	query2 := "insert into custom_field_values (field_id, owner_id, value) values (?, ?, ?)"
	stmt2, err := tx.Prepare(query2)
	if err != nil {
		return err
	}
	defer stmt2.Close()

	for _, person := range people {
		_, err = stmt.Exec(
			person.Id,
			person.UUID,
			person.FirstName,
			person.LastName,
			person.FatherName,
			person.Sex,
			person.BirthDate,
			person.BirthPlace,
			person.IdentificationData,
			person.ActivityId,
			person.BranchId,
			person.PersonalPhone,
			person.HomePhone,
			person.Email,
			person.Address1.CityId,
			person.Address1.Address,
			person.Address1.PostalCode,
			person.Address2.CityId,
			person.Address2.Address,
			person.Address2.PostalCode,
		)
		if err != nil {
			return err
		}

		for _, customValue := range person.CustomInformation {
			_, err = stmt2.Exec(customValue.Field.Id, person.UUID, customValue.Value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func setPragma(db *sql.DB) error {
	_, err := db.Exec("pragma user_version = 1")
	if err != nil {
		return err
	}
	return nil
}