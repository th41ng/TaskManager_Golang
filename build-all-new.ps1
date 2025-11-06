# ============================================================================
# BUILD & PUSH ALL - Complete Automation Script
# ============================================================================
# Builds all frontend apps, backend services, creates Docker images and pushes
# 
# Usage:
#   ./build-all-new.ps1 -DockerUsername "thangmicro"
#   ./build-all-new.ps1 -SkipFrontend
#   ./build-all-new.ps1 -SkipBackend  
#   ./build-all-new.ps1 -SkipPush
#   ./build-all-new.ps1 -Tag "v1.0.0"
# ============================================================================

param(
    [Parameter(Mandatory=$false)]
    [string]$DockerUsername = "thangmicro",
    
    [Parameter(Mandatory=$false)]
    [switch]$SkipFrontend,
    
    [Parameter(Mandatory=$false)]
    [switch]$SkipBackend,
    
    [Parameter(Mandatory=$false)]
    [switch]$SkipPush,
    
    [Parameter(Mandatory=$false)]
    [string]$Tag = "latest"
)

$ErrorActionPreference = "Stop"
$rootDir = $PSScriptRoot

Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "BUILD & PUSH ALL SERVICES" -ForegroundColor Cyan
Write-Host "Docker Username: $DockerUsername" -ForegroundColor Yellow
Write-Host "Tag: $Tag" -ForegroundColor Yellow
Write-Host "Skip Frontend: $SkipFrontend" -ForegroundColor Yellow
Write-Host "Skip Backend: $SkipBackend" -ForegroundColor Yellow
Write-Host "Skip Push: $SkipPush" -ForegroundColor Yellow
Write-Host "============================================================================" -ForegroundColor Cyan

# ============================================================================
# FRONTEND BUILD
# ============================================================================
if (-not $SkipFrontend) {
    Write-Host "`n[FRONTEND] Building all frontend apps..." -ForegroundColor Green
    
    # Frontend apps phải build theo thứ tự: remote apps trước, shell sau
    $frontendApps = @(
        @{ Name = "user"; Port = 4001 },
        @{ Name = "project"; Port = 4002 },
        @{ Name = "task"; Port = 4003 },
        @{ Name = "shell"; Port = 5173 }  # Shell build cuối cùng
    )
    
    foreach ($app in $frontendApps) {
        $appName = $app.Name
        $appDir = Join-Path $rootDir "frontend\$appName"
        
        Write-Host "`n[FRONTEND] Building $appName app..." -ForegroundColor Cyan
        
        if (-not (Test-Path $appDir)) {
            Write-Host "ERROR: Directory not found: $appDir" -ForegroundColor Red
            exit 1
        }
        
        Push-Location $appDir
        
        try {
            # Install dependencies if needed
            if (-not (Test-Path "node_modules")) {
                Write-Host "  Installing dependencies..." -ForegroundColor Yellow
                npm install
                if ($LASTEXITCODE -ne 0) { throw "npm install failed" }
            }
            
            # Build the app
            Write-Host "  Building..." -ForegroundColor Yellow
            npm run build
            if ($LASTEXITCODE -ne 0) { throw "npm run build failed" }
            
            # Build Docker image
            Write-Host "  Building Docker image..." -ForegroundColor Yellow
            $imageName = "$DockerUsername/taskmanager-$appName-frontend:$Tag"
            docker build -t $imageName .
            if ($LASTEXITCODE -ne 0) { throw "docker build failed" }
            
            # Push to Docker Hub
            if (-not $SkipPush) {
                Write-Host "  Pushing to Docker Hub..." -ForegroundColor Yellow
                docker push $imageName
                if ($LASTEXITCODE -ne 0) { throw "docker push failed" }
            }
            
            Write-Host "  ✓ $appName completed successfully" -ForegroundColor Green
            
        } catch {
            Write-Host "ERROR building $appName : $_" -ForegroundColor Red
            Pop-Location
            exit 1
        }
        
        Pop-Location
    }
    
    Write-Host "`n[FRONTEND] All frontend apps built successfully!" -ForegroundColor Green
}

