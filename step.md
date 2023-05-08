### Install migrate

```bash
brew install golang-migrate
mkdir -p db/migration
migrate create -ext sql -dir db/migration -seq init_schema
```

These will be 2 script up and down <br>
- The up script will update the schema to the new version
- The down script will revert back to the previouse version

When we run `migrate up` the old db will run throught the migrate file in the order sequentially and then update to the new db <br>

When we run `migrate down` the new db will run throught the migrate file in reverse order sequentially and down grade to the old version of db <br>

### Migrate the database

```bash
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```


### CRUD
- CREATE: Insert new records to the database
- READ: Select or search for records in the database
- UPDATE: Change some fields of the record in the databse
- DELETE: Remove records from the database

### Goland and CRUD
- DATABASE/SQL
    - Very fast & straightforward
    - Manual mapping SQL fields to variable
    - Easy to make mistakes, not caught until runtime

- GORM
    - CRUD functions already implemented, very short production code
    - Must learn to write queries using gorm's function
    - Run slowly on high load

- SQLX
    - Quite fast & easy to use
    - Field mapping via query text & struct tags
    - Failure won't occur until runtime

- SQLC
    - Very fast & easy to use
    - Automatic code generation
    - Catch SQL query errors before generating codes
    - Full support Postgres. MySQL is experimental


### Database Transaction
- A single unit of work
- Often made up of multiple db operations

> Example: Transfer 10 USD from bank account 1 to bank accoun 2
> 1. Create a transfer record with amount 10
> 2. Create an account entry for account 1 with amount = -10
> 3. Create an account entry for account 2 with amount = +10
> 4. Substract 10 from the balance of account 1
> 5. Add 10 to the balance of account 2

Why do we need db transaction?
- To provide a reliable and consistent unit of work, even in case of system failure
- To provide isolation between programs that access the database concurrently
- It must sactify the ACID Property
    - Atomicity (A)
        - Either all operations complete successfully or the transaction fails and the db is unchanged
    - Consistency (C)
        - The db state must be valid after the transaction. All contraints must be satisfied
    - Isolation (I)
        - Concurrent transactions must not affect each other.
    - Durability (D)
        - Data written by a successful transaction must be recorded in persistent storage

To run sql transaction
- Commit to commit all change
- Rollback when fail to commit and it will delete all change

Some issue with SQL <br>
Read Phenomena
- Dirty Read
    - A transaction reads data written by other concurrent uncommitted transaction
- Non-Repeatable read
    - A transaction reads the same row twice and sees different value because it has been modified by other commited transaction
- Phantom Read
    - A transaction re-executes a query to find rows that sstisfy a condition and sees a different set of rows, due to changes by other committed transaction
- Serialization Anomaly
    - The result of a group of concurrent commited transactions is impossible to achieve if we try to run them sequentially in any order without overlapping

4 Standard Isolation Levels (American National Standards Institute - ANSI)
1. Read Uncommitted: Can see data written by uncommitted transaction
2. Read Committed: Only see data written by committed transaction
3. Repeatable Read: Same read query always returns same result
4. Serializable: Can achieve same result if execute transactions serially in some order instead of concurrently

> Remember
> - Retry Mechanism
>   - There might be errors, timeout or deadlock
> - Read Documentation
>   - Each database engine might implement isolation level differently