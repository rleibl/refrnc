# example_cities
An example database.

cities.json contains example data for major european cities.
```
[
   {
      "census" : "Date of census",
      "city" : "City",
      "population" : "Official population",
      "country" : "Country"
   },
   {
      "country" : "United Kingdom",
      "city" : "London",
      "population" : "8,825,001",
      "census" : "28 June 2018"
   },
...
]
```

The .sql files cover setup of a postgres database. There are examples for
inserting and selecting data in go.

There are also examples for inserting and retrieving data from a mongodb
database. The mongodb is document oriented and does not need any special setup
like table creation and such.