# ============================================================================
# BACKEND BUILD
# ============================================================================
if (-not $SkipBackend) {
    Write-Host "`n[BACKEND] Building all backend services..." -ForegroundColor Green
    
    $backendServices = @(
        @{ Name = "gateway"; Path = "gateway" },
        @{ Name = "user-service"; Path = "microservices\user-service" },
        @{ Name = "project-service"; Path = "microservices\project-service" },
        @{ Name = "task-service"; Path = "microservices\task-service" }
    )
    
    foreach ($service in $backendServices) {
        $serviceName = $service.Name
        $serviceDir = Join-Path $rootDir $service.Path
        
        Write-Host "`n[BACKEND] Building $serviceName..." -ForegroundColor Cyan
        
        if (-not (Test-Path $serviceDir)) {
            Write-Host "ERROR: Directory not found: $serviceDir" -ForegroundColor Red
            exit 1
        }
        
        Push-Location $serviceDir
        
        try {
            # Tidy dependencies
            Write-Host "  Running go mod tidy..." -ForegroundColor Yellow
            go mod tidy
            if ($LASTEXITCODE -ne 0) { throw "go mod tidy failed" }
            
            # Build Docker image
            Write-Host "  Building Docker image..." -ForegroundColor Yellow
            $imageName = "$DockerUsername/taskmanager-$serviceName`:$Tag"
            docker build -t $imageName .
            if ($LASTEXITCODE -ne 0) { throw "docker build failed" }
            
            # Push to Docker Hub
            if (-not $SkipPush) {
                Write-Host "  Pushing to Docker Hub..." -ForegroundColor Yellow
                docker push $imageName
                if ($LASTEXITCODE -ne 0) { throw "docker push failed" }
            }
            
            Write-Host "  ✓ $serviceName completed successfully" -ForegroundColor Green
            
        } catch {
            Write-Host "ERROR building $serviceName : $_" -ForegroundColor Red
            Pop-Location
            exit 1
        }
        
        Pop-Location
    }
    
    Write-Host "`n[BACKEND] All backend services built successfully!" -ForegroundColor Green
}

# ============================================================================
# SUMMARY
# ============================================================================
Write-Host "`n============================================================================" -ForegroundColor Cyan
Write-Host "BUILD & PUSH COMPLETED SUCCESSFULLY!" -ForegroundColor Green
Write-Host "============================================================================" -ForegroundColor Cyan

if (-not $SkipFrontend) {
    Write-Host "`nFrontend Images:" -ForegroundColor Yellow
    Write-Host "  - $DockerUsername/taskmanager-user-frontend:$Tag"
    Write-Host "  - $DockerUsername/taskmanager-project-frontend:$Tag"
    Write-Host "  - $DockerUsername/taskmanager-task-frontend:$Tag"
    Write-Host "  - $DockerUsername/taskmanager-shell-frontend:$Tag"
}

if (-not $SkipBackend) {
    Write-Host "`nBackend Images:" -ForegroundColor Yellow
    Write-Host "  - $DockerUsername/taskmanager-gateway:$Tag"
    Write-Host "  - $DockerUsername/taskmanager-user-service:$Tag"
    Write-Host "  - $DockerUsername/taskmanager-project-service:$Tag"
    Write-Host "  - $DockerUsername/taskmanager-task-service:$Tag"
}

if (-not $SkipPush) {
    Write-Host "`nAll images have been pushed to Docker Hub!" -ForegroundColor Green
} else {
    Write-Host "`nImages built locally (not pushed)" -ForegroundColor Yellow
}

Write-Host "`nNext steps:" -ForegroundColor Cyan
Write-Host "  1. Deploy with Nomad: nomad job run nomad/taskmanager.nomad"
Write-Host "  2. Check status: nomad job status taskmanager"
Write-Host "  3. Access Shell: http://172.21.223.107:5173"
Write-Host "============================================================================" -ForegroundColor Cyan
