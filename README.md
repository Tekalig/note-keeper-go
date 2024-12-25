# Note Keeper

## Running the Project

### Prerequisites

- Go installed on your machine. You can download it from [Go website](https://golang.org/dl/).

### Steps

1. **Clone the Repository**

   ```sh
   git clone https://github.com/your-username/note-keeper.git
   cd note-keeper
   ```

2. **Set Environment Variables**

   Set the required environment variables `HASURA_URL` and `HASURA_ADMIN_SECRET`:

   ```sh
   export HASURA_URL="your_hasura_url"
   export HASURA_ADMIN_SECRET="your_hasura_admin_secret"
   ```

   Reload your shell profile:

   ```sh
   source ~/.bashrc  # or source ~/.zshrc
   ```

3. **Build and Run the Project**

   Navigate to the project directory and run the following commands:

   ```sh
   go mod tidy  # Ensure all dependencies are installed
   go build -o note-keeper ./cmd/server
   ./note-keeper
   ```

   This will build the project and start the server on port 8000. You should see the log message "Server running on port 8000".
