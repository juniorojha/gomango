# gomango
A utility to use on top of mgo to ease the process of making Mongo queries in Go

## Getting Started

Clone this repository into your project folder residing in the src folder of your GOPATH. See [Writing Go Code](https://golang.org/doc/code.html) for more information.

Update [queries.go](queries.go) with your MongDB details and credentials.

Import the package gomango into the code wherever required like this 

```
import <project package>/gomango
```

Inside your main project package's init function, add the following initialization command -

```
gomango.DatabaseInit()
```

That's it.

### Prerequisites

You need mgo or a related mongodb driver for Go installed. You can install mgo with the following command

```
go get gopkg.in/mgo.v2
```
If you have mgo installed, you can use this out of the box.
If you want to use this with a driver other than mgo, you will need to update the MgoSession variable with the relevant value as mentioned in that driver's documentation.

```
MgoSession      *mgo.Session
```

### Usage example

GoMango's main purpose is to simplify the purpose of making MongoDB queries. 

For e.g.,
Instead of 

```
results := *mgo.Session.DB(testdb).C(testcollection).Find(bson.M{"isActive": true}).Sort("title").All()
```

with GoMango, you can do 

```
results := gomango.GetSortedResults("testcollection", bson.M{"isActive": true}, "title")
```

which will return the same results, i.e. all objects with field 'isActive' being true, sorted by their 'title' fields from 'testcollection'

## List of functions with examples

* **GetResults** - GetResults function will return all results for a given query.

```
results := gomango.GetResults("testcollection", bson.M{"field1": "value1", "field2": "value2"})
```

* **GetSortedResults** - GetSortedResults function will return results sorted as per sorting criteria.
```
results := gomango.GetSortedResults("testcollection", bson.M{"field1": true, "field2": "value2"}, "field3")
```

* **GetFields** - GetFields function uses a projection to show only those fields that are required. Returns all matching fields.
```
results := gomango.GetFields("testcollection", bson.M{"field1": "value1"}, bson.M{"field_to_be_included_1": 1, "field_to_be_included_2": 1, "field_to_be_excluded": 0})
```

* **GetMappedFields** - GetMappedFields is similar to GetResults but returns an array of maps instead of an array of interfaces.
```
results := gomango.GetMappedFields("testcollection", bson.M{"field1": "value1"}, "field_to_be_included": 1})
```

* **GetMappedFieldsWithLimit** - GetMappedFieldsWithLimit is similar to GetMappedFields but uses Limit and Iter to return only the number of results that are required.
```
results = gomango.GetMappedFieldsWithLimit("testcollection", bson.M{"field1": "value1", "field2": "value2", "field3_bool": true, "field4_should_not_exist": bson.M{"$exists": false}, "field5_in_given_array": bson.M{"$in": []string{"arr_val1", "arr_val2", "arr_val3"}}}, bson.M{"field_to_be_included": 1}, 5)

```



## Miscellaneous

I've included comments for every function and every major variable/constant used in the file. You are free to modify the existing functions to create your own functions

## Contributing

Feel free to submit a pull request or fork this repository as per your preferences in order to improve on this utility.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Many thanks to [Abhishek](https://github.com/soniabhishek) for helping me learn Go and showing me the first examples of functions in *to-be* GoMango. He's a great guy!
* Many many thanks to [Saddy](https://github.com/Sadhanandh) for inspiring me to start this project.
