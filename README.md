# JXSaver

CLI tool for validating and saving JSON and XML as files.

- ``go run main.go --<format> <'data'>`` for saving data;
- Formats supported: ``json``, ``xml``;
- Invalid data won't be saved;
- Duplicate data won't be saved.
- Data is saved in current directory in a new file with corresponding format extension.
