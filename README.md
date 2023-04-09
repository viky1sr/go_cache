## About Struktur

.
├── app
│   ├── controllers
│   │   ├── book_controller.go       // Contains BookController, handles Book related requests
│   │   └── user_controller.go       // Contains UserController, handles User related requests
│   ├── models
│   │   ├── book.go                  // Contains Book model struct
│   │   └── user.go                  // Contains User model struct
│   ├── providers
│   │   ├── app_provider.go          // Contains AppProvider, initializes database and server
│   │   ├── book_provider.go         // Contains BookProvider, sets up Book service and repository
│   │   ├── user_provider.go         // Contains UserProvider, sets up User service and repository
│   │   └── sql_provider.go          // Contains SQLProvider, initializes SQL database
│   ├── repositories
│   │   ├── book_repository.go       // Contains BookRepository interface
│   │   ├── sql_book_repository.go   // Contains SQLBookRepository, implements BookRepository interface for SQL database
│   │   ├── user_repository.go       // Contains UserRepository interface
│   │   └── sql_user_repository.go   // Contains SQLUserRepository, implements UserRepository interface for SQL database
│   ├── routes
│   │   ├── book_route.go            // Contains book related routes
│   │   └── user_route.go            // Contains user related routes
│   ├── services
│   │   ├── book_service.go          // Contains BookService, handles business logic for Book model
│   │   └── user_service.go          // Contains UserService, handles business logic for User model
│   ├── traits
│   │   └── response_trait.go        // Contains ResponseTrait, adds standardized JSON response
│   └── validation
│       ├── book_validator.go        // Contains BookValidator, defines validation rules for Book model
│       └── user_validator.go        // Contains UserValidator, defines validation rules for User model
└── main.go                           // Entry point of the application