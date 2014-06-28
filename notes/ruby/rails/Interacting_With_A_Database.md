# Interating With A Database

## Migrations

To generate a migration:

```
rails g migration CreateBooks
```

When the command above is executed in the terminal, it will generate a Rails migration file which will get executed the next time `rake db:migrate` is executed.

```
class CreateBooksTable < ActiveRecord::Migration
    
    def change
        create_table :books do |t|
            t.string :title
            t.string :body
        end
    end

end
```

In a migration file, there are three possible methods that can be executed, which are `up`, `down`, and `change`. `up` is used to make additions to a table, `down` is used to make deletions to a table, and `change` can be used to do both.

Within these methods, there are several other methods that can be used:
* `create_table`

To execute a migration, you can use the `rake db:migate`. This command will run the unprocessed database migrations.

## Seeding

```
rake db:seed
```

Seeding the database is the process of injecting predefined values into your database.


## Associations

Associations are ways of creating a connection between database table, such as saying that a tweet belongs to a certain user or a product belongs in a certain user's shopping cart.

### Types of Associations
Associations between database tables are similar to relationships in real life.You share many relationships with members of your family. For instance, you and your mother are related. You have only one mother, but she may have several children. You and your siblings are related—you may have many brothers and sisters and, of course, they'll have many brothers and sisters as well. If you're married, both you and your spouse have a spouse—each other—but only one at a time.

There are three types of relational database associations:

**One-to-one:** Both tables can have only one record on either side of the relationship. Each primary key value relates to only one (or no) record in the related table. They're like spouses—you may or may not be married, but if you are, both you and your spouse have only one spouse. Most one-to-one relationships are forced by business rules and don't flow naturally from the data. In the absence of such a rule, you can usually combine both tables into one table without breaking any normalization rules.

**One-to-many:** The primary key table contains only one record that relates to none, one, or many records in the related table. This relationship is similar to the one between you and a parent. You have only one mother, but your mother may have several children. This relationship is also referred to **Has Many-Belongs To**.

**Many-to-many:** Each record in both tables can relate to any number of records (or no records) in the other table. For instance, if you have several siblings, so do your siblings (have many siblings). Many-to-many relationships require a third table, known as an associate or linking table, because relational systems can't directly accommodate the relationship.

**Has And Belongs To Many:** With this relation, two or more records can have a relationship with one another through a join table. The join table would be named after the records whose relations are being stored seperated by underscores, such as *Posts_Users*. The way this relation is expressed is saying that one records has many of another record through the join table. So, a user has many posts through the Post_Users table and a post has many users through the Post_Users table. In this case, our post can have many writers and your writers can have many posts that they've written together.

To create a association in a Rails:
```
# /migrations/...
def change
    change_table :posts do |t|
        t.references :user
    end
end

# app/models/post.rb
class Post < ActiveRecord::Base
    belongs_to :user
end

# app/models/user.rb
class User < ActiveRecord::Base
    has_many :posts
end
```

The lines for the migration file will create a column in the table that will hold the id number of the associated model record. So, for a blog post, that column would hold the writer's id number. The line for the model file will create the connection between the two ActiveRecord classes Post and User. With the above code, you would be able to execute `post.user` and the object of the user associated with that post would be returned. You would also be able to execute `user.posts` and an array of all posts associated with that user.

In the posts table, the primary id of the user will be stored in a column names user_id. When the primary id of a record is stored in another table to show association, the id will be referred to as a foreign key because the id is foreign to that table.