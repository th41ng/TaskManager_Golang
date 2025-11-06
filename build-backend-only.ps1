# ============================================================================
# BUILD AND PUSH BACKEND SERVICES ONLY
# ============================================================================

param(
    [string]$DockerUsername = "thangmicro",
    [switch]$SkipPush,
    [string]$Tag = "latest"
)

$ErrorActionPreference = "Stop"

Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "  BUILD BACKEND SERVICES" -ForegroundColor Cyan
Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host ""

$services = @(
    @{Name="gateway"; Path="gateway"},
    @{Name="user-service"; Path="microservices\user-service"},
    @{Name="project-service"; Path="microservices\project-service"},
    @{Name="task-service"; Path="microservices\task-service"}
)

foreach ($service in $services) {
    Write-Host "[$($service.Name)] Building..." -ForegroundColor Yellow
    Set-Location $service.Path
    
    # Build Go binary
    Write-Host "  Compiling Go binary..." -ForegroundColor Gray
    go build -o main .
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "  ❌ Build failed!" -ForegroundColor Red
        Set-Location ..\..
        if ($service.Path -like "*microservices*") {
            Set-Location ..
        }
        exit 1
    }
    
    # Docker build
    Write-Host "  Building Docker image..." -ForegroundColor Gray
    $imageTag = "$DockerUsername/$($service.Name):$Tag"
    docker build -t $imageTag .
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "  ❌ Docker build failed!" -ForegroundColor Red
        Set-Location ..\..
        if ($service.Path -like "*microservices*") {
            Set-Location ..
        }
        exit 1
    }
    
    Write-Host "  ✅ Built: $imageTag" -ForegroundColor Green
    
    # Push
    if (-not $SkipPush) {
        Write-Host "  Pushing to Docker Hub..." -ForegroundColor Gray
        docker push $imageTag
        
        if ($LASTEXITCODE -ne 0) {
            Write-Host "  ❌ Push failed!" -ForegroundColor Red
            Set-Location ..\..
            if ($service.Path -like "*microservices*") {
                Set-Location ..
            }
            exit 1
        }
        
        Write-Host "  ✅ Pushed: $imageTag" -ForegroundColor Green
    }
    
    Write-Host ""
    Set-Location ..\..
    if ($service.Path -like "*microservices*") {
        Set-Location ..
    }
}

Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "  ✅ ALL BACKEND SERVICES COMPLETED!" -ForegroundColor Cyan
Write-Host "============================================================================" -ForegroundColor Cyan
