declare module 'taskApp/TaskApp' {
  export const CreateTaskDialog: React.ComponentType<{
    open: boolean
    onOpenChange: (open: boolean) => void
    projectId?: number
    onSuccess?: (task: any) => void
  }>
  
  export const EditTaskDialog: React.ComponentType<{
    open: boolean
    onOpenChange: (open: boolean) => void
    task: any
    onSuccess?: (task: any) => void
  }>
}

declare module 'projectApp/ProjectApp' {
  export const CreateProjectDialog: React.ComponentType<{
    open: boolean
    onOpenChange: (open: boolean) => void
    onSuccess?: (project: any) => void
  }>
  
  export const EditProjectDialog: React.ComponentType<{
    open: boolean
    onOpenChange: (open: boolean) => void
    project: any
    onSuccess?: (project: any) => void
  }>
}

declare module 'userApp/UserApp' {
  const UserApp: React.ComponentType
  export default UserApp
}
