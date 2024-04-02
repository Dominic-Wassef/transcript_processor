# Transcript Processor

## Backend

### Structure

- **handlers**: Contains HTTP handlers that manage API endpoints.
- **helpers**: Provides utility functions, including text cleaning operations.
- **loader**: Responsible for loading and parsing utterance files.
- **processor**: Processes the utterances by cleaning and merging them.
- **middleware**: Contains middleware functions, including CORS configuration.
- **models**: Defines the data structures used across the application.
- **main.go**: The entry point of the application, setting up the server and routes.

## Function Signatures and Descriptions

### Handlers Package

- `func TranscriptHandler(w http.ResponseWriter, r *http.Request)`
  - **Description**: Handles the HTTP request to process and return a transcript of processed utterances.

### Helpers Package

- `func CleanText(text string) string`
  - **Description**: Applies various cleaning functions to the input text, including removing duplicate words and punctuation.

- `func RemoveDuplicateWords(text string) string`
  - **Description**: Removes consecutive duplicate words from the input text.

- `func RemoveDuplicatePunctuation(text string) string`
  - **Description**: Removes consecutive duplicate punctuation from the input text.

### Loader Package

- `func LoadUtterances(dir string) ([]models.Utterance, error)`
  - **Description**: Reads and parses utterance files from the given directory, returning a slice of utterances.

### Processor Package

- `func ProcessUtterances(utterances []models.Utterance) ([]models.Utterance, error)`
  - **Description**: Processes a slice of utterances by cleaning and merging them, based on defined logic.

- `func MergeUtterances(utterances []models.Utterance) ([]models.Utterance, error)`
  - **Description**: Combines consecutive utterances by the same speaker when appropriate, based on timing and content.

- `func shouldMergeWithPreviousUtterance(prev, current models.Utterance) bool`
  - **Description**: Determines if the current utterance should be merged with the previous one, based on timing, speaker, and punctuation.

- `func startsWithLowercase(text string) bool`
  - **Description**: Checks if the first character of the text is a lowercase letter.

### Middleware Package

- `func CorsMiddleware(next http.Handler) http.Handler`
  - **Description**: Adds basic CORS headers to responses, allowing for cross-origin requests.

### Models Package

- **Utterance Struct**
  - **Description**: Represents a segment of speech in a transcript, containing details such as speaker, text, and timestamps.

- **Transcript Struct**
  - **Description**: Aggregates multiple `Utterance` instances, representing a complete transcript.


### Special Note

Modify the directory path in `handlers/handlers.go` according to your local setup:

Currently it is set to:
```go
dir := "/Users/dominic/Desktop/golang/transcriptprocessor/utterances"
```

Please set it to:
Currently it is set to:
```go
dir := "Your/Absolute/Directory/Path"
```

### Running the Backend

1. Navigate to the backend directory within the project structure.
2. Run `go mod tidy` to ensure all dependencies are correctly installed.
3. Execute `go run main.go` to start the server. By default, it will listen on `localhost:8080`.
4. The server routes are configured to respond to API requests. Use `/api/transcript` to fetch processed transcripts.

### Testing

- Comprehensive tests are available for each package. To run tests, execute `go test -v ./...` at the root of the backend directory. This command recursively runs all tests.
- For individual packages, navigate to the package directory and run `go test`.

## Frontend

### Structure

- **src**: Source directory containing React components and TypeScript types.
- **public**: Contains static files like HTML entry point and icons.
- **.env**: Environment file to configure the backend API URL.

### Key Components

- `App.tsx`: The main React component that fetches and displays utterance data from the backend.
- `types.ts`: TypeScript definitions for the data structures used in the frontend.

### Running the Frontend

1. Navigate to the frontend directory and run `npm install` to install dependencies.
2. Start the application with `npm start`. This will launch the app in your default web browser at `localhost:3000`.
3. The application will automatically connect to the backend API to fetch and display transcripts.

### Environment Configuration

- The `.env` file in the frontend directory specifies the backend API URL. Adjust the `REACT_APP_BACKEND_URL` variable as needed to point to your backend server.