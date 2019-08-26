## Welcome to docker-test!
This is just a repo to help me learn golang, docker, and docker-compose

Downloads required: clone the repo and docker (that's it!)

## Workflow
- docker-compose up (on the first time)
- docker-compose down to stop running
- docker-compose build after making changes
- run tests: docker-compose -f docker-compose.test.yml up

## TODO 
- persist database ✅
- load environment variables securely ✅ ? (not sure if secure, but at least we're using config files)
- make it so you don't have to download go packages every time
- get pgadmin working ✅
- ship!