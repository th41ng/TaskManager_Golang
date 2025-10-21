<#
PowerShell helper to run Consul and Nomad in dev mode, build images and run Nomad jobs (local dev)
Usage: run from repository root: .\nomad\run_nomad_dev.ps1

This script will:
- verify docker, consul and nomad are available
- check Docker daemon is running
- build docker images
- start Consul and Nomad in dev mode (if available)
- run the Nomad job files in ./nomad

If a command is missing the script will print instructions and exit.
#>

function Test-Command($name) {
	return (Get-Command $name -ErrorAction SilentlyContinue) -ne $null
}

Write-Host "Checking prerequisites..."

if (-not (Test-Command docker)) {
	Write-Host "ERROR: Docker CLI not found. Please install Docker Desktop or Docker Engine and ensure 'docker' is in your PATH." -ForegroundColor Red
	Write-Host "Windows: https://www.docker.com/get-started" -ForegroundColor Yellow
	exit 1
}

# Check docker daemon
try {
	docker info | Out-Null
} catch {
	Write-Host "ERROR: Docker daemon not running or accessible. Start Docker Desktop and try again." -ForegroundColor Red
	exit 1
}

if (-not (Test-Command consul)) {
	Write-Host "ERROR: 'consul' not found. Install Consul and add to PATH." -ForegroundColor Red
	Write-Host "Windows (choco): choco install consul -y" -ForegroundColor Yellow
	Write-Host "https://www.consul.io/downloads" -ForegroundColor Yellow
	exit 1
}

if (-not (Test-Command nomad)) {
	Write-Host "ERROR: 'nomad' not found. Install Nomad and add to PATH." -ForegroundColor Red
	Write-Host "Windows (choco): choco install nomad -y" -ForegroundColor Yellow
	Write-Host "https://www.nomadproject.io/downloads" -ForegroundColor Yellow
	exit 1
}

Write-Host "All prerequisites found. Building docker images..."
try {
	docker build -t taskmanager/gateway:latest ./gateway
	docker build -t taskmanager/user-service:latest ./microservices/user-service
	docker build -t taskmanager/task-service:latest ./microservices/task-service
	docker build -t taskmanager/project-service:latest ./microservices/project-service
} catch {
	Write-Host "ERROR: Docker build failed: $_" -ForegroundColor Red
	exit 1
}

Write-Host "Starting Consul (dev)"
try {
	Start-Process -FilePath "consul" -ArgumentList "agent -dev -bind=127.0.0.1" -NoNewWindow -PassThru | Out-Null
	Start-Sleep -Seconds 3
} catch {
	Write-Host "ERROR: Failed to start consul: $_" -ForegroundColor Red
	exit 1
}

Write-Host "Starting Nomad (dev)"
try {
	Start-Process -FilePath "nomad" -ArgumentList "agent -dev -bind=127.0.0.1" -NoNewWindow -PassThru | Out-Null
	Start-Sleep -Seconds 5
} catch {
	Write-Host "ERROR: Failed to start nomad: $_" -ForegroundColor Red
	exit 1
}

Write-Host "Running Nomad jobs..."
try {
	& nomad job run .\nomad\user-service.nomad
	& nomad job run .\nomad\task-service.nomad
	& nomad job run .\nomad\project-service.nomad
	& nomad job run .\nomad\gateway.nomad
} catch {
	Write-Host "ERROR: Failed to run nomad jobs: $_" -ForegroundColor Red
	exit 1
}

Write-Host "Status:"
nomad status

Write-Host "Consul UI: http://127.0.0.1:8500  Nomad UI: http://127.0.0.1:4646"
