# goreq

`goreq` is a command-line tool for making HTTP requests. It simplifies the process of sending requests and handling responses, making it easier to interact with APIs and web services. With `goreq`, you can perform various HTTP methods, manage headers, handle authentication, follow redirects, use proxies, and more.

Key features:
- Simple syntax for making HTTP requests
- Support for GET, POST, and other HTTP methods
- Ability to add headers and parameters
- Basic authentication and proxy support
- Options to follow redirects and ignore SSL
- Save responses to files and customize user agents
- Verbose output for debugging

`goreq` is designed to be easy to use and integrate into your workflow, providing a powerful tool for developers working with HTTP APIs.

## Installation

To install `goreq` using Homebrew, follow these steps:

1. Tap the repository:
    ```sh
    brew tap vvvdwbvvv/goreq https://github.com/vvvdwbvvv/goreq
    ```

2. Install `goreq`:
    ```sh
    brew install goreq
    ```

## Usage

### Basic GET request
```sh
greq -url https://api.example.com/
```
### POST with data
```sh
greq -url https://api.example.com/ -X POST -d '{"name":"test"}'
```

## Contributing

We welcome contributions to the `goreq` project! Here are some guidelines to help you get started:

1. **Fork the repository**: Click the "Fork" button at the top right of the repository page to create a copy of the repository in your GitHub account.

2. **Clone your fork**: Clone your forked repository to your local machine.
    ```sh
    git clone https://github.com/your-username/goreq.git
    cd goreq
    ```

3. **Create a new branch**: Create a new branch for your feature or bug fix.
    ```sh
    git checkout -b my-feature-branch
    ```

4. **Make your changes**: Make your changes to the codebase. Ensure your code follows the project's coding standards and includes appropriate tests.

5. **Commit your changes**: Commit your changes with a descriptive commit message.
    ```sh
    git add .
    git commit -m "Description of your changes"
    ```

6. **Push to your fork**: Push your changes to your forked repository.
    ```sh
    git push origin my-feature-branch
    ```

7. **Create a pull request**: Go to the original repository and create a pull request from your forked repository. Provide a clear description of your changes and any relevant information.

8. **Review process**: Your pull request will be reviewed by the maintainers. Be prepared to make any necessary changes based on feedback.

Thank you for contributing!
## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.