Usul - "the strength of the base of the pillar"

Requires:

- Install go
- Install virtualbox
- Install docker toolbox

Create docker machine in virtualbox:

- Run 'docker-machine create --driver=virtualbox usul'

Source the environment variables:

- Run 'eval "$(docker-machine env usul)"'

Build the docker image locally:

- Run 'docker build -t usul .'

Build Usul:

- Run 'go build'

Run Usul:

- Run './usul'

Sample curl queries:

curl -X POST -d "language=python&code=print+%22hello+world%22" http://localhost:8080/run
curl -X POST -d "language=ruby&code=puts+%22hello+world%22" http://localhost:8080/run
curl -X POST -d "language=nodejs&code=console.log(+%22hello+world%22)" http://localhost:8080/run
