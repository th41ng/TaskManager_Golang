# ============================================================================
# BUILD, DOCKER, AND PUSH ALL - TaskManager Project
# ============================================================================
# Script này sẽ:
# 1. Build tất cả frontend apps
# 2. Build tất cả backend services
# 3. Tạo Docker images
# 4. Push lên Docker Hub

param(
    [string]$DockerUsername = "thangmicro",
    [switch]$SkipFrontend,
    [switch]$SkipBackend,
    [switch]$SkipPush,
    [string]$Tag = "latest"
)

$ErrorActionPreference = "Stop"

Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "  BUILD AND PUSH ALL - TaskManager Project" -ForegroundColor Cyan
Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host ""

# ============================================================================
# FRONTEND BUILDS
# ============================================================================
if (-not $SkipFrontend) {
    Write-Host "[1/2] Building Frontend Apps..." -ForegroundColor Yellow
    Write-Host ""
    
    $frontendApps = @("user", "project", "task", "shell")
    
    foreach ($app in $frontendApps) {
        Write-Host "  Building $app app..." -ForegroundColor Green
        Set-Location "frontend\$app"
        
        # Install dependencies if needed
        if (-not (Test-Path "node_modules")) {
            Write-Host "    Installing dependencies..." -ForegroundColor Gray
            npm install
        }
        
        # Build
        Write-Host "    Building..." -ForegroundColor Gray
        npm run build
        
        if ($LASTEXITCODE -ne 0) {
            Write-Host "    ❌ Build failed for $app!" -ForegroundColor Red
            Set-Location ..\..
            exit 1
        }
        
        Write-Host "    ✅ Build successful for $app" -ForegroundColor Green
        Write-Host ""
        
        Set-Location ..\..
    }
    
    Write-Host "✅ All frontend apps built successfully!" -ForegroundColor Green
    Write-Host ""
}

# ============================================================================
# BACKEND BUILDS
# ============================================================================
if (-not $SkipBackend) {
    Write-Host "[2/2] Building Backend Services..." -ForegroundColor Yellow
    Write-Host ""
    
    $services = @(
        @{Name="gateway"; Path="gateway"},
        @{Name="user-service"; Path="microservices\user-service"},
        @{Name="project-service"; Path="microservices\project-service"},
        @{Name="task-service"; Path="microservices\task-service"}
    )
    
    foreach ($service in $services) {
        Write-Host "  Building $($service.Name)..." -ForegroundColor Green
        Set-Location $service.Path
        
        # Build Go binary
        Write-Host "    Compiling Go binary..." -ForegroundColor Gray
        go build -o main .
        
        if ($LASTEXITCODE -ne 0) {
            Write-Host "    ❌ Build failed for $($service.Name)!" -ForegroundColor Red
            Set-Location ..\..
            exit 1
        }
        
        Write-Host "    ✅ Build successful for $($service.Name)" -ForegroundColor Green
        Write-Host ""
        
        Set-Location ..\..
        if ($service.Path -like "*microservices*") {
            Set-Location ..
        }
    }
    
    Write-Host "✅ All backend services built successfully!" -ForegroundColor Green
    Write-Host ""
}

# ============================================================================
# DOCKER IMAGES
# ============================================================================
Write-Host "Building Docker Images..." -ForegroundColor Yellow
Write-Host ""

$images = @(
    # Frontend
    @{Name="user-frontend"; Path="frontend\user"; Tag="$DockerUsername/user-frontend:$Tag"},
    @{Name="project-frontend"; Path="frontend\project"; Tag="$DockerUsername/project-frontend:$Tag"},
    @{Name="task-frontend"; Path="frontend\task"; Tag="$DockerUsername/task-frontend:$Tag"},
    @{Name="shell-frontend"; Path="frontend\shell"; Tag="$DockerUsername/shell-frontend:$Tag"},
    
    # Backend
    @{Name="gateway"; Path="gateway"; Tag="$DockerUsername/gateway:$Tag"},
    @{Name="user-service"; Path="microservices\user-service"; Tag="$DockerUsername/user-service:$Tag"},
    @{Name="project-service"; Path="microservices\project-service"; Tag="$DockerUsername/project-service:$Tag"},
    @{Name="task-service"; Path="microservices\task-service"; Tag="$DockerUsername/task-service:$Tag"}
)

foreach ($image in $images) {
    if ($SkipFrontend -and $image.Name -like "*-frontend") {
        continue
    }
    if ($SkipBackend -and $image.Name -notlike "*-frontend") {
        continue
    }
    
    Write-Host "  Building Docker image: $($image.Tag)" -ForegroundColor Green
    
    docker build -t $image.Tag $image.Path
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "    ❌ Docker build failed for $($image.Name)!" -ForegroundColor Red
        exit 1
    }
    
    Write-Host "    ✅ Image built: $($image.Tag)" -ForegroundColor Green
    Write-Host ""
}

Write-Host "✅ All Docker images built successfully!" -ForegroundColor Green
Write-Host ""

# ============================================================================
# PUSH TO DOCKER HUB
# ============================================================================
if (-not $SkipPush) {
    Write-Host "Pushing Docker Images to Docker Hub..." -ForegroundColor Yellow
    Write-Host ""
    
    # Check if logged in
    Write-Host "  Checking Docker login..." -ForegroundColor Gray
    docker info | Out-Null
    if ($LASTEXITCODE -ne 0) {
        Write-Host "  Please login to Docker Hub first:" -ForegroundColor Yellow
        docker login
    }
    
    foreach ($image in $images) {
        if ($SkipFrontend -and $image.Name -like "*-frontend") {
            continue
        }
        if ($SkipBackend -and $image.Name -notlike "*-frontend") {
            continue
        }
        
        Write-Host "  Pushing: $($image.Tag)" -ForegroundColor Green
        
        docker push $image.Tag
        
        if ($LASTEXITCODE -ne 0) {
            Write-Host "    ❌ Push failed for $($image.Name)!" -ForegroundColor Red
            exit 1
        }
        
        Write-Host "    ✅ Pushed: $($image.Tag)" -ForegroundColor Green
        Write-Host ""
    }
    
    Write-Host "✅ All images pushed successfully!" -ForegroundColor Green
    Write-Host ""
}

# ============================================================================
# SUMMARY
# ============================================================================
Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "  ✅ ALL TASKS COMPLETED SUCCESSFULLY!" -ForegroundColor Cyan
Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Summary:" -ForegroundColor White
if (-not $SkipFrontend) {
    Write-Host "  ✅ Frontend apps built and dockerized" -ForegroundColor Green
}
if (-not $SkipBackend) {
    Write-Host "  ✅ Backend services built and dockerized" -ForegroundColor Green
}
if (-not $SkipPush) {
    Write-Host "  ✅ All images pushed to Docker Hub" -ForegroundColor Green
} else {
    Write-Host "  ⏭️  Push skipped (use without -SkipPush to push)" -ForegroundColor Yellow
}
Write-Host ""
Write-Host "Next steps:" -ForegroundColor White
Write-Host "  1. Deploy to Nomad: nomad job run nomad\taskmanager.nomad" -ForegroundColor Gray
Write-Host "  2. Check status: nomad job status taskmanager" -ForegroundColor Gray
Write-Host ""
