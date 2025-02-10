# magica
This repository contains a refactored version of Google's [Magika](https://github.com/google/magika) Go module, adjusted for proper module resolution and ease of installation. 

Magika is a file type identification tool that leverages machine learning models (via Microsoft's ONNX Runtime) to classify file formats.

## Features

- Refactored module structure to allow direct usage as `github.com/postfix/magika`.
- Installation script for required dependencies, including:
  - Microsoft's ONNX Runtime
- Compatible with standard Go build workflows.

## Installation

### To install ONNX Runtime and necessary C includes, use the latest setup script from postfix/setup-scripts:

```bash
git clone https://github.com/postfix/setup-scripts.git
cd setup-scripts
sudo ./install_onnx_runtime.sh
go install github.com/postfix/magika
```
### LICENSE

This project is part of the Magika, which is licensed under the Apache 2.0 License.
