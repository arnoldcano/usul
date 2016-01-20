Usul - Sandbox Code Runner

Supported Languages:

- Ruby
- Python
- Javascript

Requires:

- Install go
- Install virtualbox
- Install docker toolbox

Create docker machine in virtualbox:

- Run 'docker-machine create --driver=virtualbox dune'

Source the environment variables:

- Run 'eval "$(docker-machine env dune)"'

Build the docker image locally:

- Run 'docker build -t usul .'

Run usul:

- Run 'docker run --rm -p 8080:8080 --name usul usul'

Get docker ip address:

- Run 'docker-machine ip dune'

Try web editor:

- Install muad'dib (http://github.com/arnoldcano/muaddib)

Try the api:

- Run `curl -H "Content-Type: application/json" -d '{ "language": "ruby", "code": "puts \"hello+world\"" }' http://$(docker-machine ip dune):8080/run`
- Run `curl -H "Content-Type: application/json" -d '{ "language": "python", "code": "print \"hello+world\"" }' http://$(docker-machine ip dune):8080/run`
- Run `curl -H "Content-Type: application/json" -d '{ "language": "js", "code": "console.log(\"hello+world\")" }' http://$(docker-machine ip dune):8080/run`
