## Welcome to docker-test!
This is just a repo to help me learn golang, docker, postgresql, pgadmin and docker-compose

## Downloads
- Clone repo
- Install docker and docker compose

## First run
- set up config files in ./config
- docker-compose up (that's it!)

## Developer Workflow
- Make a change
  - docker-compose build (after making any change to golang src)
  - docker-compose up
  - docker-compose down 
- Run tests: 
  - docker-compose -f docker-compose.test.yml up
- Inspect db with pgadmin: 
  - Navigate to localhost:5050 and sign in with the credentials specified in ./config
  - If server is not created:
    - Create a server 
      - General > Name = whatever you like (e.g. PG10)
      - Connection > Hostname = "postgres" or whatever your host is named in ./config
      - Connection > Port = "5432" or whatever your port is named in ./config
      - Connection > Username = whatever your email is in ./config
      - Connection > Password = whatever your password is in ./config
      - Save
  - Navigate to Servers > SERVER_NAME > Databases > DATABASE_NAME > Right Click > Query Tool 
    - Enter query such as "SELECT * FROM products"

## TODO 
- persist database ✅
- load environment variables securely ✅ ? (not sure if secure, but at least we're using config files)
- get pgadmin working ✅
- make it so you don't have to download go packages on every build
- ship!
