# Crosssight - Static Crossplane YAML Scanner
Overview
Crosssight is a powerful static Crossplane YAML scanner that allows you to scan your Crossplane Infrastructure-as-Code (IAC) YAML files for security and compliance issues. It leverages Checkov by Bridgecrew to provide detailed reports on potential vulnerabilities and misconfigurations.

### Architecture
Crosssight is built using the following Go packages and libraries:

Cobra CLI: Cobra is a popular library for building command-line applications in Go. It provides a simple and elegant way to define commands, flags, and subcommands for the CLI interface. In Crosssight, Cobra is used to define the command-line interface for running the scanning process.

Viper: Viper is a configuration management library for Go. It is used to handle configuration settings and flags in Crosssight. Viper allows users to define configuration options through various sources, such as environment variables, configuration files, and command-line flags.

The combination of Cobra and Viper makes Crosssight's CLI interface user-friendly and customizable, providing users with a seamless experience when running scans and configuring the tool's behavior.

### Installation
To use Crosssight, you need to have Go installed on your machine. If you don't have Go installed, you can download it from the official Go website.

Once you have Go installed, you can build Crosssight from the source code by following these steps:

Clone the Crosssight repository to your local machine:

```bash
Copy code
git clone https://github.com/crosssight/crosssight.git
cd crosssight
```
Build the Crosssight binary and place it in the bin directory:

```bash
make build
```
This will create a binary named crosssight and place it in the bin directory.

Optionally, you can move the crosssight binary to a location in your system's PATH for easier access.

### Usage
To run Crosssight and scan your Crossplane IAC YAML files, use the following command:

```bash
crosssight scan -r /path/to/rules.yaml -f /path/to/configuration.yaml
```

Replace /path/to/rules.yaml with the path to the YAML file defining the specific S3 Crossplane bucket kind rules you want to use for scanning.
Replace /path/to/configuration.yaml with the path to the YAML file containing the Crossplane configuration rules you want to scan. 

### Example:

```bash
crosssight scan -r /Users/jonathanpick/GolandProjects/static_crossplane_scanner/aws/tests/demo_rules.yaml -f /Users/jonathanpick/GolandProjects/static_crossplane_scanner/aws/tests/demo_bucket.yaml
```

#### Additional Notes
The crosssight scan command will execute Checkov using the provided rules file to scan the Crossplane IAC YAML file.
The output will provide a detailed report of any security and compliance issues found in the Crossplane IAC YAML file.
Be sure to keep your Crosssight binary and Checkov tool up-to-date to take advantage of the latest security checks.
Contributing
If you encounter any issues or have suggestions for improvement, please feel free to submit a pull request or create an issue in the Crosssight repository.

### License
This project is licensed under the MIT License.

### Acknowledgments
Thank you to the Bridgecrew team for developing Checkov, enabling Crosssight to provide valuable security scanning for Crossplane YAML files.

