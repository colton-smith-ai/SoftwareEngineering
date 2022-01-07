# Author Signature
> :raising_hand:      Colton Smith <br>
> :computer:          Data Science Engineer <br>
> :incoming_envelope: colton.smith.ai@gmail.com <br>
> :octocat:           https://github.com/colton-smith-ai <br>
> :date:              December 2021 <br>

# global-greet-stack
Simple tech-stack application housing: RESTful API Server, Client Interface,
and REDIS database. Application enhances typical `Hello World`, by greeting
a user in their native language :earth_americas:.
- Pedro from Argentina :argentina: would be greeted with, `Hola, Pedro`
- Kartik from Northern-India :india: would be greeted with, `Namaste, Kartik`
- Kaito from Japan :jp: would be greeted with, `Kon'nichiwa, Kaito`

## Microservices
Each service in the stack is deployed by separate Docker containers.
Dockerfiles are located in the respective /api and /client packages. 
Redis database is pulled from Docker Hub. 
Modifying the `.env` file enables developer maximum customization
over any configuration detail and entire application environment.

## Tech Stack
- API Server
  - Written in Go, featuring [Gin Web Framework](https://github.com/gin-gonic/gin)
  - Interacts with Redis Database to retrieve & insert language details
  - Endpoints
    - GET /greeting
      - returns list of available languages to query API
    - POST /greeting
      - adds a new greeting from JSON received in the request body
    - GET /greeting/[language]
      - returns greeting for specified language
  - `$ curl http://localhost:$LOCAL_HOST_API_PORT`
    ``` 
      {
         "about": "Welcome to this API about greetings in various languages.",
         "easterEgg": "There is a hidden endpoint that takes two numbers and adds them together. See if you can solve this MATH problem ;)",
         "help": "This API serves 3 endpoints: 1) GET /greeting = returns list of available languages to query API 2) POST /greeting = adds a new greeting from JSON received in the request body 3) GET /greeting/<language> = returns greeting for specified language"
      }
    ```
- Client Interface
  - Written in Go
  - Uses command line interface to interact with user
  - Receives user input, then greets user in their language
  - Allows user to add new language (which is updated in database via API)
  - GET & POST interactions with API server are performed in background
- Redis Database
  - Schema
    - KEY   : language 
    - VALUE : {language:<language>, hello:<hello>, welcome:<welcome>}
  - Example
    - `{"English":"{"language":"English","hello":"Hello","welcome":"Welcome"}"}`
    - `{"Espanol":"{"language":"Espanol","hello":"Hola","welcome":"Bienvenido"}"}`

# Docker :whale:
Use `docker-compose.yml` to build application ecosystem. `docker-compose` will build the API and
Client images, pull redis image, then start all three containers on same docker network. Client
container expected to be run via command line interface (cli). Client container must be run in
interactive mode (i.e. `docker run -it ...`) because the Go application requires user input.
Redis container uses a docker volume for storage persistent. When finished running the application, 
terminate services and cleanup docker artifacts (i.e. docker network, containers, etc.) with 
`$ docker-compose down`.
- ## Build Docker Images & Run Containers
    ` $ docker-compose run --rm --name gg-client client`
- ## Terminate Services & Cleanup
    `$ docker-compose down`

# EasterEgg
Included in this application is a fun capture-the-flag endpoint.

Steps to get started
1. View/modify LOCAL_HOST_API_PORT in `.env` file
   - Remember localhost port number
2. In terminal, start application
   - `$ docker-compose up`
   - client app will error and fail, that is okay for this exercise
3. Navigate to favorite browser
4. Insert into address bar: localhost + port number found from step 1
   - `http://localhost:LOCAL_HOST_API_PORT`
   - `$ curl http://localhost:$LOCAL_HOST_API_PORT`
5. Read the "easterEgg" hint
6. Have FUN trying different endpoint URLs
   - Few endpoints to explore server
     - `http://localhost:LOCAL_HOST_API_PORT/greeting`
     - `http://localhost:LOCAL_HOST_API_PORT/greeting/English`
     - `http://localhost:LOCAL_HOST_API_PORT/easterEgg` :grimacing:
7. Use [Postman](https://www.postman.com) or `$ curl ...` commands to send GET, POST, or PUT requests to server
