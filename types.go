package main

import "github.com/graphql-go/graphql"

/*
type Author {
	Name: String
	Tutorials: [Int]
}
*/
var authorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Author",
		Fields: graphql.Fields{
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Tutorials": &graphql.Field{
				// we'll use NewList to deal with an array
				// of int values
				Type: graphql.NewList(graphql.Int),
			},
		},
	},
)

/*
type Comment {
	body:String
}
*/
var commentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		// we define the name and the fields of our
		// object. In this case, we have one solitary
		// field that is of type string
		Fields: graphql.Fields{
			"body": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

/*
type Tutorial {
	id : Int
	title : String
	author : Author
	comments : [Comment]
}
*/
var tutorialType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tutorial",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"title": &graphql.Field{
				Type: graphql.String,
			},
			"author": &graphql.Field{
				// here, we specify type as authorType
				// which we've already defined.
				// This is how we handle nested objects
				Type: authorType,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
			},
		},
	},
)
