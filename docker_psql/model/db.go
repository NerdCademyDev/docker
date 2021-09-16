package model

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver, use _ so we don't have to init anything
)

// package level variable so we can use in methods here
var db *sql.DB

// connection data type, lowercase cannot access outside of this package
type connection struct {
	Host string
	Port string
	User string
	Password string
	DBName string
}

// Post data type so we can package the data easier
type Post struct {
	ID int
	Title string
	Content string
}

// Init will initialize our db connection
// returns an error if there's an issue connecting or pinging DB
func Init() error {
	// create a connection data type
	conn := connection{
		Host: "localhost",
		Port: "5432",
		User: "dev",
		Password: "secret",
		DBName: "dev",
	}
	// create a string with our connection information
	connInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conn.Host,
		conn.Port,
		conn.User,
		conn.Password,
		conn.DBName,
	)

	var err error

	// try to open our db connection, using the postres driver
	db, err = sql.Open("postgres", connInfo)

	// check for error
	if err != nil {
		return err
	}
	// try to ping our DB
	err = db.Ping()
	// check for error
	if err != nil {
		return err
	}
	// return nil because no error if we got here!
	return nil
}

// GetPost returns the post by ID specified.
// returns an error if there is any issue in the SQL query
func GetPost(postID int) (Post, error) {
	// declare a Post to hold our data
	var post Post

	// create the SQL query statement.  $1 will be replaced by the postID
	sqlQuery := `SELECT "id", "title", "content" FROM "post" where id = $1;`

	// execute the query, should return a row from the table if ID is good
	// returns an error if there is any issues.  $1 is replaced with postID
	row, err := db.Query(sqlQuery, postID)

	// there's an error!
	if err != nil {
		// return the empty post and error
		return post, err
	}

	// declare variables to hold our post data
	var id int
	var title, content string

	// loop through the row
	for row.Next() {
		// scan the current row and pass in by reference our post data variables
		// if no error, it will contain the data from the db.
		// these should be in order as our sqlQuery string from above
		err := row.Scan(&id, &title, &content)

		// check for error
		if err != nil {
			return post, err
		}

		// set our post data to our struct
		post.ID = id
		post.Title = title
		post.Content = content
	}
	// return our filled in post struct and nil (no error if we got here!)
	return post, nil
}

// CreatePost creates a post by inserting the data into the table `post`
func CreatePost(post Post) error {

	sqlQuery := `INSERT into post(title, content) values($1, $2);`

	_, err := db.Exec(sqlQuery, post.Title, post.Content)
	
	return err
}

// Close will attempt to close the connection to the DB
// will return an error if an issue closing
func Close() error {
	return db.Close()
}