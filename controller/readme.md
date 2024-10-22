# Prerequisites
## Environment file
You need to have a `.env` file in the *conf* folder.
By default, the controller looks for `.env.dev`
```
# db variables
DB_HOST=localhost
DB_USER=postgres
DB_PASS=postgres
DB_NAME=test_db
DB_PORT=5432
DB_SSL=disable
DB_TZ=Europe/Helsinki


# go backend variables
GO_BACKEND_HOST=localhost
GO_BACKEND_PORT=3000
GO_BACKEND_API_PATH=/api/v1
HOST_URL=http://localhost:3000 # GO_BACKEND_HOST + GO_BACKEND_PORT
RUNNING_IN_DEV=true
```

## Database
You need postgresql intalled on your device.
```
sudo -u postgres psql
```

Then, create a new database using `DB_NAME` and `DB_PORT` from your .env file
```
create db -p 5432 test_db
```

You need to alter the default user with your `DB_PASS` from your .env file
```
ALTER USER postgres WITH PASSWORD 'postgres';
```

We can now load the sql files from *db/sql* to create the tables and their triggers.
```
sudo -u postgres psql test_db < functions.sql 
sudo -u postgres psql test_db < cron.sql 
sudo -u postgres psql test_db < exec.sql 
```

There will be errors due to extra manipulations with an argument ($x). You can ignore them;

## SuperToken 
To do
