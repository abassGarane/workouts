## Golang && React Fullstack application

![Muscles](images/portfolio.png)


## Building application
__Steps__
1. Create a mongo database locally or remotely.
2. Set up the environmental variables for application and database
- `MONGO_URL` - url to mongodb eg ``mongodb://root:root@localhost:27017``
- `MONGO_TIMEOUT` - Max timeout for db
- `MONGO_DB` - database to be used in application
- `PORT` - server port. Default is 4000
3. Ensure that `make` is installed in your system
4. Run `make ui`
5. Run `make run`
6. Go to your browser at `http://localhost:PORT`

