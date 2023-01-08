## Golang && React Fullstack application

![Muscles](images/portfolio.png)


## Building application
__Steps__
1. Create a mongo database locally or remotely.
2. Ensure that docker is installed. 
  Run mongodb from docker using `make up`
  Drop the docker container with `make down`
3. Set up the environmental variables for application and database
- `MONGO_URL` - url to mongodb eg ``mongodb://root:root@localhost:27017``
- `MONGO_TIMEOUT` - Max timeout for db
- `MONGO_DB` - database to be used in application
- `PORT` - server port. Default is 4000
4. Ensure that `make` is installed in your system
5. Run `make ui`
6. Run `make run`
7. Go to your browser at `http://localhost:PORT`

