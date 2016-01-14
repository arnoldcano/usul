Usul - Docker Code Runner

Supported Languages:

- Ruby
- Python
- Javascript

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

Try web form:

- http://localhost:8080

Try REST api:

- Run `curl -H "Content-Type: application/json" -d '{ "language": "ruby", "code": "puts \"hello+world\"" }' http://localhost:8080/compile`
- Run `curl -H "Content-Type: application/json" -d '{ "language": "python", "code": "print \"hello+world\"" }' http://localhost:8080/compile`
- Run `curl -H "Content-Type: application/json" -d '{ "language": "nodejs", "code": "console.log(\"hello+world\")" }' http://localhost:8080/compile`
