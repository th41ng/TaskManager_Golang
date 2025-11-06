# ============================================================================
# BUILD AND PUSH FRONTEND ONLY
# ============================================================================

param(
    [string]$DockerUsername = "thangmicro",
    [switch]$SkipPush,
    [string]$Tag = "latest"
)

$ErrorActionPreference = "Stop"

Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "  BUILD FRONTEND APPS" -ForegroundColor Cyan
Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host ""

$frontendApps = @("user", "project", "task", "shell")

foreach ($app in $frontendApps) {
    Write-Host "[$app] Building..." -ForegroundColor Yellow
    Set-Location "frontend\$app"
    
    # Install dependencies
    if (-not (Test-Path "node_modules")) {
        Write-Host "  Installing dependencies..." -ForegroundColor Gray
        npm install
    }
    
    # Build
    Write-Host "  Building..." -ForegroundColor Gray
    npm run build
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "  ❌ Build failed!" -ForegroundColor Red
        Set-Location ..\..
        exit 1
    }
    
    # Docker build
    Write-Host "  Building Docker image..." -ForegroundColor Gray
    $imageTag = "$DockerUsername/${app}-frontend:$Tag"
    docker build -t $imageTag .
    
    if ($LASTEXITCODE -ne 0) {
        Write-Host "  ❌ Docker build failed!" -ForegroundColor Red
        Set-Location ..\..
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
            exit 1
        }
        
        Write-Host "  ✅ Pushed: $imageTag" -ForegroundColor Green
    }
    
    Write-Host ""
    Set-Location ..\..
}

Write-Host "============================================================================" -ForegroundColor Cyan
Write-Host "  ✅ ALL FRONTEND APPS COMPLETED!" -ForegroundColor Cyan
Write-Host "============================================================================" -ForegroundColor Cyan
