// ============================================================================
// SHARED TYPES FOR MICRO FRONTEND COMMUNICATION
// ============================================================================
// File này define types cho việc communication giữa Shell và Remote apps

/**
 * Shell Data - Dữ liệu Shell truyền xuống tất cả Remote apps
 */
export interface ShellData {
  user?: {
    id?: number
    name?: string
    email?: string
    avatar?: string
  }
  theme?: 'light' | 'dark'
  token?: string
  [key: string]: any  // Allow thêm properties
}

/**
 * Base props cho tất cả Remote Content Components
 */
export interface RemoteContentProps {
  shellData?: ShellData
  onNavigate?: (path: string) => void
}
