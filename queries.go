package gomango

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	MongoDBHosts = "127.0.0.1:27017" // MongoDBHosts is your MongoDB URL.
	AuthDatabase = "admin"           // The database that contains the authorization credentials for your MongoDB database. Leave this and all further fields blank if you're running MongoDB without authentication.
	AuthUserName = "user"            // The user through which you would like to access your MongoDB database. This user could have read or write access or both. In case it has only read access, some of GoMongo's functions may result in failure.
	AuthPassword = "password"        // The password for the user which has authenticated access to your MongoDB database.
	TestDatabase = "testdb"          // This constant holds the name of a test database, so that you can quickly change from the main DB to a test DB using the Dbname variable below.Aggregation
	// Add any more constants as required.
)

var (
	MgoSession      *mgo.Session
	Dbname          = "targetdb" // Database that contains the actual information that is required to be queried.
	mongoDBDialInfo = &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}
)

// DatabaseInit function initializes the database connection.
func DatabaseInit() {
	var err error
	MgoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Println(err)
	}
}

// GetResults function will return all results for a given query.
func GetResults(collectionName string, query interface{}) (result []interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).All(&result)
	return
}

// GetSortedResults function will return results sorted as per sorting criteria.
func GetSortedResults(collectionName string, query interface{}, sortingcriteria string) (result []interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).Sort(sortingcriteria).All(&result)
	return
}

// GetFields function uses a projection to show only those fields that are required. Returns all matching fields.
func GetFields(collectionName string, query interface{}, selector interface{}) (result []interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).Select(selector).All(&result)
	return
}

// GetMappedFields is similar to GetResults but returns an array of maps instead of an array of interfaces.
func GetMappedFields(collectionName string, query interface{}, selector interface{}) (result []map[string]interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).Select(selector).All(&result)
	return
}

// GetMappedFieldsWithLimit is similar to GetMappedFields but uses Limit and Iter to return only the number of results that are required.
func GetMappedFieldsWithLimit(collectionName string, query interface{}, selector interface{}, limit int) (result []map[string]interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).Select(selector).All(&result)
	iter := MgoSession.DB(Dbname).C(collectionName).Find(query).Select(selector).Limit(limit).Iter()
	err := iter.All(&result)
	if err != nil {
		log.Println(err)
	}
	return
}

// GetSortedResultsMap function will return results sorted as per sorting criteria in the form of an array of maps.
func GetSortedResultsMap(collectionName string, query interface{}, sortingcriteria string) (result []map[string]interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).Sort(sortingcriteria).All(&result)
	return
}

// GetSortedMappedFields function will return only the requested fields sorted as per sorting criteria in the form of an array of maps.
func GetSortedMappedFields(collectionName string, query interface{}, selector interface{}, sortingcriteria string) (result []map[string]interface{}) {
	MgoSession.DB(Dbname).C(collectionName).Find(query).Select(selector).Sort(sortingcriteria).All(&result)
	return
}

// DeleteFields is actually an Update operation but can delete whole documents as well.
func DeleteFields(collectionName string, selector interface{}, updater interface{}) bool {
	err := MgoSession.DB(Dbname).C(collectionName).Update(selector, updater)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// RemoveDocument to remove all documents matching given conditions. Avoid using DeleteFields for this purpose.
func RemoveDocument(collectionName string, selector interface{}) (err error) {
	return MgoSession.DB(Dbname).C(collectionName).Remove(selector)
}

// GetResultsById is a variant of GetResults which returns one result - matching a given set of criteria, which is supposed to be an ID. This function could be modified in later versions to match only ID as that's a frequent requirement in MongoDB-backed data.
func GetResultsById(collectionName string, query interface{}) (result interface{}) {
	err := MgoSession.DB(Dbname).C(collectionName).Find(query).One(&result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// InsertDocument function accepts any number of queries to insert to a document. Returns success or failure as boolean.
func InsertDocument(collectionName string, query ...interface{}) bool {
	err := MgoSession.DB(Dbname).C(collectionName).Insert(query...)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println(err)
	return true
}

// UpsertCollection is a standard Upsert operation. Returns changeInfo and error.
func UpsertCollection(collectionName string, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return MgoSession.DB(Dbname).C(collectionName).Upsert(selector, update)
}

// FindOneDocument is a lot like GetResultsById but the return type here is map[string]interface{} instead of just an interface{}.
func FindOneDocument(collectionName string, query interface{}) (result map[string]interface{}) {
	err := MgoSession.DB(Dbname).C(collectionName).Find(query).One(&result)
	if err != nil {
		log.Println(err)
	}
	return
}

// FindDocuments is an 'All' variant of FindOneDocument. Returns all matching results as an array of maps.
func FindDocuments(collectionName string, query interface{}) (result []map[string]interface{}) {
	err := MgoSession.DB(Dbname).C(collectionName).Find(query).All(&result)
	if err != nil {
		log.Println(err)
	}
	return
}

// UpdateDocument will update fields matching the given query. Use $set with this to retain other fields.
func UpdateDocument(collectionName string, selector interface{}, update interface{}) (err error) {
	return MgoSession.DB(Dbname).C(collectionName).Update(selector, update)
}

// GetCount returns the count of results of a given query.
func GetCount(collectionName string, selector interface{}) (int, error) {
	return MgoSession.DB(Dbname).C(collectionName).Find(selector).Count()
}

// Aggregation function provides a medium to use the MongoDB Aggregation Framework.
func Aggregation(collectionName string, aggregationquery []bson.M) (result []map[string]interface{}) {
	iter := MgoSession.DB(Dbname).C(collectionName).Pipe(aggregationquery).Iter()
	err := iter.All(&result)
	if err != nil {
		log.Println(err)
	}
	return
}

// Distinct function to return distinct records for a given field in a query.
func Distinct(collectionName string, query interface{}, field string) (result []interface{}) {
	err := MgoSession.DB(Dbname).C(collectionName).Find(query).Distinct(field, &result)
	if err != nil {
		log.Println(err)
	}
	return
}

// FindAndModify function to find and modify in a single step
func FindAndModify(collectionName string, find interface{}, modify interface{}) (result, info interface{}, err error) {
	change := mgo.Change{
		Update:    modify,
		ReturnNew: true,
	}
	info, err = MgoSession.DB(Dbname).C(collectionName).Find(find).Apply(change, &result)
	return
}
