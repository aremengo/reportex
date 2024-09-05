# reportex

Very basic skeleton of a containerized tool for generating pdf using the XeLaTeX engine.

## Build and run

```bash
docker build -t reportex:latest .

docker run -v ./reports:/app/reports -it reportex:latest
```