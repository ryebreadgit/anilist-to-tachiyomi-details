# Manga Details

This program, `manga-details`, is designed to download Anilist data for a manga and save it in the `details.json` format used by Tachiyomi. It provides a command-line interface for the user to input a manga ID, and it fetches the necessary information from Anilist's GraphQL API.

## Prerequisites

- Go programming language (version 1.16 or higher)

## Installation

1. Clone the repository or download the source code files.
2. Open a terminal and navigate to the project directory.
3. Build the program by running the following command:

```shell
go build
```

## Usage

1. Run the compiled executable or use the Go command:
```shell
go run main.go
```
2. Enter the manga ID when prompted. You can find the manga ID on Anilist's website or through other sources.
3. Wait for the program to download the manga details and save them in the details.json file.
4. The program will also download the manga cover image and save it as cover.jpg in the same directory.
5. Once the process is complete, a success message will be displayed.
6. Press Enter to exit the program.

## Dependencies

This program relies on the following third-party Go packages:
- buger/jsonparser: A low-level JSON parser to extract data from JSON structures efficiently.
- imroc/req: A Go HTTP client library with chainable API, built-in JSON support, and other useful features.

## Contributing

Contributions to this program are welcome. If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on the GitHub repository.

## License

This program is licensed under the MIT License. Feel free to modify and distribute it as needed.

## Acknowledgements

This program was inspired by the need to import manga details from Anilist into Tachiyomi. Thanks to the developers of Anilist and Tachiyomi for providing the necessary APIs and resources.


Please note that you may need to adjust the links and placeholders in the Markdown file according to your project's actual structure and requirements.